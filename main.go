package main

import (
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	http.Handle("/", http.FileServer(http.Dir("./site")))
	http.ListenAndServe(":"+port, nil)
}
