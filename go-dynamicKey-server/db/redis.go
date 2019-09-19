package db

import (
    "github.com/garyburd/redigo/redis"
    "github.com/sirupsen/logrus"
    "github.com/spf13/viper"
    "time"
)

var RedisPool *redis.Pool

func init() {
    RedisPool = NewRedisPool()
}

func NewRedisPool() *redis.Pool {
    return &redis.Pool {
        MaxIdle:        viper.GetInt("redis.MaxIdle"),
        MaxActive:      viper.GetInt("redis.MaxActive"),
        IdleTimeout:    time.Duration(viper.GetInt("redis.IdleTimeout")),
        Dial: func() (c redis.Conn, e error) {
            if c, e = redis.Dial(viper.GetString("redis.Proto"), viper.GetString("redis.Addr")); e != nil {
                return nil, e
            }

            if _, err := c.Do("AUTH", viper.GetString("redis.Auth")); err != nil {
                c.Close()
                logrus.Error(err)
                return nil, err
            }
            return c, e
        },
        TestOnBorrow: func(c redis.Conn, t time.Time) error {
            _, err := c.Do("PING")
            if err != nil {
                logrus.Error(err)
            }
            return err
        },
    }
}
