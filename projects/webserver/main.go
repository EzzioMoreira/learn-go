package main

import (
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("/tmp"))
	log.Println("Serving /tmp on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", fs))
}
