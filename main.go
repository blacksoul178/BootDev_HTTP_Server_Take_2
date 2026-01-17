package main

import (
	"HTTP_Server_2/internal/config"
	"HTTP_Server_2/internal/logger"
	"log"
	"net/http"
)

func main() {
	// 1- Load app configs
	appConfig, err := config.LoadAppConfig("config.json")
	if err != nil {
		log.Fatalf("Error loading application configuration: %v", err)
	}

	// 2- Init logger
	err = logger.InitLogger(appConfig)
	if err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}

	const port = "8080"
	const filepathRoot = "."

	mux := http.NewServeMux()

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	mux.Handle("/", http.FileServer(http.Dir(filepathRoot)))

	log.Printf("Serving on port: %s\n", port)
	log.Fatal(srv.ListenAndServe())

}
