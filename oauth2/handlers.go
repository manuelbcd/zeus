//
package oauth2

import (
    "log"
    "net/http"
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

    err = SetWithExpire(state.String(),true,120)
    if err != nil {
        log.Panic(err)
    }

    urlA := conf.AuthCodeURL(state.String())
    http.Redirect(response, request, urlA, http.StatusTemporaryRedirect)
}

//
func GitHubCallBack(response http.ResponseWriter, request *http.Request) {

    response.Header().Set("Content-Type","application/json")

    code := request.URL.Query().Get("code")
    state := request.URL.Query().Get("state")
    if code == "" || state == "" {
        http.Error(response, "Bad request baby", http.StatusBadRequest)
        return
    }

    unused, err := CheckState(state)
    if err != nil || !unused {
        log.Panic(err)
    }

    user, err := GetUser(request.Context(), code)
    if err != nil {
        log.Panic(err)
    }

    if err := json.NewEncoder(response).Encode(user); err != nil {
        log.Panic(err)
    }
}

func RedisHealth(response http.ResponseWriter, request *http.Request) {
    _,err := newSession()
    if err != nil {
        log.Panic(err)
    }
}
