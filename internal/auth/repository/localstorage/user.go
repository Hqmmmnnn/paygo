package localstorage

import (
	"context"
	"sync"

	"github.com/Hqqm/paygo/internal/auth"
	"github.com/Hqqm/paygo/internal/auth/entities"
	uuid "github.com/satori/go.uuid"
)

type UsersLocalStorage struct {
	users map[string]*entities.User
	mutex    *sync.Mutex
}

func NewUsersLocalStorage() *UsersLocalStorage {
	return &UsersLocalStorage{
		users: make(map[string] *entities.User),
		mutex:    new(sync.Mutex),
	}
}

func (usersStorage *UsersLocalStorage) CreateUserID(ctx context.Context, userID uuid.UUID) error {
	user := &entities.User{
		ID: userID,
		FirstName: "",
		LastName: "",
		Patronymic: "",
	}
	usersStorage.mutex.Lock()
	usersStorage.users[userID.String()] = user
	usersStorage.mutex.Unlock()
	return nil
}

func (usersStorage *UsersLocalStorage) GetUser(ctx context.Context, userID uuid.UUID) (*entities.User, error) {
	usersStorage.mutex.Lock()
	defer  usersStorage.mutex.Unlock()

	for _, user := range usersStorage.users {
		if user.ID == userID {
			return user, nil
		}
	}

	return nil, auth.ErrUserNotFound
}
