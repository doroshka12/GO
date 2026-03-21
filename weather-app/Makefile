.PHONY: run build clean test

# Переменные
BINARY_NAME=weather-cli
BUILD_DIR=build

# Запуск CLI приложения
run:
	go run ./cmd/linux/cli/main.go

# Сборка бинарника
build:
	mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(BINARY_NAME) ./cmd/linux/cli/main.go

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

# Линтер
lint:
	golangci-lint run

# Запуск с отладкой
run-debug:
	go run ./cmd/linux/cli/main.go -debug

run:
	go run ./cmd/linux/cli/main.go

run-debug:
	DEBUG=true go run ./cmd/linux/cli/main.go

build:
	mkdir -p build
	go build -o build/weather-cli ./cmd/linux/cli/main.go