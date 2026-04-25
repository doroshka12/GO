package cli

import (
    "fmt"

    "github.com/doroshka12/GO/weather-app/internal/domain/models"
    "github.com/doroshka12/GO/weather-app/pkg/config"
)

// Logger интерфейс логгера
type Logger interface {
    Info(string)
    Debug(string)
    Error(string, error)
}

// WeatherInfo интерфейс для получения информации о погоде (теперь возвращает ошибку)
type WeatherInfo interface {
    GetTemperature(lat, long float64) (models.TempInfo, error)
}

// cliApp структура CLI приложения
type cliApp struct {
    l    Logger
    wi   WeatherInfo
    cfg  config.Config
}

// New создает новое CLI приложение
func New(l Logger, wi WeatherInfo, cfg config.Config) *cliApp {
    return &cliApp{
        l:    l,
        wi:   wi,
        cfg:  cfg,
    }
}

// Run запускает приложение
func (c *cliApp) Run() error {
    tempInfo, err := c.wi.GetTemperature(c.cfg.L.Lat, c.cfg.L.Long)
    if err != nil {
        c.l.Error("can't get temp info", err)
        return err
    }

    fmt.Printf(
        "Температура воздуха - %.2f градусов цельсия\n",
        tempInfo.Temp,
    )
    return nil
}
