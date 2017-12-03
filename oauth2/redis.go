//
package oauth2

import (
    "github.com/garyburd/redigo/redis"
    "os"
)

var (
    rawurl = os.Getenv("REDIS_ADDRESS")
)

const (
    // Commands
    SET = "SET"
    GET = "GET"
    EXPIRE = "EXPIRE"
)

//
func newSession() (redis.Conn, error){

    c, err := redis.DialURL(rawurl)
    if err != nil {
        return nil, err
    }

    return c, nil
}

//
func GetBoolValue(key interface{}) (bool, error) {

    redisConn,err := newSession()
    if err != nil {
        return false, err
    }

    return redis.Bool(redisConn.Do(GET, key))
}

//
func SetWithExpire(key interface{}, value interface{}, expire int) error{

    redisConn,err := newSession()
    if err != nil {
        return err
    }
    defer redisConn.Close()

    _,err = redisConn.Do(SET,key,value)
    if err != nil {
        return err
    }

    _,err = redisConn.Do(EXPIRE,key,expire)
    return err
}

//
func Set(key interface{}, value interface{}) error{

    redisConn,err := newSession()
    if err != nil {
        return err
    }

    _, err = redisConn.Do(SET, key, value)
    return err
}
