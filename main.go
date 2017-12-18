package utils

import (
    "log"
    "flag"
    "net/http"
    "github.com/gorilla/mux"
    "github.com/marco2704/zeus/oauth2"
    "github.com/marco2704/zeus/users"
)

const (
    addrName = "addr"
    addrUsage = "address where the go server is going to be listening"
    addrDefaultValue = ":80"
)

func main() {

    addr := flag.String(addrName, addrDefaultValue, addrUsage)

    router := mux.NewRouter()
    router.HandleFunc("/", oauth2.GitHubSignIn).Methods("GET")
    router.HandleFunc("/home", oauth2.GitHubCallBack).Methods("GET")
    router.HandleFunc("/oauth2/github", oauth2.GitHubOAuth).Methods("GET")
    router.HandleFunc("/users", users.CreateUser).Methods("POST")
    router.HandleFunc("/users/{id}", users.GetUser).Methods("GET")
    log.Fatal(http.ListenAndServe(*addr, router))
}
