package ftp

import "net/http"

type PageController struct {
}

func NewPageController() *PageController {
	return &PageController{}
}

func (this *PageController) RegisterEndpoints() error {
	http.Handle("/", http.FileServer(http.Dir("./wwwroot/")))
	return nil
}
