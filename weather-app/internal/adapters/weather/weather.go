package weather

iimport (
    "encoding/json"
    "errors"
    "fmt"
    "io"
    "net/http"

    "github.com/doroshka12/GO/weather-app/internal/domain/models"
)

const apiURL = "https://api.open-meteo.com/v1/forecast"

// Logger интерфейс логгера
type Logger interface {
    Info(string)
    Debug(string)
    Error(string, error)
}

// current структура текущей погоды из ответа API
type current struct {
    Temp float32 `json:"temperature_2m"`
}

// response структура ответа от API
type response struct {
    Curr current `json:"current"`
}

// WeatherInfo структура для получения информации о погоде
type WeatherInfo struct {
    c        current
    l        Logger
    isLoaded bool
}

// New создает новый экземпляр WeatherInfo
func New(l Logger) *WeatherInfo {
    return &WeatherInfo{
        l: l,
    }
}

// getWeatherInfo получает данные о погоде из API
func (wi *WeatherInfo) getWeatherInfo(lat, long float64) error {
    var respData response

    params := fmt.Sprintf(
        "latitude=%f&longitude=%f&current=temperature_2m",
        lat, long,
    )
    url := fmt.Sprintf("%s?%s", apiURL, params)

    wi.l.Debug(fmt.Sprintf("url was generated success - %s", url))

    resp, err := http.Get(url)
    if err != nil {
        wi.l.Error("can't get weather data", err)
        customErr := errors.New("can't get weather data from openmeteo")
        return errors.Join(customErr, err)
    }
    defer func() {
        if err := resp.Body.Close(); err != nil {
            wi.l.Error("can't close body", err)
        }
    }()

    data, err := io.ReadAll(resp.Body)
    if err != nil {
        wi.l.Error("can't read data from body", err)
        customErr := errors.New("can't read data from response")
        return errors.Join(customErr, err)
    }

    wi.l.Debug(fmt.Sprintf("data was readed successfully size - %d", len(data)))

    if err := json.Unmarshal(data, &respData); err != nil {
        wi.l.Error("can't unmarshal json data", err)
        customErr := errors.New("can't unmarshal data from response")
        return errors.Join(customErr, err)
    }

    wi.c = respData.Curr
    wi.isLoaded = true
    return nil
}

// GetTemperature возвращает температуру для указанных координат
func (wi *WeatherInfo) GetTemperature(lat, long float64) models.TempInfo {
    if !wi.isLoaded {
        wi.getWeatherInfo(lat, long)
    }
    return models.TempInfo{
        Temp: wi.c.Temp,
    }
}