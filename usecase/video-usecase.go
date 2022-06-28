package usecase

import (
	"fmt"
	"log"
	"mime/multipart"
	"os"
	"strings"

	"github.com/terry960302/go-clean-media-upload-server/adapter/repository"
	"github.com/terry960302/go-clean-media-upload-server/domain"
)

type VideoMetadataUsecase struct {
	domain.VideoMetadataUsecase
	videoRepo repository.VideoMetadataRepository
	mediaRepo repository.MediaMetadataRepository
}

func NewVideoMetadataUsecase(videoRepo repository.VideoMetadataRepository, mediaRepo repository.MediaMetadataRepository) *VideoMetadataUsecase {
	usecase := &VideoMetadataUsecase{
		videoRepo: videoRepo,
		mediaRepo: mediaRepo,
	}
	usecase.VideoMetadataUsecase = interface{}(usecase).(*VideoMetadataUsecase)
	return usecase
}

type VideoConfig struct {
	Volume int64
	Format string
}

func (v *VideoMetadataUsecase) GetConfigAndSaveVideoData(header *multipart.FileHeader, file *os.File, url string) error {
	videoConfig, err := getVideoConfig(file, header)
	if err != nil {
		log.Fatal(err)
		return err
	}

	if err := v.SaveVideoMedia(url, videoConfig); err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

func (v *VideoMetadataUsecase) SaveVideoMedia(url string, config VideoConfig) error {

	media := domain.MediaMetadata{
		Url:       url,
		MediaType: "video",
	}
	mediaId, err := v.mediaRepo.Create(media)
	if err != nil {
		log.Fatal(err)
		return err
	}

	image := domain.VideoMetadata{
		MediaID: mediaId,
		Volume:  fmt.Sprint(config.Volume),
		Format:  config.Format,
	}
	imgId, err := v.videoRepo.Create(image)
	if err != nil {
		log.Fatal(err)
		return err
	}

	fmt.Printf("%s img metadata is created\n", fmt.Sprint(imgId))
	return nil
}

func getVideoConfig(file *os.File, header *multipart.FileHeader) (VideoConfig, error) {
	// File read index could set to end because of previous functions related to os.File
	// Set file read index to begining before run 'Decode' function
	if _, err := file.Seek(0, 0); err != nil {
		log.Fatal(err)
		return VideoConfig{}, err
	}

	fileInfo, err := file.Stat()
	if err != nil {
		log.Fatal(err)
		return VideoConfig{}, err
	}

	if err != nil {
		log.Fatal(err)
		return VideoConfig{}, err
	}

	volume := fileInfo.Size() / 1024 // KB
	var format string
	mime := header.Header["Content-Type"][0]
	format = strings.Split(mime, "/")[1]

	config := VideoConfig{

		Format: format,
		Volume: volume,
	}

	fmt.Printf("[Completed] file config extracted: %v\n", config)

	return config, nil

}
