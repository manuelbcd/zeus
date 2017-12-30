//
package oauth2

import (
	"github.com/garyburd/redigo/redis"
	"github.com/marco2704/zeus/internal/config"
)

const (
	SET    = "SET"
	GET    = "GET"
	EXPIRE = "EXPIRE"
)

// newSession creates a new connection to redis server
func newSession() (redis.Conn, error) {

	c, err := redis.DialURL(config.Config.RedisAddress())
	if err != nil {
		return nil, err
	}

	return c, nil
}

// CheckState gets the value by using state param as a key. If the value is true,
// false is set as new value for the next time this state is checked.
func CheckState(state string) (bool, error) {

	value, err := GetBoolValue(state)
	if err != nil {
		return false, err
	}

	if value {
		return value, Set(state, false)
	}

	return value, nil
}

// GetBoolValue gets a bool value by the given key param.
func GetBoolValue(key string) (bool, error) {

	redisConn, err := newSession()
	if err != nil {
		return false, err
	}

	value, err := redis.Bool(redisConn.Do(GET, key))
	if err != nil {
		return false, err
	}

	return value, nil
}

// SetWithExpire sets a key-value entry with timeout.
func SetWithExpire(key interface{}, value interface{}, seconds int) error {

	redisConn, err := newSession()
	if err != nil {
		return err
	}
	defer redisConn.Close()

	_, err = redisConn.Do(SET, key, value)
	if err != nil {
		return err
	}

	_, err = redisConn.Do(EXPIRE, key, seconds)
	return err
}

//
func Set(key interface{}, value interface{}) error {

	redisConn, err := newSession()
	if err != nil {
		return err
	}

	_, err = redisConn.Do(SET, key, value)
	return err
}
