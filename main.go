package main

import (

	"os"
	"log"
	"net/http"
	"gopkg.in/mgo.v2"
	"github.com/gorilla/mux"
	"golang.org/x/oauth2"
	 OAuth2GitHub "golang.org/x/oauth2/github"
	"github.com/google/go-github/github"
	"github.com/google/uuid"
)


var (

	conf = &oauth2.Config{

		ClientID: os.Getenv("GITHUB_CLIENT_ID"),
		ClientSecret: os.Getenv("GITHUB_CLIENT_SECRET"),
		Scopes: []string{

		},

		Endpoint: OAuth2GitHub.Endpoint,
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

	state, err := uuid.NewRandom()

	//TODO: Find best way to handle errors.
	if err != nil {
		log.Print("It has been a error generating 'state', a version 4 uuid.")
	}

    urlA := conf.AuthCodeURL(state.String())
	http.Redirect(response, request, urlA, http.StatusTemporaryRedirect)
}


func GetUser(response http.ResponseWriter, request *http.Request) {

	code := request.URL.Query().Get("code")
	state := request.URL.Query().Get("state")

	// redis key - value ?
	if code == "" || state == "" {
		http.Error(response, "Bad request baby", http.StatusBadRequest)
		return
	}

	ctx := request.Context()
	token, err := conf.Exchange(ctx, code)

	if err != nil {
		log.Fatal(err)
	}

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{ AccessToken: token.AccessToken },
	)

	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)
	user, _ , err  := client.Users.Get(ctx,"")

	if err != nil{
		log.Panic(err);
	}

	log.Print(*user.Name)
	log.Print(*user.Email)

	session := GetMongoDBSession()
	defer session.Close()
}
