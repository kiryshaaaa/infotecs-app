package service

import (
	//"fmt"
	"net/http"
)

func (s *Service) MyHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}
