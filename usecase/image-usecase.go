package usecase

import (
	"errors"
	"fmt"
	"log"
	"media-upload-server/adapter/repository"
	"media-upload-server/domain"
	"mime/multipart"
	"os"
	"sort"
	"sync"
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

func (i *ImageMetadataUsecase) UploadImages(fileHeaders []*multipart.FileHeader) []string {

	uploadResChan := make(chan UploadImagesRes)
	var wg sync.WaitGroup
	wg.Add(len(fileHeaders))

	for index, header := range fileHeaders {
		go uploadImage(&wg, index, header, uploadResChan)
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
			fmt.Println("No channel is ready")
		}
	}
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

func uploadImage(wg *sync.WaitGroup, index int, header *multipart.FileHeader, uploadResChan chan UploadImagesRes) {
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

	file, ok := multipartFile.(*os.File)
	if !ok {
		if err != nil {
			var err error = errors.New("Could not convert multipart.File to os.File")
			log.Fatal(err)
			uploadResChan <- UploadImagesRes{
				index: index,
				url:   "",
				err:   err,
			}
			return
		}
	}
	defer file.Close()

	// TODO : upload file to remote storage like GCP, AWS, AZURE... and fetch file path url

	url := file.Name() // arbitrary value for URL
	uploadResChan <- UploadImagesRes{
		index: index,
		url:   url,
		err:   nil,
	}
	return
}
