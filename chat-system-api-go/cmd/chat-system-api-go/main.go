package main

import (
	"log"
	"net/http"
	"github.com/wael-anwar/chat-system-api-go/internal/router"
)

func main() {
	r := router.InitRouter()

	log.Println("Listening on 5000 ......")
	log.Fatal(http.ListenAndServe(":5000", r))
}
