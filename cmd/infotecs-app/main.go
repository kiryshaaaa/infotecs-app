package main

import (
	"fmt"
	"net/http"

	"github.com/kiryshaaaa/infotecs-app/internal/routers"
	"github.com/kiryshaaaa/infotecs-app/internal/service"
)

func main() {
	service := service.NewService()
	router := routers.NewRouter(service)
	router.GetHelloWorld()
	fmt.Println("Hello World")
	http.ListenAndServe(":8080", router.Router)
}
