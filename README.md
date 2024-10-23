# Weather-Monitoring-
Develop a real-time data processing system to monitor weather conditions and provide summarized insights using rollups and aggregates. The system will utilize data from the OpenWeatherMap API (https://openweathermap.org/).
# Objective
The aim of this project is to develop a real-time data processing system to monitor weather conditions and provide summarized insights using rollups and aggregates. The system retrieves weather data from the OpenWeatherMap API and processes it to deliver useful information such as daily weather summaries, alerts based on configurable thresholds, and visualizations of weather trends.

# Features
Real-Time Data Retrieval: Continuously fetches weather data from OpenWeatherMap for major Indian metros (Delhi, Mumbai, Chennai, Bangalore, Kolkata, Hyderabad).
Weather Summarization: Provides daily weather summaries, including average, maximum, minimum temperatures, and dominant weather condition.
Configurable Alerting System: Users can set thresholds for specific weather parameters to receive alerts when conditions are met.
Data Visualization: Displays historical weather trends, daily summaries, and triggered alerts.
Golang Backend: Efficient backend processing, leveraging the concurrency features of Golang.
Data Source
The system uses the OpenWeatherMap API to fetch weather data. Sign up here for a free API key to access the data.

# Weather Parameters Used:
main: Main weather condition (e.g., Rain, Snow, Clear)
temp: Current temperature in Celsius
feels_like: Perceived temperature in Celsius
dt: Time of the data update (Unix timestamp)
Processing and Analysis
The system retrieves real-time weather data at configurable intervals (e.g., every 5 minutes).
Temperature values from the API (in Kelvin) are converted to Celsius (and optionally, Fahrenheit) based on user preferences.
Rollups and Aggregates
Daily Weather Summary:

Roll up weather data to provide daily aggregates, including:
Average Temperature
Maximum Temperature
Minimum Temperature
Dominant Weather Condition: Determined by the most frequent condition during the day (e.g., if it's mostly sunny, "Clear" will be the dominant condition).
Alerting Thresholds:

Users can set thresholds for specific weather conditions (e.g., alert if the temperature exceeds 35°C for two consecutive updates).
Alerts are triggered when thresholds are breached and can be displayed on the console or sent via email.
# Visualization:

Provides visual representation of daily weather summaries, historical trends, and triggered alerts.
# System Architecture
The system follows a 3-tier architecture:

UI Layer: User interface for configuring settings and viewing data.
API Layer: Exposes endpoints for data retrieval, user configuration, and alerting.
Backend Layer: Processes data, manages API calls, and handles storage.
Setup Instructions
Prerequisites
Golang: Make sure Go is installed (version 1.18 or above).
API Key: Register and get an OpenWeatherMap API key.
Installation
Clone the Repository:

git clone https://github.com/yourusername/weather-data-processor.git
cd weather-data-processor
Install Dependencies:
API_KEY=your_openweathermap_api_key
INTERVAL=300  # Time interval in seconds (e.g., 5 minutes)
Run the Application:


go run main.go
# API Design
/fetch-weather

Description: Fetches real-time weather data for predefined metros.
Method: GET
/set-thresholds

Description: Allows users to set custom alert thresholds for weather conditions.
Method: POST
Input: JSON containing threshold settings (e.g., max temperature)
/get-summary

Description: Retrieves daily weather summaries.
Method: GET
Test Cases
System Setup:

Ensure the system connects to the OpenWeatherMap API using a valid API key.
Data Retrieval:

Simulate periodic API calls and verify data is correctly parsed.
Temperature Conversion:

Validate conversion from Kelvin to Celsius and Fahrenheit.
Daily Weather Summary:

Verify accurate calculation of daily aggregates including average, maximum, and minimum temperatures, and dominant weather condition.
Alerting Thresholds:

Simulate weather conditions that exceed user-configured thresholds to test alert triggers.
Sample Code
Main Weather Data Processing

func FetchWeatherData(apiKey string, city string) WeatherData {
    // Logic to call OpenWeatherMap API and retrieve data
}

func ConvertKelvinToCelsius(temp float64) float64 {
    return temp - 273.15
}

func MonitorWeather() {
    ticker := time.NewTicker(time.Duration(interval) * time.Second)
    for range ticker.C {
        // Fetch, process, and store weather data
    }
}
Bonus Features
Additional Weather Parameters: Support for other metrics like humidity, wind speed, etc.
Weather Forecast Summaries: Extend functionality to provide insights based on weather forecasts.
Notification Enhancements: Enable email alerts and push notifications for better user experience.
Evaluation Criteria
Functionality and Correctness: Ensure smooth, real-time data processing.
Data Parsing Accuracy: Verify correct temperature conversion and aggregate calculations.
Efficiency: Optimal data retrieval and processing without lag.
Test Case Coverage: Comprehensive testing for various weather scenarios and configurations.
Code Clarity and Maintenance: Clean, maintainable codebase.
(Bonus) Additional Features: Implementation of extra features.
Contributing
We welcome contributions! Please fork the repository, make your changes, and submit a pull request.

# License
This project is licensed under the MIT License. See LICENSE for more information.
