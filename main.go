package main

import (
	//"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
	//"os"
)

func main() {

	session, err := mgo.Dial("127.0.0.1:27017")

	if err != nil {
		panic(err)
	}

	defer session.Close()

	router := mux.NewRouter()
	router.HandleFunc("/users", CreateUser).Methods("POST")
	log.Fatal(http.ListenAndServe("localhost:8080", router))

}

func CreateUser(response http.ResponseWriter, request *http.Request) {

}