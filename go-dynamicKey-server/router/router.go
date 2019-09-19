package router

import (
    "github.com/gin-gonic/gin"
    "go-dynamicKey-server/handle/user"
    "net/http"
)

func Router(g *gin.Engine) {
    g.Use(gin.Recovery())

    g.NoRoute(func(c *gin.Context) {
        c.String(http.StatusNotFound, "The incorrect API route")
    })

    g.GET("/v1/user/list", user.List)
}

