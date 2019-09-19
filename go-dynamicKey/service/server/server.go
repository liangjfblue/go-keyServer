package server

import (
    "fmt"
    "github.com/garyburd/redigo/redis"
    "github.com/sirupsen/logrus"
    "go-dynamicKey/config"
    "go-dynamicKey/db"
    "net/http"
    "os"
    "syscall"
)

type Server struct {
    ServerConf  *config.ServerConfig
    RedisPool   *redis.Pool
}

func NewServer(serverConf *config.ServerConfig) *Server {
    return &Server{
        ServerConf: serverConf,
        RedisPool:  db.NewRedisPool(serverConf.RedisConf),
    }
}

func PingHealth(c chan os.Signal, config *config.HTTPConfig) {
    go func() {
        pingURL := fmt.Sprintf("http://127.0.0.1%s/health?heartbeat=ok", config.Addr)
        for i := 0; i < config.PingMax; i++ {
            resp, err := http.Get(pingURL);
            if err == nil && resp != nil && resp.StatusCode == http.StatusOK {
                logrus.Info("ping ok")
                return
            }
        }
        //can not ping server
        c <- syscall.SIGQUIT
    }()
}
