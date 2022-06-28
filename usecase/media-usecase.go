package usecase

import (
	"errors"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"os"
	"sort"
	"strings"
	"sync"

	"cloud.google.com/go/storage"
	"github.com/labstack/echo/v4"
	"github.com/terry960302/go-clean-media-upload-server/adapter/repository"
	"github.com/terry960302/go-clean-media-upload-server/config"
	"github.com/terry960302/go-clean-media-upload-server/domain"
	"google.golang.org/api/option"
)

type MediaMetadataUsecase struct {
	domain.MediaMetadataUsecase
	mediaRepo    repository.MediaMetadataRepository
	imgUsecase   ImageMetadataUsecase
	videoUsecase VideoMetadataUsecase
}

func NewMediaMetadataUsecase(mediaRepo repository.MediaMetadataRepository, imgUsecase ImageMetadataUsecase, videoUsecase VideoMetadataUsecase) *MediaMetadataUsecase {
	usecase := &MediaMetadataUsecase{mediaRepo: mediaRepo, imgUsecase: imgUsecase, videoUsecase: videoUsecase}
	usecase.MediaMetadataUsecase = interface{}(usecase).(*MediaMetadataUsecase)
	return usecase
}

type UploadFilesRes struct {
	index int
	url   string
	err   error
}

func (m *MediaMetadataUsecase) UploadFiles(fileHeaders []*multipart.FileHeader, ctx echo.Context) []string {

	uploadResChan := make(chan UploadFilesRes)
	var wg sync.WaitGroup
	wg.Add(len(fileHeaders))

	for index, header := range fileHeaders {
		var mediaType string
		mime := strings.Split(header.Header["Content-Type"][0], "/")
		mediaType = mime[0]

		if mediaType == "image" {
			go m.UploadFile(ctx, &wg, index, header, uploadResChan, m.imgUsecase.GetConfigAndSaveImgData)
		} else {
			go m.UploadFile(ctx, &wg, index, header, uploadResChan, m.videoUsecase.GetConfigAndSaveVideoData)
		}
	}

	responses := []UploadFilesRes{}

	go func() {
		wg.Wait()
		close(uploadResChan)
	}()

	for {
		select {
		case res := <-uploadResChan:
			responses = append(responses, res)

			if len(responses) == len(fileHeaders) {
				urls := processUploadedFiless(responses)
				return urls
			}
			// default:
			// 	fmt.Println("[UploadImages] No channel is ready")
		}
	}
}

func (m *MediaMetadataUsecase) UploadFile(ctx echo.Context, wg *sync.WaitGroup, index int, header *multipart.FileHeader, uploadResChan chan UploadFilesRes, getConfigAndSaveData func(header *multipart.FileHeader, file *os.File, url string) error) {
	defer wg.Done()

	multipartFile, err := header.Open()
	if err != nil {
		log.Fatal(err)
		uploadResChan <- UploadFilesRes{
			index: index,
			url:   "",
			err:   err,
		}
		return
	}

	defer multipartFile.Close()

	file, err := os.Create(header.Filename)
	if err != nil {
		log.Fatal(err)
		uploadResChan <- UploadFilesRes{
			index: index,
			url:   "",
			err:   err,
		}
		return
	}
	defer file.Close()

	if _, err = io.Copy(file, multipartFile); err != nil {
		log.Fatal(err)
		uploadResChan <- UploadFilesRes{
			index: index,
			url:   "",
			err:   err,
		}
		return
	}
	defer func() {
		path := header.Filename
		if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
			return
		}
		os.Remove(path)
		return
	}()

	url, err := uploadToStorage(ctx, file)
	if err != nil {
		log.Fatal(err)
		uploadResChan <- UploadFilesRes{
			index: index,
			url:   "",
			err:   err,
		}
		return

	}

	if err := getConfigAndSaveData(header, file, url); err != nil {
		log.Fatal(err)
		uploadResChan <- UploadFilesRes{
			index: index,
			url:   "",
			err:   err,
		}
		return
	}

	uploadResChan <- UploadFilesRes{
		index: index,
		url:   url,
		err:   nil,
	}
	return
}

// Uplaod to GCP Storage
func uploadToStorage(ctx echo.Context, file *os.File) (string, error) {
	// File read index could set to end because of previous functions related to os.File
	// Set file read index to begining before run 'Decode' function
	if _, err := file.Seek(0, 0); err != nil {
		log.Fatal(err)
		return "", err
	}

	c := ctx.Request().Context()
	client, err := storage.NewClient(c, option.WithCredentialsFile((config.C.Storage.CredPath)))
	if err != nil {
		panic(err)
	}
	defer client.Close()

	dir := "test"
	newFilename := file.Name()
	dst := dir + "/" + newFilename

	storageWriter := client.Bucket(config.C.Storage.BucketName).Object(dst).NewWriter(c)

	if _, err = io.Copy(storageWriter, file); err != nil {
		return "", fmt.Errorf("Upload Failed > io.Copy: %v", err)
	}
	if err := storageWriter.Close(); err != nil {
		return "", fmt.Errorf("Upload Failed > Writer.Close: %v", err)
	}

	url := prefix + config.C.Storage.BucketName + "/" + dst
	fmt.Printf("[Completed] storage image uploaded : %s\n", file.Name())
	return url, nil
}

func processUploadedFiless(responses []UploadFilesRes) []string {
	rawRes := []UploadFilesRes{}
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

	fmt.Println("[Completed] process images output")

	return urls
}
