package main

import (
	"net/http"

	"github.com/kiryshaaaa/infotecs-app/internal/transport"
)

func main() {
	//service := transport.NewService()
	router := transport.NewRouter()
	http.ListenAndServe(":8080", router.Router)
}
