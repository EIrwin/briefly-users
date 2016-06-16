package main

import (
	"github.com/eirwin/briefly-users/api"
	"log"
	"net/http"
)

func main() {
	log.Print("starting user service...")
	router := api.NewRouter()
	log.Fatal(http.ListenAndServe(":8181", router))
}
