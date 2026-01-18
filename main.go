package main

import (
	"HTTP_Server_2/internal/config"
	"HTTP_Server_2/internal/handlers"
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

	mux := http.NewServeMux()
	handlers.App(mux)

	srv := &http.Server{
		Addr:    ":" + appConfig.Server.Port,
		Handler: mux,
	}

	mux.HandleFunc("/healthz", handlers.Healthz)

	log.Printf("Serving on port: %s\n", appConfig.Server.Port)
	log.Fatal(srv.ListenAndServe())

}
