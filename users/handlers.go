package users

import (
    "net/http"
    "github.com/gorilla/mux"
    "encoding/json"
    "log"
)

// GetUser retrieves a existing user.
func GetUser(response http.ResponseWriter, request *http.Request) {

    // TODO: Token Checker
    vars := mux.Vars(request)
    user := getUser(vars["id"])
    if user == nil {
        http.Error(response, "The requested user does not exist", http.StatusNotFound)
        return
    }

    if err := json.NewEncoder(response).Encode(user); err != nil {
        log.Print(err)
    }
}

// CreateUser creates a new user and return it if everything went well.
func CreateUser(response http.ResponseWriter, request *http.Request) {

    user, err := createUser(
            request.FormValue("email"),
            request.FormValue("name"),
            request.FormValue("lastName"),
    )
    if err != nil {
        log.Print(err)
    }

    if err := json.NewEncoder(response).Encode(user); err != nil {
        log.Print(err)
    }
}
