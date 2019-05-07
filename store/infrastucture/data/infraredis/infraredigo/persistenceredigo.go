package infraredigo

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

var pool = newPool()

func newPool() *redis.Pool {
	return &redis.Pool{

		// Maximum number of idle connections in the pool.
		MaxIdle: 5,

		// max number of connections
		MaxActive: 200,

		// Dial is an application supplied function for creating and
		// configuring a connection.
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", "192.168.228.150:6379", redis.DialDatabase(11))
			if err != nil {
				panic(err.Error())
			}
			return c, err
		},
	}
}

// PersistenceRedigo struct to provides methods redis
type persistenceRedigo struct {
}

func NewPersistenceRedigo() *persistenceRedigo {
	return &persistenceRedigo{}
}

// Ping test connection server
func (pr persistenceRedigo) Ping() (string, error) {
	conn := pool.Get()
	defer conn.Close()

	resp, err := redis.String(conn.Do("PING"))
	if err != nil {
		return "", err
	}

	fmt.Printf("PING Response = %s\n", resp)

	return resp, nil
}

// GetKey GET value
func (pr persistenceRedigo) GetKey(key string) (string, error) {
	conn := pool.Get()
	defer conn.Close()

	s, err := redis.String(conn.Do("GET", key))
	if err != nil && err != redis.ErrNil {
		return "", (err)
	}

	return s, nil
}
