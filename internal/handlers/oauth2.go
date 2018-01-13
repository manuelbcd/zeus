package handlers

import (
	"encoding/json"
	"github.com/marco2704/zeus/internal/oauth2"
	"github.com/marco2704/zeus/pkg/types"
	"log"
	"net/http"
)

const (
	providerParam = "provider"
	codeParam     = "code"
	stateParam    = "state"
)

// GitHubOAuth redirects to authorization provider URL with a unique state.
func GitHubOAuth(response http.ResponseWriter, request *http.Request) {

	provider := request.URL.Query().Get(providerParam)
	client := oauth2.GetClient(provider)
	if client != nil {
		http.Error(response, "Bad request baby", http.StatusBadRequest)
	}

	url := client.AuthCodeURL()
	http.Redirect(response, request, url, http.StatusTemporaryRedirect)
}

// GitHubCallBack handles a callback from the provider taking a code and state
// query parameters.
func GitHubCallBack(response http.ResponseWriter, request *http.Request) {

	response.Header().Set("Content-Type", "application/json")

	code := request.URL.Query().Get(codeParam)
	state := request.URL.Query().Get(stateParam)
	if types.IsAnyEmpty(code, state) {
		http.Error(response, "Bad request baby", http.StatusBadRequest)
		return
	}

	client := oauth2.GetClientByState(state)
	user, err := client.GetUser(request.Context(), code)
	if err != nil {
		log.Println(err)
	}

	if err := json.NewEncoder(response).Encode(user); err != nil {
		log.Println(err)
	}
}
