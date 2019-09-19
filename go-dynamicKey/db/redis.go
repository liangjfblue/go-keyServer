package db

import (
    "github.com/garyburd/redigo/redis"
    "github.com/sirupsen/logrus"
    "go-dynamicKey/config"
    "time"
)

func NewRedisPool(cf *config.RedisConfig) *redis.Pool {
    return &redis.Pool {
        MaxIdle: cf.MaxIdle,
        MaxActive: cf.MaxActive,
        IdleTimeout: cf.IdleTimeout,
        Dial: func() (c redis.Conn, e error) {
            if c, e = redis.Dial(cf.Proto, cf.Addr); e != nil {
                return nil, e
            }

            if _, err := c.Do("AUTH", cf.Auth); err != nil {
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
