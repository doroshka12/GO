package main

import (
    "context"
    "database/sql"
    "fmt"
    "log"
    "os"
    "time"

    _ "github.com/mattn/go-sqlite3"
)

func main() {
    // Удаляем старую БД если есть
    os.Remove("test.db")
    
    // Открываем БД
    db, err := sql.Open("sqlite3", "test.db")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()
    
    // Создаем таблицу
    _, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS users (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            name TEXT NOT NULL,
            email TEXT NOT NULL UNIQUE,
            created_at DATETIME DEFAULT CURRENT_TIMESTAMP
        )
    `)
    if err != nil {
        log.Fatal(err)
    }
    
    // Создаем репозиторий
    repo := NewSQLiteRepository(db)
    ctx := context.Background()
    
    // Демонстрация работы
    fmt.Println("🚀 Демонстрация работы с SQLite...")
    
    // Создаем пользователей
    users := []*User{
        {Name: "Иван Петров", Email: "ivan@mail.com"},
        {Name: "Мария Сидорова", Email: "maria@mail.com"},
    }
    
    for _, user := range users {
        err := repo.CreateUser(ctx, user)
        if err != nil {
            log.Printf("Ошибка создания %s: %v", user.Name, err)
            continue
        }
        fmt.Printf("✅ Создан пользователь: %s (ID: %d)\n", user.Name, user.ID)
    }
    
    // Получаем список пользователей
    allUsers, err := repo.ListUsers(ctx)
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Println("\n📋 Список пользователей:")
    for _, u := range allUsers {
        fmt.Printf("  ID: %d | %s | %s\n", u.ID, u.Name, u.Email)
    }
}