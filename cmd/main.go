package main

import (
	"fmt"
	"groupie-tracker/internal/handlers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handlers.HomePage)
	http.HandleFunc("/artist", handlers.ArtistPage)
	fileServer := http.FileServer(http.Dir("ui/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))
	fmt.Println("Server is running on :http://localhost:8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
