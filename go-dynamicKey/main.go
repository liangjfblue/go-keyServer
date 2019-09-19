package main

import (
    "github.com/gin-gonic/gin"
    "github.com/sirupsen/logrus"
    "go-dynamicKey/config"
    "go-dynamicKey/router"
    "go-dynamicKey/service/security"
    "go-dynamicKey/service/server"
    "net/http"
    "os"
    "os/signal"
    "syscall"
)

func main() {
    serverConf := config.NewServerConf()

    g := gin.New()
    gin.SetMode(serverConf.HTTPConf.RunMode)

    router.Router(g)

    srv := server.NewServer(serverConf)

    security.SyncUpdateKey(srv)

    go func() {
        logrus.Infof("server start name %s, addr %s", serverConf.HTTPConf.Name, serverConf.HTTPConf.Addr)
        if err := http.ListenAndServe(serverConf.HTTPConf.Addr, g); err != nil {
            logrus.Error("server start error : ", err)
            panic(err)
        }
    }()

    c := make(chan os.Signal, 1)

    server.PingHealth(c, serverConf.HTTPConf)

    signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
    for {
        s := <-c
        logrus.Infof("server get a signal %s", s.String())
        switch s {
        case syscall.SIGTERM, syscall.SIGINT:
            logrus.Info("server exit")
            return
        case syscall.SIGQUIT:
            logrus.Info("ping error")
            return
        case syscall.SIGHUP:
        default:
            return
        }
    }
}
