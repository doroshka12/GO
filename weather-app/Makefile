.PHONY: run run-gui build build-gui test clean

# Переменные
BINARY_NAME=weather-cli
GUI_BINARY_NAME=weather-gui
BUILD_DIR=build

# Запуск CLI приложения
run:
	go run ./cmd/linux/cli/main.go

# Запуск GUI приложения
run-gui:
	go run ./cmd/linux/gui/main.go

# Сборка CLI бинарника
build:
	mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(BINARY_NAME) ./cmd/linux/cli/main.go

# Сборка GUI бинарника
build-gui:
	mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(GUI_BINARY_NAME) ./cmd/linux/gui/main.go

# Тестирование
test:
	go test -v ./...

# Запуск с покрытием
test-coverage:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out

# Очистка
clean:
	rm -rf $(BUILD_DIR)
	rm -f coverage.out

# Установка зависимостей
deps:
	go mod tidy
	go mod download

# Форматирование кода
fmt:
	go fmt ./...
