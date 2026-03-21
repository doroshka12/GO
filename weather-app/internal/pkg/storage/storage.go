package storage

import (
    "encoding/json"
    "errors"
    "os"
)

// Location структура местоположения
type Location struct {
    Latitude  float64 `json:"latitude"`
    Longitude float64 `json:"longitude"`
    City      string  `json:"city,omitempty"`
}

// Storage интерфейс хранилища
type Storage interface {
    GetLocation() (*Location, error)
    SaveLocation(loc *Location) error
}

// FileStorage хранилище в файле
type FileStorage struct {
    path string
}

func NewFileStorage(path string) *FileStorage {
    return &FileStorage{path: path}
}

func (s *FileStorage) GetLocation() (*Location, error) {
    data, err := os.ReadFile(s.path)
    if err != nil {
        return nil, err
    }
    
    var loc Location
    if err := json.Unmarshal(data, &loc); err != nil {
        return nil, err
    }
    return &loc, nil
}

func (s *FileStorage) SaveLocation(loc *Location) error {
    data, err := json.MarshalIndent(loc, "", "  ")
    if err != nil {
        return err
    }
    return os.WriteFile(s.path, data, 0644)
}

// SQLiteStorage хранилище в SQLite
type SQLiteStorage struct {
    dbPath string
}

func NewSQLiteStorage(dbPath string) *SQLiteStorage {
    return &SQLiteStorage{dbPath: dbPath}
}

func (s *SQLiteStorage) GetLocation() (*Location, error) {
    // Здесь будет код для SQLite
    return nil, errors.New("sqlite storage not implemented yet")
}

func (s *SQLiteStorage) SaveLocation(loc *Location) error {
    // Здесь будет код для SQLite
    return errors.New("sqlite storage not implemented yet")
}

// NewStorage фабрика для создания хранилища
func NewStorage(storageType string, path string) (Storage, error) {
    switch storageType {
    case "file":
        return NewFileStorage(path), nil
    case "sqlite":
        return NewSQLiteStorage(path), nil
    default:
        return nil, errors.New("unknown storage type")
    }
}