package user

import (
    "encoding/json"
    "net/http"
    "strconv"

    "github.com/gorilla/mux"
)

// Handler содержит репозиторий для работы с БД
type Handler struct {
    repo Repository
}

// NewHandler создает новый обработчик
func NewHandler(repo Repository) *Handler {
    return &Handler{repo: repo}
}

// CreateUserRequest структура запроса на создание пользователя
type CreateUserRequest struct {
    Name  string `json:"name"`
    Email string `json:"email"`
}

// CreateUserResponse структура ответа после создания
type CreateUserResponse struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}

// CreateUser — POST /api/users
func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
    var req CreateUserRequest
    
    // Декодируем JSON запрос
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, `{"error": "Invalid request body"}`, http.StatusBadRequest)
        return
    }
    
    // Валидация
    if req.Name == "" {
        http.Error(w, `{"error": "Name is required"}`, http.StatusBadRequest)
        return
    }
    if req.Email == "" {
        http.Error(w, `{"error": "Email is required"}`, http.StatusBadRequest)
        return
    }
    
    // Создаем пользователя
    user := &User{
        Name:  req.Name,
        Email: req.Email,
    }
    
    if err := h.repo.CreateUser(r.Context(), user); err != nil {
        http.Error(w, `{"error": "`+err.Error()+`"}`, http.StatusInternalServerError)
        return
    }
    
    // Отправляем ответ
    resp := CreateUserResponse{
        ID:    user.ID,
        Name:  user.Name,
        Email: user.Email,
    }
    
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(resp)
}

// GetUser — GET /api/users/{id}
func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) {
    // Получаем ID из URL
    vars := mux.Vars(r)
    idStr := vars["id"]
    
    id, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, `{"error": "Invalid user ID"}`, http.StatusBadRequest)
        return
    }
    
    // Получаем пользователя из БД
    user, err := h.repo.GetUser(r.Context(), id)
    if err != nil {
        http.Error(w, `{"error": "User not found"}`, http.StatusNotFound)
        return
    }
    
    // Отправляем ответ
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(user)
}

// ListUsers — GET /api/users
func (h *Handler) ListUsers(w http.ResponseWriter, r *http.Request) {
    // Получаем список пользователей
    users, err := h.repo.ListUsers(r.Context())
    if err != nil {
        http.Error(w, `{"error": "`+err.Error()+`"}`, http.StatusInternalServerError)
        return
    }
    
    // Отправляем ответ
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(users)
}

// UpdateUser — PUT /api/users/{id}
func (h *Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    idStr := vars["id"]
    
    id, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, `{"error": "Invalid user ID"}`, http.StatusBadRequest)
        return
    }
    
    var req CreateUserRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, `{"error": "Invalid request body"}`, http.StatusBadRequest)
        return
    }
    
    // Обновляем пользователя
    user := &User{
        ID:    id,
        Name:  req.Name,
        Email: req.Email,
    }
    
    if err := h.repo.UpdateUser(r.Context(), user); err != nil {
        http.Error(w, `{"error": "`+err.Error()+`"}`, http.StatusInternalServerError)
        return
    }
    
    // Отправляем ответ
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(user)
}

// DeleteUser — DELETE /api/users/{id}
func (h *Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    idStr := vars["id"]
    
    id, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, `{"error": "Invalid user ID"}`, http.StatusBadRequest)
        return
    }
    
    // Удаляем пользователя
    if err := h.repo.DeleteUser(r.Context(), id); err != nil {
        http.Error(w, `{"error": "`+err.Error()+`"}`, http.StatusInternalServerError)
        return
    }
    
    // Отправляем успешный ответ
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusNoContent)
}