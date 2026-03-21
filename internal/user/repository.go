package user

import (
    "context"
    "database/sql"
)

type User struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}

type Repository interface {
    GetUser(ctx context.Context, id int) (*User, error)
    CreateUser(ctx context.Context, user *User) error
    UpdateUser(ctx context.Context, user *User) error
    DeleteUser(ctx context.Context, id int) error
    ListUsers(ctx context.Context) ([]User, error)
}

// SQLiteRepository реализация
type SQLiteRepository struct {
    db *sql.DB
}

func NewSQLiteRepository(db *sql.DB) *SQLiteRepository {
    return &SQLiteRepository{db: db}
}

func (r *SQLiteRepository) GetUser(ctx context.Context, id int) (*User, error) {
    var user User
    err := r.db.QueryRowContext(ctx,
        "SELECT id, name, email FROM users WHERE id = ?", id).
        Scan(&user.ID, &user.Name, &user.Email)
    if err != nil {
        return nil, err
    }
    return &user, nil
}

func (r *SQLiteRepository) CreateUser(ctx context.Context, user *User) error {
    result, err := r.db.ExecContext(ctx,
        "INSERT INTO users (name, email) VALUES (?, ?)",
        user.Name, user.Email)
    if err != nil {
        return err
    }
    
    id, err := result.LastInsertId()
    if err != nil {
        return err
    }
    user.ID = int(id)
    return nil
}

func (r *SQLiteRepository) UpdateUser(ctx context.Context, user *User) error {
    _, err := r.db.ExecContext(ctx,
        "UPDATE users SET name = ?, email = ? WHERE id = ?",
        user.Name, user.Email, user.ID)
    return err
}

func (r *SQLiteRepository) DeleteUser(ctx context.Context, id int) error {
    _, err := r.db.ExecContext(ctx, "DELETE FROM users WHERE id = ?", id)
    return err
}

func (r *SQLiteRepository) ListUsers(ctx context.Context) ([]User, error) {
    rows, err := r.db.QueryContext(ctx, "SELECT id, name, email FROM users ORDER BY id")
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    
    var users []User
    for rows.Next() {
        var user User
        if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
            return nil, err
        }
        users = append(users, user)
    }
    return users, rows.Err()
}