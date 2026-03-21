package main

import (
    "os"

    "github.com/doroshka12/GO/weather-app/internal/adapters/weather"
    "github.com/doroshka12/GO/weather-app/internal/pkg/app/cli"
    "github.com/doroshka12/GO/weather-app/pkg/logger"
)

func main() {
    // Создаем логгер
    l := logger.New()

    // Создаем адаптер для погоды
    wi := weather.New(l)

    // Создаем приложение с логгером и погодным адаптером
    app := cli.New(l, wi)

    // Запускаем приложение
    err := app.Run()
    if err != nil {
        l.Error("Some error", err)
        os.Exit(1)
    }

    os.Exit(0)
}