package users

import (
	"github.com/google/uuid"
	"github.com/marco2704/zeus/internal/mongo"
	"github.com/marco2704/zeus/internal/repository"
)

// User
type User struct {
	*user
}

// user
type user struct {
	Id         string                 `json:"id,omitempty"`
	Email      string                 `json:"email,omitempty"`
	Name       string                 `json:"name,omitempty"`
	Repository *repository.Repository `json:"repository,omitempty"`
}

// CreateUser creates a new user with a random id. Also it initializes its repository, if
// everything was well a User pointer is returned with a nil error, otherwise, a no nil error
// is returned with a nil User pointer.
func CreateUser(email string, name string) (*User, error) {

	uuidV1, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	id := uuid.NewSHA1(uuidV1, []byte(email)).String()
	user := newUser(id, email, name)
	err = mongo.Insert("zeus", "users", user)
	if err != nil {
		return nil, err
	}

	if err := user.Repository.Init(); err != nil {
		return nil, err
	}

	return user, nil
}

// newUser creates a new User struct with the given params.
func newUser(id string, email string, name string) *User {

	return &User{&user{
		id,
		email,
		name,
		repository.NewRepository(id),
	},
	}
}

// TODO: implement it
func GetUser(id string) *User {
	return nil
}
