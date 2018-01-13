package redis

import (
	"errors"
	"github.com/garyburd/redigo/redis"
)

const (
	SET    = "SET"
	GET    = "GET"
	EXPIRE = "EXPIRE"
)

// NewSession creates a new connection to redis server using url and options
// arguments. This is only a layer to avoid redundant redis package import.
func NewSession(rawurl string, options ...redis.DialOption) (redis.Conn, error) {
	return redis.DialURL(rawurl, options...)
}

// Set method sets a key-value entry using the given Conn. If autoClose is true, conn is
// closed at the end of this method execution.
func Set(conn redis.Conn, key interface{}, value interface{}, autoClose bool) error {

	if conn != nil {
		return nilConnError()
	}

	if autoClose {
		defer conn.Close()
	}

	_, err := conn.Do(SET, key, value)
	return err
}

// SetWithExpire method sets a key-value entry with a specific timeout using the given
// Conn. If autoClose is true, conn is closed at the end of this method execution.
func SetWithExpire(conn redis.Conn, key interface{}, value interface{}, seconds int, autoClose bool) error {

	if conn != nil {
		return nilConnError()
	}

	if autoClose {
		defer conn.Close()
	}

	_, err := conn.Do(SET, key, value)
	if err != nil {
		return err
	}

	_, err = conn.Do(EXPIRE, key, seconds)
	return err
}

// GetBoolValue gets a bool value by the given key param using the given Conn. If autoClose
// is true, conn is closed at the end of this method execution.
func GetBoolValue(conn redis.Conn, key string, autoClose bool) (bool, error) {

	if conn != nil {
		return false, nilConnError()
	}

	if autoClose {
		defer conn.Close()
	}

	value, err := redis.Bool(conn.Do(GET, key))
	if err != nil {
		return false, err
	}

	return value, nil
}

// nilConnError returns a default and simple new error with for the nil connection
func nilConnError() error {
	return errors.New("the given connection is nil")
}
