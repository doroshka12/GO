package user

import (
    "context"
    "errors"
    "testing"

    "github.com/doroshka12/GO/sqlite-example/mocks" // путь к вашим мокам
    "github.com/stretchr/testify/assert"
    "go.uber.org/mock/gomock"
)

func TestSQLiteRepository_GetUser_Success(t *testing.T) {
    // Создаем контроллер gomock
    ctrl := gomock.NewController(t)
    defer ctrl.Finish() // проверяет, что все ожидаемые вызовы были сделаны
    
    // Создаем мок репозитория
    mockRepo := mocks.NewMockUserRepository(ctrl)
    
    // Подготавливаем тестовые данные
    expectedUser := &User{
        ID:    1,
        Name:  "Тестовый Пользователь",
        Email: "test@example.com",
    }
    
    // Настраиваем ожидание: метод GetUser должен быть вызван с id=1
    // и вернуть expectedUser и nil (ошибку)
    mockRepo.EXPECT().
        GetUser(gomock.Any(), 1). // gomock.Any() - любой context
        Return(expectedUser, nil).
        Times(1) // ожидаем ровно один вызов
    
    // Вызываем тестируемый метод
    user, err := mockRepo.GetUser(context.Background(), 1)
    
    // Проверяем результаты
    assert.NoError(t, err)
    assert.Equal(t, expectedUser, user)
}

func TestSQLiteRepository_GetUser_NotFound(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()
    
    mockRepo := mocks.NewMockUserRepository(ctrl)
    
    // Настраиваем ожидание: метод вернет ошибку "sql: no rows in result set"
    expectedErr := errors.New("sql: no rows in result set")
    mockRepo.EXPECT().
        GetUser(gomock.Any(), 999).
        Return(nil, expectedErr).
        Times(1)
    
    user, err := mockRepo.GetUser(context.Background(), 999)
    
    assert.Error(t, err)
    assert.Nil(t, user)
    assert.Equal(t, expectedErr.Error(), err.Error())
}

func TestSQLiteRepository_CreateUser_Success(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()
    
    mockRepo := mocks.NewMockUserRepository(ctrl)
    
    // Новый пользователь (без ID)
    newUser := &User{
        Name:  "Новый Пользователь",
        Email: "new@example.com",
    }
    
    // Ожидаем, что после создания у пользователя появится ID
    // В моке мы должны имитировать это поведение
    mockRepo.EXPECT().
        CreateUser(gomock.Any(), gomock.Eq(newUser)). // проверяем, что передается нужный user
        DoAndReturn(func(_ context.Context, u *User) error {
            // Имитируем присвоение ID
            u.ID = 42
            return nil
        }).
        Times(1)
    
    // Вызываем метод
    err := mockRepo.CreateUser(context.Background(), newUser)
    
    // Проверяем результаты
    assert.NoError(t, err)
    assert.Equal(t, 42, newUser.ID) // ID должен быть установлен
}

func TestSQLiteRepository_CreateUser_DuplicateEmail(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()
    
    mockRepo := mocks.NewMockUserRepository(ctrl)
    
    duplicateUser := &User{
        Name:  "Дубликат",
        Email: "existing@example.com",
    }
    
    expectedErr := errors.New("UNIQUE constraint failed: users.email")
    mockRepo.EXPECT().
        CreateUser(gomock.Any(), duplicateUser).
        Return(expectedErr).
        Times(1)
    
    err := mockRepo.CreateUser(context.Background(), duplicateUser)
    
    assert.Error(t, err)
    assert.Equal(t, expectedErr.Error(), err.Error())
}

func TestSQLiteRepository_ListUsers_Success(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()
    
    mockRepo := mocks.NewMockUserRepository(ctrl)
    
    expectedUsers := []User{
        {ID: 1, Name: "User 1", Email: "user1@example.com"},
        {ID: 2, Name: "User 2", Email: "user2@example.com"},
        {ID: 3, Name: "User 3", Email: "user3@example.com"},
    }
    
    mockRepo.EXPECT().
        ListUsers(gomock.Any()).
        Return(expectedUsers, nil).
        Times(1)
    
    users, err := mockRepo.ListUsers(context.Background())
    
    assert.NoError(t, err)
    assert.Equal(t, expectedUsers, users)
    assert.Len(t, users, 3)
}

func TestSQLiteRepository_UpdateUser_Success(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()
    
    mockRepo := mocks.NewMockUserRepository(ctrl)
    
    updatedUser := &User{
        ID:    1,
        Name:  "Обновленное Имя",
        Email: "updated@example.com",
    }
    
    mockRepo.EXPECT().
        UpdateUser(gomock.Any(), updatedUser).
        Return(nil).
        Times(1)
    
    err := mockRepo.UpdateUser(context.Background(), updatedUser)
    
    assert.NoError(t, err)
}

func TestSQLiteRepository_DeleteUser_Success(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()
    
    mockRepo := mocks.NewMockUserRepository(ctrl)
    
    mockRepo.EXPECT().
        DeleteUser(gomock.Any(), 42).
        Return(nil).
        Times(1)
    
    err := mockRepo.DeleteUser(context.Background(), 42)
    
    assert.NoError(t, err)
}