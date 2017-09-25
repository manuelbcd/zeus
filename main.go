package main

import (

	"log"
	"net/http"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
	"os"
)

var mongoAddress = os.Getenv("MONGO_ADDRESS")
var listeningAddress = os.Getenv("LISTENING_ADDRESS")

func GetMongoDBSession() *mgo.Session{

	session, err := mgo.Dial(mongoAddress)

	if err != nil {
		panic(err)
	}

	return session
}

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/users", CreateUser).Methods("POST")
	log.Fatal(http.ListenAndServe(listeningAddress, router))
}

func CreateUser(response http.ResponseWriter, request *http.Request) {

	session := GetMongoDBSession()
	defer session.Close()
	response.Write([]byte("Everything is under controll"))
}