package usecase

import (
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"mime/multipart"
	"os"
	"strings"

	"github.com/terry960302/go-clean-media-upload-server/adapter/repository"
	"github.com/terry960302/go-clean-media-upload-server/domain"
)

const prefix = "https://storage.googleapis.com/"

type ImageMetadataUsecase struct {
	domain.ImageMetadataUsecase
	imageRepo repository.ImageMetadataRepository
	mediaRepo repository.MediaMetadataRepository
}

func NewImageMetadataUsecase(imageRepo repository.ImageMetadataRepository, mediaRepo repository.MediaMetadataRepository) *ImageMetadataUsecase {
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
		Url:       url,
		MediaType: "image",
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

	fmt.Printf("%s img metadata is created\n", fmt.Sprint(imgId))
	return nil
}

func (i *ImageMetadataUsecase) GetConfigAndSaveImgData(header *multipart.FileHeader, file *os.File, url string) error {

	imgConfig, err := getImageConfig(file, header)
	if err != nil {
		log.Fatal(err)
		return err
	}

	if err := i.SaveImageMedia(url, imgConfig); err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

func getImageConfig(file *os.File, header *multipart.FileHeader) (ImageConfig, error) {
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

	imgConfig, _, err := image.Decode(file)
	if err != nil {
		log.Fatal(err)
		return ImageConfig{}, err
	}
	var format string
	mime := header.Header["Content-Type"][0]
	format = strings.Split(mime, "/")[1]

	width := imgConfig.Bounds().Dx()
	height := imgConfig.Bounds().Dy()
	volume := fileInfo.Size() / 1024 // KB

	config := ImageConfig{
		Width:  width,
		Height: height,
		Format: format,
		Volume: volume,
	}

	fmt.Printf("[Completed] file config extracted: %v\n", config)

	return config, nil

}
