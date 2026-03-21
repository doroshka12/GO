package main

import (
    "database/sql"
    "fmt"
    "log"
    "net/http"
    "os"

    "github.com/doroshka12/GO/sqlite-example/internal/user"
    "github.com/gorilla/mux"
    _ "github.com/mattn/go-sqlite3"
)

func main() {
    // Инициализация базы данных
    db, err := initDB()
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    // Создаем репозиторий и обработчики
    repo := user.NewSQLiteRepository(db)
    handler := user.NewHandler(repo)

    // Настройка маршрутов
    router := mux.NewRouter()
    
    // API маршруты
    api := router.PathPrefix("/api").Subrouter()
    api.HandleFunc("/users", handler.CreateUser).Methods("POST")
    api.HandleFunc("/users", handler.ListUsers).Methods("GET")
    api.HandleFunc("/users/{id:[0-9]+}", handler.GetUser).Methods("GET")
    api.HandleFunc("/users/{id:[0-9]+}", handler.UpdateUser).Methods("PUT")
    api.HandleFunc("/users/{id:[0-9]+}", handler.DeleteUser).Methods("DELETE")
    
    // Корневой маршрут
    router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "text/html")
        fmt.Fprintf(w, `
            <h1>User API</h1>
            <ul>
                <li>POST /api/users - создать пользователя</li>
                <li>GET /api/users - список всех пользователей</li>
                <li>GET /api/users/{id} - получить пользователя</li>
                <li>PUT /api/users/{id} - обновить пользователя</li>
                <li>DELETE /api/users/{id} - удалить пользователя</li>
            </ul>
        `)
    })

    // Запуск сервера
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }
    
    log.Printf("🚀 Сервер запущен на http://localhost:%s", port)
    log.Fatal(http.ListenAndServe(":"+port, router))
}

// initDB инициализирует базу данных
func initDB() (*sql.DB, error) {
    // Удаляем старую БД для чистого старта (опционально)
    // os.Remove("test.db")
    
    db, err := sql.Open("sqlite3", "test.db")
    if err != nil {
        return nil, err
    }
    
    // Создаем таблицу если не существует
    _, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS users (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            name TEXT NOT NULL,
            email TEXT NOT NULL UNIQUE,
            created_at DATETIME DEFAULT CURRENT_TIMESTAMP
        )
    `)
    if err != nil {
        return nil, err
    }
    
    return db, nil
}