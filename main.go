package main

import (
	"github.com/gorilla/mux"
	"github.com/marco2704/zeus/internal/config"
	"github.com/marco2704/zeus/internal/handlers"
	"log"
	"net/http"
)

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/home", handlers.GitHubCallBack).Methods("GET")
	router.HandleFunc("/oauth2/github", handlers.GitHubOAuth).Methods("GET")
	router.HandleFunc("/users", handlers.CreateUser).Methods("POST")
	// TODO: un-comment when handlers.GetUser is completely implemented.
	//router.HandleFunc("/users/{id}", handlers.GetUser).Methods("GET")
	log.Fatal(http.ListenAndServe(config.Config.ListeningAddress(), router))
}
