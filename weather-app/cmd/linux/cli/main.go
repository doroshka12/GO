package main

import (
    "os"

    "github.com/doroshka12/GO/weather-app/internal/pkg/app/cli"
    "github.com/doroshka12/GO/weather-app/pkg/logger"
)

func main() {
    // Создаем логгер
    l := logger.New()

    // Создаем приложение с логгером
    app := cli.New(l)

    // Запускаем приложение
    err := app.Run()
    if err != nil {
        l.Error("Some error", err)
        os.Exit(1)
    }

    os.Exit(0)
}