package flags

import "flag"

// Flags структура хранения флагов
type Flags struct {
    Path string
}

// Parse парсит аргументы командной строки
func Parse() *Flags {
    config := flag.String("config", "./config/config.yaml", "path to config")
    flag.Parse()

    return &Flags{
        Path: *config,
    }
}
