package pkg

import (
    "encoding/json"
    "errors"
    "github.com/garyburd/redigo/redis"
    "github.com/robfig/cron"
    "github.com/sirupsen/logrus"
    "go-dynamicKey-server/db"
)

type SecurityKey struct {
    ID  string `json:"id"`
    Key string `json:"key"`
    Iv  string  `json:"iv"`
}

var (
    KEY = "pub_key"
    Key SecurityKey
)

func updateKey() {
    cli := db.RedisPool.Get()
    defer cli.Close()

    key, _ := redis.String(cli.Do("GET", KEY))
    if key == "" {
        logrus.Error("key empty")
        panic(errors.New("key empty"))
    }

    logrus.Info(key)
    setSecretKey(key)
}

func SyncUpdateKey() {
    updateKey()

    c := cron.New()
    if err := c.AddFunc("*/30 * * * * ?", func() { updateKey() }); err != nil {
        logrus.Error("add timer update key error : ", err)
        panic(err)
    }
    c.Start()
}

func setSecretKey(key string) {
    if err := json.Unmarshal([]byte(key), &Key); err != nil {
        logrus.Error(err)
    }
}
