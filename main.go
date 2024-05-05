package main

import (
	application "generate-manifest/application"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/health-check", application.HealthCheck)
	http.HandleFunc("/create-manifest", application.GenerateManifest)

	log.Println("Server started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
