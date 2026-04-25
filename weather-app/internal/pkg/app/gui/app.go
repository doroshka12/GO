package gui

import (
    "github.com/doroshka12/GO/weather-app/internal/domain/gui_settings"
    "github.com/doroshka12/GO/weather-app/internal/domain/models"
    "github.com/doroshka12/GO/weather-app/pkg/config"
)

// Logger интерфейс логгера
type Logger interface {
    Info(string)
    Debug(string)
    Error(string, error)
}

// WeatherInfo интерфейс для получения погоды
type WeatherInfo interface {
    GetTemperature(lat, long float64) (models.TempInfo, error)
}

// GUI приложение
type App struct {
    l        Logger
    provider guisettings.Provider
    wi       WeatherInfo
    cfg      config.Config
}

// New создает новое GUI приложение
func New(l Logger, p guisettings.Provider, wi WeatherInfo, cfg config.Config) *App {
    return &App{
        l:        l,
        provider: p,
        wi:       wi,
        cfg:      cfg,
    }
}

// Run запускает GUI приложение
func (a *App) Run() error {
    // Создаем окно
    windowSize := guisettings.NewWS(400, 300)
    win, err := a.provider.CreateWindow("Weather App", windowSize)
    if err != nil {
        return err
    }

    // Создаем текстовый виджет
    tw := a.provider.GetTextWidget("Загрузка...")
    win.SetTemperatureWidget(tw)

    // Получаем погоду
    tempInfo, err := a.wi.GetTemperature(a.cfg.L.Lat, a.cfg.L.Long)
    if err != nil {
        a.l.Error("Failed to get weather", err)
        tw.SetText("Ошибка получения данных о погоде")
    } else {
        win.UpdateTemperature(tempInfo.Temp)
    }

    win.Render()

    // Запускаем приложение
    runner := a.provider.GetAppRunner()
    runner.Run()

    return nil
}
