package handlers

import (
    "encoding/json"
    "net/http"
    "weatherMonitoring/services"
)

func GetWeather(w http.ResponseWriter, r *http.Request) {
    city := r.URL.Query().Get("city")
    if city == "" {
        http.Error(w, "City parameter is missing", http.StatusBadRequest)
        return
    }

    weatherData, err := services.FetchWeather(city)
    if err != nil {
        http.Error(w, "Error fetching weather data", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(weatherData)
}
