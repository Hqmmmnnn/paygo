package localstorage

import (
	"context"
	"testing"

	"github.com/Hqqm/paygo/internal/auth"
	"github.com/Hqqm/paygo/internal/auth/entities"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

func TestGetUser(t *testing.T) {
	usersStorage := NewUsersLocalStorage()

	userId := uuid.NewV4()
	err := usersStorage.CreateUserID(context.Background(), userId)
	assert.NoError(t, err)

	newUser := &entities.User{
		ID:         userId,
		FirstName: "",
		LastName: "",
		Patronymic: "",
	}
	user, err := usersStorage.GetUser(context.Background(), userId)
	assert.Equal(t, newUser, user)
	assert.NoError(t, err)

	fakeId := uuid.NewV4()
	user, err = usersStorage.GetUser(context.Background(), fakeId)
	assert.Error(t, err)
	assert.Equal(t, err, auth.ErrUserNotFound)
}
