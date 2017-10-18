package main

import (

	"os"
	"log"
	"context"
	"net/http"
	"gopkg.in/mgo.v2"
	"github.com/gorilla/mux"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)


var (

	conf = &oauth2.Config{

		ClientID: os.Getenv("GITHUB_CLIENT_ID"),
		ClientSecret: os.Getenv("GITHUB_CLIENT_SECRET"),
		Endpoint: github.Endpoint,
		// TODO: what does scopes mean in this context?
	}

	mongoURL = os.Getenv("MONGO_ADDRESS")
	serverListeningURL = os.Getenv("LISTENING_ADDRESS")

)

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/github-oauth", GitHubOAuth).Methods("GET")
	router.HandleFunc("/users", GetUser).Methods("GET")

	log.Fatal(http.ListenAndServe(serverListeningURL, router))
}


func GetMongoDBSession() *mgo.Session{

	session, err := mgo.Dial(mongoURL)

	if err != nil {
		panic(err)
	}

	return session
}


func GitHubOAuth(response http.ResponseWriter, request *http.Request) {

	//TODO: replace 'state' by a real
    urlA := conf.AuthCodeURL("state")
	http.Redirect(response, request, urlA, http.StatusTemporaryRedirect)

}


func GetUser(response http.ResponseWriter, request *http.Request) {

	code := request.URL.Query().Get("code")
	state := request.URL.Query().Get("state")

	if code == "" || state == "" {
		http.Error(response, "Bad request baby", http.StatusBadRequest)
		return
	}

	// TODO: what the context does
	token, err := conf.Exchange(context.TODO(), code)

	if err != nil {
		log.Fatal(err)
	}

	// README we got the access_token !!
	log.Print(token.AccessToken)

	session := GetMongoDBSession()
	defer session.Close()
}



