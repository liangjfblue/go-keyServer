package main

import (
    "github.com/gin-gonic/gin"
    "github.com/sirupsen/logrus"
    "github.com/spf13/viper"
    "go-dynamicKey-server/config"
    "go-dynamicKey-server/pkg"
    "go-dynamicKey-server/router"
    "net/http"
)

func main() {
    g := gin.Default()

    config.Init()
    
    gin.SetMode(viper.GetString("runmode"))

    //sync update security key
    pkg.SyncUpdateKey()

    router.Router(g)
    logrus.Error(http.ListenAndServe(viper.GetString("addr"), g))
}

