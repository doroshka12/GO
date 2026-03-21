package weather

import (
    "encoding/json"
    "errors"
    "fmt"
    "io"
    "net/http"
    "time"
)

// WeatherData структура данных о погоде
type WeatherData struct {
    Temperature float64   `json:"temperature"`
    Latitude    float64   `json:"latitude"`
    Longitude   float64   `json:"longitude"`
    Time        time.Time `json:"time"`
}

// OpenMeteoResponse ответ от API
type OpenMeteoResponse struct {
    Current struct {
        Temperature2m float64 `json:"temperature_2m"`
        Time         string  `json:"time"`
    } `json:"current"`
    Latitude  float64 `json:"latitude"`
    Longitude float64 `json:"longitude"`
}

// WeatherService интерфейс сервиса погоды
type WeatherService interface {
    GetWeather(lat, lon float64) (*WeatherData, error)
}

// OpenMeteoService реализация через Open-Meteo
type OpenMeteoService struct {
    client *http.Client
}

func NewOpenMeteoService() *OpenMeteoService {
    return &OpenMeteoService{
        client: &http.Client{Timeout: 10 * time.Second},
    }
}

func (s *OpenMeteoService) GetWeather(lat, lon float64) (*WeatherData, error) {
    params := fmt.Sprintf("latitude=%f&longitude=%f&current=temperature_2m", lat, lon)
    url := fmt.Sprintf("https://api.open-meteo.com/v1/forecast?%s", params)
    
    resp, err := s.client.Get(url)
    if err != nil {
        return nil, errors.Join(errors.New("failed to get weather data"), err)
    }
    defer resp.Body.Close()
    
    data, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil, errors.Join(errors.New("failed to read response"), err)
    }
    
    var openMeteoResp OpenMeteoResponse
    if err := json.Unmarshal(data, &openMeteoResp); err != nil {
        return nil, errors.Join(errors.New("failed to parse JSON"), err)
    }
    
    return &WeatherData{
        Temperature: openMeteoResp.Current.Temperature2m,
        Latitude:    openMeteoResp.Latitude,
        Longitude:   openMeteoResp.Longitude,
        Time:        time.Now(),
    }, nil
}