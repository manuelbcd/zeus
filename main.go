package main

import (
	"github.com/gorilla/mux"
	"github.com/marco2704/zeus/handlers"
	"github.com/marco2704/zeus/internal/config"
	"log"
	"net/http"
)

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/home", handlers.GitHubCallBack).Methods("GET")
	router.HandleFunc("/oauth2/github", handlers.GitHubOAuth).Methods("GET")
	router.HandleFunc("/user", handlers.CreateUser).Methods("POST")
	router.HandleFunc("/user/{id}", handlers.GetUser).Methods("GET")
	log.Fatal(http.ListenAndServe(config.Config.ListeningAddress(), router))
}
