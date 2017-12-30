package handlers

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/marco2704/zeus/internal/oauth2"
	"log"
	"net/http"
)

//
func GitHubOAuth(response http.ResponseWriter, request *http.Request) {

	state, err := uuid.NewRandom()
	if err != nil {
		panic(err)
	}

	err = oauth2.SetWithExpire(state.String(), true, 120)
	if err != nil {
		log.Panic(err)
	}

	url := oauth2.ProviderAuthURL(state.String())
	http.Redirect(response, request, url, http.StatusTemporaryRedirect)
}

//
func GitHubCallBack(response http.ResponseWriter, request *http.Request) {

	response.Header().Set("Content-Type", "application/json")

	code := request.URL.Query().Get("code")
	state := request.URL.Query().Get("state")
	if code == "" || state == "" {
		http.Error(response, "Bad request baby", http.StatusBadRequest)
		return
	}

	unused, err := oauth2.CheckState(state)
	if err != nil || !unused {
		log.Panic(err)
	}

	user, err := oauth2.GetUser(request.Context(), code)
	if err != nil {
		log.Panic(err)
	}

	if err := json.NewEncoder(response).Encode(user); err != nil {
		log.Panic(err)
	}
}
