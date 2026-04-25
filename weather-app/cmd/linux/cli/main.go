package main

import (
    "os"

    pogodaby "github.com/doroshka12/GO/weather-app/internal/adapters/pogoda_by"
    "github.com/doroshka12/GO/weather-app/internal/adapters/weather"
    "github.com/doroshka12/GO/weather-app/internal/pkg/app/cli"
    "github.com/doroshka12/GO/weather-app/internal/pkg/flags"
    "github.com/doroshka12/GO/weather-app/pkg/config"
    "github.com/doroshka12/GO/weather-app/pkg/logger"
)

func main() {
    // Парсим аргументы командной строки
    arguments := flags.Parse()

    // Открываем конфигурационный файл
    r, err := os.Open(arguments.Path)
    if err != nil {
        panic(err)
    }
    defer r.Close()

    // Парсим конфигурацию
    cfg, err := config.Parse(r)
    if err != nil {
        panic(err)
    }

    // Создаем логгер
    l := logger.New()

    // Получаем провайдер погоды в зависимости от конфигурации
    wi := getProvider(cfg, l)

    // Создаем приложение
    app := cli.New(l, wi, cfg)

    // Запускаем приложение
    err = app.Run()
    if err != nil {
        l.Error("Some error", err)
        os.Exit(1)
    }

    os.Exit(0)
}

// getProvider возвращает реализацию интерфейса WeatherInfo
// в зависимости от типа провайдера в конфигурации
func getProvider(cfg config.Config, l cli.Logger) cli.WeatherInfo {
    var wi cli.WeatherInfo

    switch cfg.P.Type {
    case "open-meteo":
        wi = weather.New(l)
    case "pogoda":
        wi = pogodaby.New(l)
    default:
        wi = weather.New(l)
    }

    return wi
}
