//
package users

import (
    "github.com/marco2704/zeus/git"
    "github.com/marco2704/zeus/mongo"
    "github.com/google/uuid"
)

//
type User struct {
    *user
}

//
type user struct {
    Id string                   `json:"id,omitempty"`
    Email string                `json:"email,omitempty"`
    Name string                 `json:"name,omitempty"`
    LastName string             `json:"lastName,omitempty"`
    Repository *git.Repository  `json:"repository,omitempty"`
}

//
func createUser(email string, name string, lastName string) (*User, error) {

    uuidV1, err := uuid.NewUUID()
    if err != nil {
        return nil, err
    }

    id := uuid.NewSHA1(uuidV1, []byte(email)).String()
    user := newUser(id, email, name, lastName)
    err = mongo.Insert("zeus", "users", user)
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
        git.NewRepository(id),
        },
    }
}

//
func getUser(id string) *User {
    return nil
}

// Field appears in JSON as key "myName".
//Field int `json:"myName"`

// Field appears in JSON as key "myName" and
// the field is omitted from the object if its value is empty,
// as defined above.
//Field int `json:"myName,omitempty"`

// Field appears in JSON as key "Field" (the default), but
// the field is skipped if empty.
// Note the leading comma.
//Field int `json:",omitempty"`

// Field is ignored by this package.
//Field int `json:"-"`

// Field appears in JSON as key "-".
//Field int `json:"-,"`
