package main

import (
    "database/sql"
    "fmt"
    "log"
    "os"

    _ "github.com/mattn/go-sqlite3"
)

type User struct {
    ID    int
    Name  string
    Email string
}

func main() {
    // Удаляем старую БД если есть (для чистого старта)
    os.Remove("users.db")

    fmt.Println("🚀 Запуск SQLite примера...")
    fmt.Println("==========================")

    // Открываем (создаем) базу данных
    db, err := sql.Open("sqlite3", "users.db")
    if err != nil {
        log.Fatal("Ошибка открытия БД:", err)
    }
    defer db.Close()

    // Проверяем соединение
    err = db.Ping()
    if err != nil {
        log.Fatal("Ошибка соединения с БД:", err)
    }
    fmt.Println("✅ Соединение с БД установлено")

    // Создаем таблицу
    createTableSQL := `
    CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT NOT NULL,
        email TEXT NOT NULL UNIQUE,
        created_at DATETIME DEFAULT CURRENT_TIMESTAMP
    );`

    _, err = db.Exec(createTableSQL)
    if err != nil {
        log.Fatal("Ошибка создания таблицы:", err)
    }
    fmt.Println("✅ Таблица users создана успешно")

    // Вставляем данные
    fmt.Println("\n📝 Добавление пользователей...")
    users := []User{
        {Name: "Иван Петров", Email: "ivan@example.com"},
        {Name: "Мария Сидорова", Email: "maria@example.com"},
        {Name: "Петр Иванов", Email: "petr@example.com"},
    }

    for _, user := range users {
        result, err := db.Exec(
            "INSERT INTO users (name, email) VALUES (?, ?)", 
            user.Name, user.Email,
        )
        if err != nil {
            log.Printf("❌ Ошибка вставки %s: %v\n", user.Name, err)
            continue
        }
        
        id, _ := result.LastInsertId()
        fmt.Printf("  ✅ Добавлен: %s (ID: %d)\n", user.Name, id)
    }

    // Запрашиваем данные
    fmt.Println("\n📋 Список всех пользователей:")
    fmt.Println("------------------------")
    rows, err := db.Query("SELECT id, name, email FROM users ORDER BY id")
    if err != nil {
        log.Fatal("Ошибка запроса:", err)
    }
    defer rows.Close()

    for rows.Next() {
        var user User
        err = rows.Scan(&user.ID, &user.Name, &user.Email)
        if err != nil {
            log.Fatal("Ошибка сканирования:", err)
        }
        fmt.Printf("  ID: %d | Имя: %s | Email: %s\n", 
            user.ID, user.Name, user.Email)
    }

    // Поиск по ID
    fmt.Println("\n🔍 Поиск пользователя с ID=2:")
    var user User
    err = db.QueryRow(
        "SELECT id, name, email FROM users WHERE id = ?", 2,
    ).Scan(&user.ID, &user.Name, &user.Email)
    
    if err != nil {
        if err == sql.ErrNoRows {
            fmt.Println("  ❌ Пользователь не найден")
        } else {
            log.Printf("Ошибка поиска: %v\n", err)
        }
    } else {
        fmt.Printf("  ✅ Найден: ID: %d, Имя: %s, Email: %s\n", 
            user.ID, user.Name, user.Email)
    }

    // Обновление данных
    fmt.Println("\n✏️ Обновление email пользователя с ID=1:")
    result, err := db.Exec(
        "UPDATE users SET email = ? WHERE id = ?", 
        "ivan.updated@example.com", 1,
    )
    if err != nil {
        log.Printf("❌ Ошибка обновления: %v\n", err)
    } else {
        rowsAffected, _ := result.RowsAffected()
        if rowsAffected > 0 {
            fmt.Println("  ✅ Email обновлен успешно")
        } else {
            fmt.Println("  ⚠️ Пользователь не найден")
        }
    }

    // Удаление пользователя
    fmt.Println("\n🗑️ Удаление пользователя с ID=3:")
    result, err = db.Exec("DELETE FROM users WHERE id = ?", 3)
    if err != nil {
        log.Printf("❌ Ошибка удаления: %v\n", err)
    } else {
        rowsAffected, _ := result.RowsAffected()
        if rowsAffected > 0 {
            fmt.Println("  ✅ Пользователь удален успешно")
        } else {
            fmt.Println("  ⚠️ Пользователь не найден")
        }
    }

    // Финальный запрос
    fmt.Println("\n📊 Финальный список пользователей:")
    fmt.Println("------------------------")
    rows, err = db.Query("SELECT id, name, email FROM users ORDER BY id")
    if err != nil {
        log.Fatal("Ошибка запроса:", err)
    }
    defer rows.Close()

    count := 0
    for rows.Next() {
        var user User
        err = rows.Scan(&user.ID, &user.Name, &user.Email)
        if err != nil {
            log.Fatal("Ошибка сканирования:", err)
        }
        fmt.Printf("  ID: %d | Имя: %s | Email: %s\n", 
            user.ID, user.Name, user.Email)
        count++
    }
    
    if count == 0 {
        fmt.Println("  📭 Таблица пуста")
    } else {
        fmt.Printf("\n  Всего записей: %d\n", count)
    }

    fmt.Println("\n✨ Программа успешно завершена!")
    fmt.Println("==========================")
}