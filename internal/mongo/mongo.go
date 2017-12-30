package mongo

import (
	"gopkg.in/mgo.v2"
	"os"
)

var (
	mongoURL = os.Getenv("MONGO_ADDRESS")
)

//
func getMongoDBSession() (*mgo.Session, error) {

	session, err := mgo.Dial(mongoURL)
	if err != nil {
		return nil, err
	}

	return session, nil
}

//
func Insert(database string, collection string, docs ...interface{}) error {

	session, err := getMongoDBSession()
	if err != nil {
		return err
	}
	defer session.Close()

	return session.DB(database).C(collection).Insert(docs)
}
