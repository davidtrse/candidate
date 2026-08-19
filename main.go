package main

import (
	"log"
	"net/http"

	"goassessment/internal/handler"
)

func main() {
	http.HandleFunc("/users", handler.GetUserHandler)

	log.Println("listening on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
