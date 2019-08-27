package gredis

import (
    "encoding/json"
    "gin-test/utils/setting"
    "github.com/gomodule/redigo/redis"
    "log"
    "time"
)

var RedisConn *redis.Pool

// Setup Initialize the Redis instance
func Setup() {
    RedisConn = &redis.Pool{
        MaxIdle:     setting.RedisSetting.MaxIdle,
        MaxActive:   setting.RedisSetting.MaxActive,
        IdleTimeout: setting.RedisSetting.IdleTimeout,
        Dial: func() (redis.Conn, error) {
            c, err := redis.Dial("tcp", setting.RedisSetting.Host)
            if err != nil {
                return nil, err
            }

            if setting.RedisSetting.Password != "" {
                if _, err := c.Do("AUTH", setting.RedisSetting.Password); err != nil {
                    c.Close()
                    return nil, err
                }
            }

            if _, err := c.Do("SELECT", setting.RedisSetting.DB); err != nil {
                c.Close()
                return nil, err
            }

            return c, err
        },
        TestOnBorrow: func(c redis.Conn, t time.Time) error {
            _, err := c.Do("PING")
            return err
        },
    }

    TestConnection()

    log.Printf("[info] Redis connected %s DB: %d", setting.RedisSetting.Host, setting.RedisSetting.DB)
}

func TestConnection() {
    conn := RedisConn.Get()
    defer conn.Close()

    _, err := conn.Do("PING")
    if err != nil {
        panic(err)
    }
}

// Set a key/value
func Set(key string, data interface{}, time int) error {
    conn := RedisConn.Get()
    defer conn.Close()

    value, err := json.Marshal(data)
    if err != nil {
        return err
    }

    _, err = conn.Do("SET", key, value)
    if err != nil {
        return err
    }

    _, err = conn.Do("EXPIRE", key, time)
    if err != nil {
        return err
    }

    return nil
}

// Exists check a key
func Exists(key string) bool {
    conn := RedisConn.Get()
    defer conn.Close()

    exists, err := redis.Bool(conn.Do("EXISTS", key))
    if err != nil {
        return false
    }

    return exists
}

// Get get a key
func Get(key string) ([]byte, error) {
    conn := RedisConn.Get()
    defer conn.Close()

    reply, err := redis.Bytes(conn.Do("GET", key))
    if err != nil {
        return nil, err
    }

    return reply, nil
}

// Delete delete a kye
func Delete(key string) (bool, error) {
    conn := RedisConn.Get()
    defer conn.Close()

    return redis.Bool(conn.Do("DEL", key))
}