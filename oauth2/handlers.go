//
package oauth2

import (
    "net/http"
    "log"
    "encoding/json"
    "github.com/google/uuid"
    "html/template"
)

//
func GitHubSignIn(response http.ResponseWriter, _ *http.Request) {
    t, _ := template.ParseFiles( "client.html")
    t.Execute(response, nil)
}

//
func GitHubOAuth(response http.ResponseWriter, request *http.Request) {

    state, err := uuid.NewRandom()
    if err != nil {
        panic(err)
    }

    err = SetWithExpire(state,true,120)
    if err != nil {
        log.Panic(err)
    }

    urlA := conf.AuthCodeURL(state.String())
    http.Redirect(response, request, urlA, http.StatusTemporaryRedirect)
}

//
func GitHubCallBack(response http.ResponseWriter, request *http.Request) {

    code := request.URL.Query().Get("code")
    state := request.URL.Query().Get("state")

    if code == "" || state == "" {
        http.Error(response, "Bad request baby", http.StatusBadRequest)
        return
    }

    unused, err := GetBoolValue(state)
    if err != nil || !unused {
        log.Panic(err)
    }
    Set(state,false)

    user, err := GetUser(request.Context(), code)
    if err != nil {
        log.Panic(err)
    }

    response.Header().Set("Content-Type","application/json")
    if err := json.NewEncoder(response).Encode(user); err != nil {
        log.Panic(err)
    }
}