package user

import (
	"github.com/google/uuid"
	"github.com/marco2704/zeus/internal/mongo"
	"github.com/marco2704/zeus/internal/repository"
)

//
type User struct {
	*user
}

//
type user struct {
	Id         string                 `json:"id,omitempty"`
	Email      string                 `json:"email,omitempty"`
	Name       string                 `json:"name,omitempty"`
	LastName   string                 `json:"lastName,omitempty"`
	Repository *repository.Repository `json:"repository,omitempty"`
}

//
func CreateUser(email string, name string, lastName string) (*User, error) {

	uuidV1, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	id := uuid.NewSHA1(uuidV1, []byte(email)).String()
	user := newUser(id, email, name, lastName)
	err = mongo.Insert("zeus", "user", user)
	if err != nil {
		return nil, err
	}

	if err := user.Repository.Init(); err != nil {
		return nil, err
	}

	return user, nil
}

//
func newUser(id string, email string, name string, lastName string) *User {

	return &User{&user{
		id,
		email,
		name,
		lastName,
		repository.NewRepository(id),
	},
	}
}

//
func GetUser(id string) *User {
	return nil
}
