package main

import (

    "os"
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "golang.org/x/oauth2"
     OAuth2GitHub "golang.org/x/oauth2/github"
    "github.com/google/go-github/github"
    "github.com/google/uuid"
    "github.com/garyburd/redigo/redis"
    "encoding/json"
)


var (

	conf = &oauth2.Config{

            ClientID: os.Getenv("GITHUB_CLIENT_ID"),
            ClientSecret: os.Getenv("GITHUB_CLIENT_SECRET"),
            Scopes: []string{

            },

		Endpoint: OAuth2GitHub.Endpoint,
	}

	redisURL = os.Getenv("REDIS_ADDRESS")
	serverListeningURL = os.Getenv("LISTENING_ADDRESS")
)

func main() {

    router := mux.NewRouter()
    router.HandleFunc("/github-oauth", GitHubOAuth).Methods("GET")
    router.HandleFunc("/users", GetUser).Methods("GET")

    log.Fatal(http.ListenAndServe(serverListeningURL, router))
}

func GetRedisDBSession() (*redis.Conn){

    c, err := redis.DialURL(redisURL)

    if err != nil {
        panic(err)
    }

    return &c
}

func GitHubOAuth(response http.ResponseWriter, request *http.Request) {

	state, err := uuid.NewRandom()
    if err != nil {
        panic(err)
    }

    redisConn := *GetRedisDBSession()
    defer redisConn.Close()

    redisConn.Do("SET", state, true)

    urlA := conf.AuthCodeURL(state.String())
    http.Redirect(response, request, urlA, http.StatusTemporaryRedirect)
}

func GetUser(response http.ResponseWriter, request *http.Request) {

    code := request.URL.Query().Get("code")
    state := request.URL.Query().Get("state")

    if code == "" || state == "" {
        http.Error(response, "Bad request baby", http.StatusBadRequest)
        return
    }

    redisConn := *GetRedisDBSession()
    defer redisConn.Close()

    unused, err:= redis.Bool(redisConn.Do("GET", state));

    if err != nil || !unused{
        panic(err)
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

    if err := json.NewEncoder(response).Encode(user); err != nil {
        panic(err)
    }
}
