package main

import (
    "log"
    "net/http"
    "weatherMonitoring/handlers"
)

func main() {
    // Serve static files from the frontend directory
    http.Handle("/", http.FileServer(http.Dir("../frontend")))

    // API endpoint for fetching weather data
    http.HandleFunc("/weather", handlers.GetWeather)

    log.Println("Starting server on port 8080...")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
