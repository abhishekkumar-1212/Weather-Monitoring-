package services

import (
    "encoding/json"
    "fmt"
    "net/http"
    "os"
    "weatherMonitoring/models"
)

type WeatherAPIResponse struct {
    Main struct {
        Temp     float64 `json:"temp"`
        Humidity int     `json:"humidity"`
    } `json:"main"`
    Wind struct {
        Speed float64 `json:"speed"`
    } `json:"wind"`
    Name string `json:"name"`
}

func FetchWeather(city string) (*models.Weather, error) {
    apiKey := os.Getenv("OPENWEATHER_API_KEY")
    url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric", city, apiKey)

    resp, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("error fetching weather data: %s", resp.Status)
    }

    var weatherAPIResp WeatherAPIResponse
    if err := json.NewDecoder(resp.Body).Decode(&weatherAPIResp); err != nil {
        return nil, err
    }

    return &models.Weather{
        Temp:      weatherAPIResp.Main.Temp,
        City:      weatherAPIResp.Name,
        Humidity:  weatherAPIResp.Main.Humidity,
        WindSpeed: weatherAPIResp.Wind.Speed,
    }, nil
}
