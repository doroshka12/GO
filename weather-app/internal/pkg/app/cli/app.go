package cli

import (
    "fmt"

    "github.com/doroshka12/GO/weather-app/internal/domain/models"
)

// Logger интерфейс логгера
type Logger interface {
    Info(string)
    Debug(string)
    Error(string, error)
}

// WeatherInfo интерфейс для получения информации о погоде
type WeatherInfo interface {
    GetTemperature(float64, float64) models.TempInfo
}

// cliApp структура CLI приложения
type cliApp struct {
    l  Logger
    wi WeatherInfo
}

// New создает новое CLI приложение
func New(l Logger, wi WeatherInfo) *cliApp {
    return &cliApp{
        l:  l,
        wi: wi,
    }
}

// Run запускает приложение
func (c *cliApp) Run() error {
    // Координаты Гродно (можно потом вынести в конфиг)
    latitude := 53.6688
    longitude := 23.8223

    c.l.Debug("Getting weather information...")
    
    tempInfo := c.wi.GetTemperature(latitude, longitude)
    
    fmt.Printf(
        "Температура воздуха - %.2f градусов цельсия\n",
        tempInfo.Temp,
    )
    return nil
}