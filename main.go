package main

import (

	"log"
	"net/http"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
	"os"
)

var GITHUB_AUTH_URL = os.Getenv("GITHUB_AUTH_URL")
var MONGO_ADDRESS = os.Getenv("MONGO_ADDRESS")
var LISTENING_ADDRESS = os.Getenv("LISTENING_ADDRESS")

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/github-oauth", GitHubOAuth).Methods("GET")
	router.HandleFunc("/github-oauth/access-token", GetAccessToken).Methods("GET")

	log.Fatal(http.ListenAndServe(LISTENING_ADDRESS, router))
}


func GetMongoDBSession() *mgo.Session{

	session, err := mgo.Dial(MONGO_ADDRESS)

	if err != nil {
		panic(err)
	}

	return session
}


func GitHubOAuth(response http.ResponseWriter, request *http.Request) {
    http.Redirect(response, request, GITHUB_AUTH_URL, 301)
}


func GetAccessToken(response http.ResponseWriter, request *http.Request) {

	temp_code := request.URL.Query().Get("code")

	if temp_code != "" {

		session := GetMongoDBSession()
		defer session.Close()

	}

	//https://github.com/login/oauth/access_token
	//SE LLAMA A ESTA PAGINA VARIAS VECES?
	// POST REQUEST 
}



