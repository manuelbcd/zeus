package mongo

import (
	"github.com/marco2704/zeus/internal/config"
	"gopkg.in/mgo.v2"
)

// getMongoDBSession returns a pointer of a new mgo.Session.
func getMongoDBSession() (*mgo.Session, error) {

	session, err := mgo.Dial(config.Config.MongoAddress())
	if err != nil {
		return nil, err
	}

	return session, nil
}

// Insert inserts any quantity of any type into any collection of any existing database.
func Insert(database string, collection string, docs ...interface{}) error {

	session, err := getMongoDBSession()
	if err != nil {
		return err
	}
	defer session.Close()

	return session.DB(database).C(collection).Insert(docs...)
}
