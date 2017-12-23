package config

import (
    "sync"
    "log"
)

var (
    Config *configuration
    once sync.Once
)

type configuration struct {
    listeningAddress string
    redisAddress string
}

func (config *configuration) ListeningAddress() string {
    return config.listeningAddress
}

func (config *configuration) RedisAddress() string {
    return config.redisAddress
}

func init() {

    once.Do(func() {
        // TODO: Parse .json file.
        // TODO: Standardize logs
        log.Println("[INFO] initializing config package")
        Config  =  &configuration{
            listeningAddress: ":8080",
            redisAddress: "redis://redis",
        }
    })

}
