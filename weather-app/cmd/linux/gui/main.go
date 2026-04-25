package main

import (
    "os"

    "github.com/doroshka12/GO/weather-app/internal/pkg/app/gui"
    "github.com/doroshka12/GO/weather-app/internal/pkg/config"
    "github.com/doroshka12/GO/weather-app/internal/pkg/flags"
    fyneProvider "github.com/doroshka12/GO/weather-app/internal/pkg/gui/fyne"
    "github.com/doroshka12/GO/weather-app/internal/pkg/providers"
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

    // Получаем провайдер погоды
    provider := providers.GetProvider(cfg, l)

    // Создаем GUI провайдер (Fyne)
    guiProvider := fyneProvider.NewP()

    // Создаем GUI приложение
    app := gui.New(l, guiProvider, provider, cfg)

    // Запускаем приложение
    err = app.Run()
    if err != nil {
        panic(err)
    }
}
