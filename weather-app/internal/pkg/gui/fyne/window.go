package fyne

import (
    "fmt"

    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/widget"

    guisettings "github.com/doroshka12/GO/weather-app/internal/domain/gui_settings"
)

type window struct {
    w   fyne.Window
    tw  guisettings.TextWidget
}

func NewW(w fyne.Window) *window {
    return &window{w: w}
}

func (win *window) Resize(ws guisettings.WindowSize) error {
    if !ws.IsFull() {
        win.w.Resize(fyne.NewSize(float32(ws.Width()), float32(ws.Height())))
    }
    return nil
}

func (win *window) UpdateTemperature(t float32) error {
    if win.tw == nil {
        return nil
    }
    win.tw.SetText(fmt.Sprintf("Температура: %.1f°C", t))
    return nil
}

func (win *window) SetTemperatureWidget(tw guisettings.TextWidget) error {
    win.tw = tw
    win.w.SetContent(tw.Render().(*widget.Label))
    return nil
}

func (win *window) Render() error {
    return nil
}
