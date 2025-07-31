package main

import (
	"fmt"
	"net/http"

	"heintzz/notion-reminder/apps/notes"

	"github.com/go-chi/chi/v5"
)

func main() {
	router := chi.NewRouter()
	notes.Run(router)

	fmt.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		fmt.Printf("Error starting server: %v\n", err)
		return
	}
}
