package pogodaby

import (
    "encoding/json"
    "net/http"

    "github.com/doroshka12/GO/weather-app/internal/domain/models"
)

const url = "https://pogoda.by/api/v2/weather-fact?station=26820"

// response структура ответа от pogoda.by
type response struct {
    Temp float32 `json:"t"`
}

// Logger интерфейс логгера
type Logger interface {
    Info(string)
    Debug(string)
    Error(string, error)
}

// Pogoda структура для получения погоды с pogoda.by
type Pogoda struct {
    l Logger
}

// New создает новый экземпляр Pogoda
func New(l Logger) *Pogoda {
    return &Pogoda{l: l}
}

// GetTemperature возвращает температуру из pogoda.by
func (p *Pogoda) GetTemperature(lat, long float64) (models.TempInfo, error) {
    p.l.Debug("Getting weather from pogoda.by...")
    
    response, err := http.Get(url)
    if err != nil {
        p.l.Error("can't get data from pogoda.by", err)
        return models.TempInfo{}, err
    }
    defer func() {
        err := response.Body.Close()
        if err != nil {
            p.l.Error("can't close response body", err)
        }
    }()

    var resp response
    if err := json.NewDecoder(response.Body).Decode(&resp); err != nil {
        p.l.Error("can't decode JSON", err)
        return models.TempInfo{}, err
    }

    return models.TempInfo{Temp: resp.Temp}, nil
}
