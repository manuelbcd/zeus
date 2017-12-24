package config

import (
    "sync"
    "log"
    "encoding/json"
    "io/ioutil"
    "errors"
)

var (
    Config *configuration
    once sync.Once
)

type configuration struct {
    listeningAddress string
    redisAddress string
}

func (config *configuration) UnmarshalJSON(b []byte) error {

    if config.redisAddress != "" {
        return errors.New("UnmarshalJSON of configuration struct is not allowed out its package")
    }

    aux := &struct{
        ListeningAddress string
        RedisAddress string
    }{}

    err := json.Unmarshal(b, &aux)
    if err != nil {
        return err
    }

    config.listeningAddress = aux.ListeningAddress
    config.redisAddress = aux.RedisAddress
    return nil
}

func (config *configuration) ListeningAddress() string {
    return config.listeningAddress
}

func (config *configuration) RedisAddress() string {
    return config.redisAddress
}

func init() {

    once.Do(func() {

        log.Println("[INFO] initializing config package")

        file, err := ioutil.ReadFile("./config.json")
        if err != nil {
            log.Println(err.Error())
        }

        json.Unmarshal(file, &Config)
    })

}
