package usecase

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"media-upload-server/adapter/repository"
	"media-upload-server/config"
	"media-upload-server/domain"
	"mime/multipart"
	"sort"
	"sync"

	"cloud.google.com/go/storage"
	"github.com/labstack/echo/v4"
)

type ImageMetadataUsecase struct {
	domain.ImageMetadataUsecase
	imageRepo repository.ImageMetadataRepository
}

func NewImageUsecase(imageRepo repository.ImageMetadataRepository) *ImageMetadataUsecase {
	usecase := &ImageMetadataUsecase{imageRepo: imageRepo}
	usecase.ImageMetadataUsecase = interface{}(usecase).(*ImageMetadataUsecase)
	return usecase
}

type UploadImagesRes struct {
	index int
	url   string
	err   error
}

func (i *ImageMetadataUsecase) UploadImages(fileHeaders []*multipart.FileHeader, ctx echo.Context) []string {

	uploadResChan := make(chan UploadImagesRes)
	var wg sync.WaitGroup
	wg.Add(len(fileHeaders))

	for index, header := range fileHeaders {
		go uploadImage(ctx, &wg, index, header, uploadResChan)
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
		default:
			fmt.Println("[UploadImages] No channel is ready")
		}
	}
}

func uploadImage(ctx echo.Context, wg *sync.WaitGroup, index int, header *multipart.FileHeader, uploadResChan chan UploadImagesRes) {
	defer wg.Done()
	// defer os.Remove(header.Filename)

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

	// Method1 : [File to Buffer] //
	//////////////
	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, multipartFile); err != nil {
		uploadResChan <- UploadImagesRes{
			index: index,
			url:   "",
			err:   err,
		}
		return
	}
	//////////////

	// Method2 : [Download File obj on Server] //
	//////////////
	// file, ok := multipartFile.(*os.File)
	// if !ok {
	// 	if err != nil {
	// 		var err error = errors.New("Could not convert multipart.File to os.File")
	// 		log.Fatal(err)
	// 		uploadResChan <- UploadImagesRes{
	// 			index: index,
	// 			url:   "",
	// 			err:   err,
	// 		}
	// 		return
	// 	}
	// }
	// defer file.Close()
	//////////////

	uploadToStorage(ctx, buf, header.Filename)

	url := header.Filename // arbitrary value for URL
	uploadResChan <- UploadImagesRes{
		index: index,
		url:   url,
		err:   nil,
	}
	return
}

// GCP Storage
func uploadToStorage(ctx echo.Context, buffer *bytes.Buffer, fileName string) (string, error) {
	c := ctx.Request().Context()
	client, err := storage.NewClient(c)
	if err != nil {
		panic(err)
	}
	defer client.Close()

	folder := "test"
	dst := folder + "/" + fileName

	storageWriter := client.Bucket(config.C.Storage.BucketName).Object(dst).NewWriter(c)

	r := bytes.NewReader(buffer.Bytes())
	if _, err = io.Copy(storageWriter, r); err != nil {
		return "", fmt.Errorf("Upload Failed > io.Copy: %v", err)
	}
	if err := storageWriter.Close(); err != nil {
		return "", fmt.Errorf("Upload Failed > Writer.Close: %v", err)
	}
	return "", nil
}

func processUploadedImgs(responses []UploadImagesRes) []string {
	rawRes := []UploadImagesRes{}
	urls := []string{}

	// logging error
	for _, res := range responses {
		if res.err != nil {
			log.Fatalf("%s th file error occured : %s", string(res.index), res.err.Error())
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
