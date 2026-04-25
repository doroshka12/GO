package providers

import (
    pogodaby "github.com/doroshka12/GO/weather-app/internal/adapters/pogoda_by"
    "github.com/doroshka12/GO/weather-app/internal/adapters/weather"
    "github.com/doroshka12/GO/weather-app/internal/pkg/app/cli"
    "github.com/doroshka12/GO/weather-app/pkg/config"
)

// GetProvider возвращает реализацию интерфейса WeatherInfo
func GetProvider(cfg config.Config, l cli.Logger) cli.WeatherInfo {
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
