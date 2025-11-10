package main

import (
	"log"
	"net/http"

	"github.com/chunyukuo88/rainy-day/handlers"
	"github.com/chunyukuo88/rainy-day/logger"
)

func main() {
	logInstance := initializeLogger()
	http.Handle("/", http.FileServer(http.Dir("public")))

	movieHandler := handlers.MovieHandler{}
	http.HandleFunc("/api/movies/top", movieHandler.GetTopMovies)

	const address = ":8080"

	err := http.ListenAndServe(address, nil)

	if err != nil {
		log.Fatalf("Server failed: %v", err)
		logInstance.Error("Server failed", err)
		return
	}

	logInstance.Info("Server is available")
}

func initializeLogger() *logger.Logger {
	logInstance, err := logger.NewLogger("rainy-day.log")
	if err != nil {
		log.Fatalf("Failed to initialize the logger: %v", err)
	}
	defer logInstance.Close()
	return logInstance
}
