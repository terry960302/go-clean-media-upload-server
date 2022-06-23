package controller

type mediaCtrl struct {
}

type MediaController interface {
	GetMedia() error
	CreateMedia() error
}

func testFunc() {
	// var mediaService service.MediaServiceImpl = service.MediaServiceImpl(s)

	// mediaService.GetAllMedia()

}
