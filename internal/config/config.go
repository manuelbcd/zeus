package config

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"sync"
)

// configuration unexported struct that holds most configuration
// parameters which are taken from config.json file.
type configuration struct {
	listeningAddress string
	repoRootPath     string
	mongoAddress     string
	redisAddress     string
}

// UnmarshalJSON allows unmarshal a .json file to configuration
// struct using a anonymous struct with exported files.
func (config *configuration) UnmarshalJSON(b []byte) error {

	if config.redisAddress != "" {
		return errors.New("error - UnmarshalJSON of configuration struct is not allowed out its package")
	}

	aux := &struct {
		ListeningAddress string
		RepoRootPath     string
		MongoAddress     string
		RedisAddress     string
	}{}

	err := json.Unmarshal(b, &aux)
	if err != nil {
		return err
	}

	config.listeningAddress = aux.ListeningAddress
	config.repoRootPath = aux.RepoRootPath
	config.mongoAddress = aux.MongoAddress
	config.redisAddress = aux.RedisAddress
	return nil
}

// ListeningAddress returns the unexported listeningAddress field from
// configuration struct.
func (config *configuration) ListeningAddress() string {
	return config.listeningAddress
}

// RedisAddress returns the unexported redisAddress field from
// configuration struct.
func (config *configuration) RedisAddress() string {
	return config.redisAddress
}

// MongoAddress returns the unexported mongoAddress field from
// configuration struct.
func (config *configuration) MongoAddress() string {
	return config.mongoAddress
}

// RepoRootPath returns the unexported repoRootPath field from
// configuration struct.
func (config *configuration) RepoRootPath() string {
	return config.repoRootPath
}

var (
	// Config holds all base configuration.
	Config *configuration
	once   sync.Once
)

// init method initialize a configuration struct from a .json file just once.
func init() {

	once.Do(func() {

		log.Println("info - initializing config package")

		file, err := ioutil.ReadFile("./config.json")
		if err != nil {
			log.Println(err.Error())
		}

		json.Unmarshal(file, &Config)
		log.Print(string(file))
	})
}
