package main

import (
	"log"
	"net/http"
	"r01/internal/delivery/http/router"
	"r01/pkg/db"
)

func main() {
	db, err := db.ConnectPostgresDB()
	if err != nil {
		log.Fatal("Failed to connect to database", err)
	}
	defer db.Close()

	r := router.NewRouter(db)
	log.Println("Server running on port 8080")
	http.ListenAndServe(":8080", r)
}
