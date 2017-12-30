package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/marco2704/zeus/internal/user"
	"log"
	"net/http"
)

//
func GetUser(response http.ResponseWriter, request *http.Request) {

	// TODO: Token Checker
	vars := mux.Vars(request)
	user := user.GetUser(vars["id"])
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

	createdUser, err := user.CreateUser(
		request.FormValue("email"),
		request.FormValue("name"),
		request.FormValue("lastName"),
	)

	if err != nil {
		log.Print(err)
	}

	if err := json.NewEncoder(response).Encode(createdUser); err != nil {
		log.Print(err)
	}
}
