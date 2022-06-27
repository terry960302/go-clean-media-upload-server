package usecase

import (
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"log"
	"mime/multipart"
	"os"
	"sort"
	"sync"

	"github.com/terry960302/go-clean-media-upload-server/adapter/repository"
	"github.com/terry960302/go-clean-media-upload-server/config"
	"github.com/terry960302/go-clean-media-upload-server/domain"

	"cloud.google.com/go/storage"
	"github.com/labstack/echo/v4"
	"google.golang.org/api/option"
)

const prefix = "https://storage.googleapis.com/"

type ImageMetadataUsecase struct {
	domain.ImageMetadataUsecase
	imageRepo repository.ImageMetadataRepository
	mediaRepo repository.MediaMetadataRepository
}

func NewImageUsecase(imageRepo repository.ImageMetadataRepository, mediaRepo repository.MediaMetadataRepository) *ImageMetadataUsecase {
	usecase := &ImageMetadataUsecase{imageRepo: imageRepo, mediaRepo: mediaRepo}
	usecase.ImageMetadataUsecase = interface{}(usecase).(*ImageMetadataUsecase)
	return usecase
}

type UploadImagesRes struct {
	index int
	url   string
	err   error
}

type ImageConfig struct {
	Width  int
	Height int
	Volume int64
	Format string
}

func (i *ImageMetadataUsecase) SaveImageMedia(url string, config ImageConfig) error {

	media := domain.MediaMetadata{
		Url: url,
	}
	mediaId, err := i.mediaRepo.Create(media)
	if err != nil {
		log.Fatal(err)
		return err
	}

	image := domain.ImageMetadata{
		MediaID: mediaId,
		Width:   fmt.Sprint(config.Width),
		Height:  fmt.Sprint(config.Height),
		Volume:  fmt.Sprint(config.Volume),
		Format:  config.Format,
	}
	imgId, err := i.imageRepo.Create(image)
	if err != nil {
		log.Fatal(err)
		return err
	}

	fmt.Printf("%s img metadata is created", fmt.Sprint(imgId))
	return nil
}

func (i *ImageMetadataUsecase) UploadImages(fileHeaders []*multipart.FileHeader, ctx echo.Context) []string {

	uploadResChan := make(chan UploadImagesRes)
	var wg sync.WaitGroup
	wg.Add(len(fileHeaders))

	for index, header := range fileHeaders {
		go i.uploadImage(ctx, &wg, index, header, uploadResChan)
	}

	responses := []UploadImagesRes{}

	go func() {
		wg.Wait()
		close(uploadResChan)
	}()

	for {
		select {
		case res := <-uploadResChan:
			responses = append(responses, res)

			if len(responses) == len(fileHeaders) {
				urls := processUploadedImgs(responses)
				return urls
			}
			// default:
			// 	fmt.Println("[UploadImages] No channel is ready")
		}
	}
}

func (i *ImageMetadataUsecase) uploadImage(ctx echo.Context, wg *sync.WaitGroup, index int, header *multipart.FileHeader, uploadResChan chan UploadImagesRes) {
	defer wg.Done()

	multipartFile, err := header.Open()
	if err != nil {
		log.Fatal(err)
		uploadResChan <- UploadImagesRes{
			index: index,
			url:   "",
			err:   err,
		}
		return
	}

	defer multipartFile.Close()

	// buffer 로 처리
	// buf := bytes.NewBuffer(nil)
	// if _, err := io.Copy(buf, multipartFile); err != nil {
	// 	uploadResChan <- UploadImagesRes{
	// 		index: index,
	// 		url:   "",
	// 		err:   err,
	// 	}
	// 	return
	// }
	file, err := os.Create(header.Filename)
	if err != nil {
		log.Fatal(err)
		uploadResChan <- UploadImagesRes{
			index: index,
			url:   "",
			err:   err,
		}
		return
	}
	defer file.Close()

	if _, err = io.Copy(file, multipartFile); err != nil {
		log.Fatal(err)
		uploadResChan <- UploadImagesRes{
			index: index,
			url:   "",
			err:   err,
		}
		return
	}
	defer os.Remove(header.Filename)

	imgConfig, err := getImageConfig(header, file)
	if err != nil {
		log.Fatal(err)
		uploadResChan <- UploadImagesRes{
			index: index,
			url:   "",
			err:   err,
		}
		return
	}

	url, err := uploadToStorage(ctx, file, header.Filename)
	if err != nil {
		log.Fatal(err)
		uploadResChan <- UploadImagesRes{
			index: index,
			url:   "",
			err:   err,
		}
		return

	}

	if err := i.SaveImageMedia(url, imgConfig); err != nil {
		log.Fatal(err)
		uploadResChan <- UploadImagesRes{
			index: index,
			url:   "",
			err:   err,
		}
		return
	}

	uploadResChan <- UploadImagesRes{
		index: index,
		url:   url,
		err:   nil,
	}
	return
}

// Uplaod to GCP Storage
func uploadToStorage(ctx echo.Context, file *os.File, fileName string) (string, error) {
	c := ctx.Request().Context()
	client, err := storage.NewClient(c, option.WithCredentialsFile((config.C.Storage.CredPath)))
	if err != nil {
		panic(err)
	}
	defer client.Close()

	dir := "test"
	dst := dir + "/" + fileName

	storageWriter := client.Bucket(config.C.Storage.BucketName).Object(dst).NewWriter(c)
	defer storageWriter.Close()

	if _, err = io.Copy(storageWriter, file); err != nil {
		return "", fmt.Errorf("Upload Failed > io.Copy: %v", err)
	}
	if err := storageWriter.Close(); err != nil {
		return "", fmt.Errorf("Upload Failed > Writer.Close: %v", err)
	}

	url := prefix + config.C.Storage.BucketName + "/" + dst
	fmt.Printf("[Storage Uploaded] file : %s", fileName)
	return url, nil
}

func processUploadedImgs(responses []UploadImagesRes) []string {
	rawRes := []UploadImagesRes{}
	urls := []string{}

	// logging error
	for _, res := range responses {
		if res.err != nil {
			log.Fatalf("%s th file error occured : %s", fmt.Sprint(res.index), res.err.Error())
		} else {
			rawRes = append(rawRes, res)
		}
	}

	// sort
	sort.Slice(rawRes, func(i, j int) bool {
		return rawRes[i].index < rawRes[j].index
	})

	// extract only urls
	for _, res := range rawRes {
		urls = append(urls, res.url)
	}

	return urls
}

func getImageConfig(header *multipart.FileHeader, file *os.File) (ImageConfig, error) {
	// File read index could set to end because of previous functions related to os.File
	// Set file read index to begining before run 'Decode' function
	if _, err := file.Seek(0, 0); err != nil {
		log.Fatal(err)
		return ImageConfig{}, err
	}

	fileInfo, err := file.Stat()
	if err != nil {
		log.Fatal(err)
		return ImageConfig{}, err
	}

	imgConfig, format, err := image.Decode(file)
	if err != nil {
		log.Fatal(err)
		return ImageConfig{}, err
	}

	width := imgConfig.Bounds().Dx()
	height := imgConfig.Bounds().Dy()
	volume := fileInfo.Size() / 1024 // KB

	config := ImageConfig{
		Width:  width,
		Height: height,
		Format: format,
		Volume: volume,
	}

	fmt.Printf("[ImgConfig] file : %s", config)

	return config, nil

}
