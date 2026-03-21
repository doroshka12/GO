package config

import (
    "encoding/json"
    "errors"
    "os"
)

// StorageType тип хранилища
type StorageType string

const (
    FileStorage StorageType = "file"
    SQLiteStorage StorageType = "sqlite"
)

// Config структура конфигурации
type Config struct {
    Storage      StorageType `json:"storage"`
    DatabasePath string      `json:"database_path,omitempty"`
    FilePath     string      `json:"file_path,omitempty"`
    Latitude     float64     `json:"latitude"`
    Longitude    float64     `json:"longitude"`
    LogLevel     string      `json:"log_level"` // debug, info, error
}

// LoadConfig загружает конфигурацию из файла
func LoadConfig(path string) (*Config, error) {
    data, err := os.ReadFile(path)
    if err != nil {
        return nil, err
    }
    
    var cfg Config
    if err := json.Unmarshal(data, &cfg); err != nil {
        return nil, err
    }
    
    return &cfg, nil
}

// SaveConfig сохраняет конфигурацию
func SaveConfig(path string, cfg *Config) error {
    data, err := json.MarshalIndent(cfg, "", "  ")
    if err != nil {
        return err
    }
    return os.WriteFile(path, data, 0644)
}

// Validate проверяет конфигурацию
func (c *Config) Validate() error {
    if c.Latitude < -90 || c.Latitude > 90 {
        return errors.New("invalid latitude")
    }
    if c.Longitude < -180 || c.Longitude > 180 {
        return errors.New("invalid longitude")
    }
    return nil
}