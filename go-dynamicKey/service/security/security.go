package security

import (
    "errors"
    "github.com/garyburd/redigo/redis"
    "github.com/robfig/cron"
    "github.com/sirupsen/logrus"
    "go-dynamicKey/service/server"
)

var (
    KEY = "pub_key"
    GCipher string
)

func SyncUpdateKey(srv *server.Server) {
    updateKey(srv)

    timerUpdateKey(srv)
}

func updateKey(srv *server.Server) {
    cli := srv.RedisPool.Get()
    defer cli.Close()

    key, err := redis.String(cli.Do("GET", KEY))
    if err != nil || key == "" {
        logrus.Error("key empty")
        panic(errors.New("key empty"))
    }

    logrus.Info(key)
    setSecretKey(key)
}

func timerUpdateKey(srv *server.Server) {
    c := cron.New()
    //every 60 s update once
    if err := c.AddFunc("*/60 * * * * ?", func() { updateKey(srv) }); err != nil {
       logrus.Error("add timer update key error : ", err)
       panic(err)
    }
    c.Start()
}

func setSecretKey(key string) {
    GCipher = key
}

func GetSecretKey() string {
    return GCipher
}

