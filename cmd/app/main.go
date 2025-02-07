package main

import (
	"log"
	"net/http"
	"r01/internal/delivery/http/router"
)

func main() {
	r := router.NewRouter()
	log.Println("Server running on port 8080")
	http.ListenAndServe(":8080", r)
}
