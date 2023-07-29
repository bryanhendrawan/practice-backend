package redis

import (
	"github.com/gomodule/redigo/redis"
)

var REDIS redis.Conn

func RedisInit() {
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		panic("Failed connect to Redis")
	}

	REDIS = conn
}
