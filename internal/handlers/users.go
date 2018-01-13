package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/marco2704/zeus/internal/users"
	"log"
	"net/http"
)

// TODO: implement it
func GetUser(response http.ResponseWriter, request *http.Request) {

	// TODO: Token Checker
	vars := mux.Vars(request)
	user := users.GetUser(vars["id"])
	if user == nil {
		http.Error(response, "The requested users does not exist", http.StatusNotFound)
		return
	}

	if err := json.NewEncoder(response).Encode(user); err != nil {
		log.Print(err)
	}
}

// CreateUser creates a new users and return it if everything went well.
func CreateUser(response http.ResponseWriter, request *http.Request) {

	response.Header().Set("Content-Type", "application/json")

	createdUser, err := users.CreateUser(
		request.FormValue("email"),
		request.FormValue("name"),
	)

	if err != nil {
		log.Print(err)
	}

	if err := json.NewEncoder(response).Encode(createdUser); err != nil {
		log.Print(err)
	}
}
