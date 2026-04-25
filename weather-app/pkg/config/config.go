package config

import (
    "io"

    "gopkg.in/yaml.v3"
)

// Provider структура провайдера погоды
type Provider struct {
    Type string `yaml:"type"`
}

// Location структура местоположения
type Location struct {
    Lat  float64 `yaml:"lat"`
    Long float64 `yaml:"long"`
}

// Config структура конфигурации
type Config struct {
    P Provider `yaml:"provider"`
    L Location `yaml:"location"`
}

// ConfigFile обертка для парсинга YAML
type ConfigFile struct {
    C Config `yaml:"service"`
}

// Parse парсит конфигурацию из io.Reader
func Parse(r io.Reader) (Config, error) {
    var c ConfigFile
    if err := yaml.NewDecoder(r).Decode(&c); err != nil {
        return Config{}, err
    }
    return c.C, nil
}
