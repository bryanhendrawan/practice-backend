package model

import (
	"practice-backend/redis"

	redigo "github.com/gomodule/redigo/redis"
	"github.com/vmihailenco/msgpack"
)

func Set(key string, value interface{}) error {
	bytes, err := msgpack.Marshal(value)
	if err != nil {
		return err
	}
	_, err = redis.REDIS.Do("SET", key, bytes)
	return err
}

func Get(key string, variable interface{}) (bool, error) {
	bytes, err := redigo.Bytes(redis.REDIS.Do("GET", key))
	if err == redigo.ErrNil {
		return false, nil
	}
	if err != nil {
		return false, err
	}

	err = msgpack.Unmarshal(bytes, variable)
	return err == nil, err
}
