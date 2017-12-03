//
package main

import (
    "os"
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "github.com/marco2704/zeus/oauth2"
)


var (
    redisURL = os.Getenv("REDIS_ADDRESS")
    serverListeningURL = os.Getenv("LISTENING_ADDRESS")
)

func main() {
    router := mux.NewRouter()
    router.HandleFunc("/", oauth2.GitHubSignIn).Methods("GET")
    router.HandleFunc("/home", oauth2.GitHubCallBack).Methods("GET")
    router.HandleFunc("/oauth2/github", oauth2.GitHubOAuth).Methods("GET")
    log.Fatal(http.ListenAndServe(serverListeningURL, router))
}

/**
package main

import (
    "github.com/marco2704/zeus-core/git"
)

var repositoryBasePath = "/opt/zeus/"

func main() {

    // TODO: Create or retrieve user
    // TODO: Deeper interfaces knowledge
    // TODO: Create version
    // TODO: Tests
}

type User struct {
    *user
}

type user struct {
    id string
    email string
    name string
    lastName string
    repository *git.Repository
}

func CreateUser() *User{

    //
    return nil
}

func GetUser(userId string) *User{
    return nil
}

func newUser(id string, email string, name string, lastName string) *User{

    return &User{&user{
        id,
        email,
        name,
        lastName,
        git.NewRepository(id),
        }}
}

 */