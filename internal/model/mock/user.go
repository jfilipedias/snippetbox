package mock

import (
	"time"

	"github.com/jfilipedias/snippetbox/internal/model"
)

var mockUser = model.User{
	ID:             1,
	Name:           "Alice",
	Email:          "alice@example.com",
	HashedPassword: []byte{},
	Created:        time.Now(),
}

type UserModel struct{}

func (m *UserModel) Insert(name, email, password string) error {
	switch email {
	case "dupe@example.com":
		return model.ErrDuplicatedEmail
	default:
		return nil
	}
}

func (m *UserModel) Authenticate(email, password string) (int, error) {
	if email == "alice@example.com" && password == "pa$$word" {
		return 1, nil
	}

	return 0, model.ErrInvalidCredentials
}

func (m *UserModel) Exists(id int) (bool, error) {
	switch id {
	case 1:
		return true, nil
	default:
		return false, nil
	}
}

func (m *UserModel) Get(id int) (model.User, error) {
	switch id {
	case 1:
		return mockUser, nil
	default:
		return model.User{}, model.ErrNoRecord
	}
}
