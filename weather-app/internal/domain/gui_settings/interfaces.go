package guisettings

// TextWidget интерфейс для текстового виджета
type TextWidget interface {
    Render() any
    SetText(text string)
}

// Window интерфейс для окна приложения
type Window interface {
    Resize(ws WindowSize) error
    UpdateTemperature(t float32) error
    SetTemperatureWidget(tw TextWidget) error
    Render() error
}

// AppRunner интерфейс для запуска приложения
type AppRunner interface {
    Run()
}

// Provider интерфейс для создания GUI компонентов
type Provider interface {
    CreateWindow(name string, size WindowSize) (Window, error)
    GetAppRunner() AppRunner
    GetTextWidget(text string) TextWidget
}
