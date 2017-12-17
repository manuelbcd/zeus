package utils

import (
    "os"
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "github.com/marco2704/zeus/oauth2"
    "github.com/marco2704/zeus/users"
)

var (
    serverListeningURL = os.Getenv("LISTENING_ADDRESS")
)

func main() {
    router := mux.NewRouter()
    router.HandleFunc("/", oauth2.GitHubSignIn).Methods("GET")
    router.HandleFunc("/home", oauth2.GitHubCallBack).Methods("GET")
    router.HandleFunc("/oauth2/github", oauth2.GitHubOAuth).Methods("GET")
    router.HandleFunc("/users", users.CreateUser).Methods("POST")
    router.HandleFunc("/users/{id}", users.GetUser).Methods("GET")
    log.Fatal(http.ListenAndServe(serverListeningURL, router))
}
