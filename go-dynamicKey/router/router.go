package router

import (
    "github.com/gin-gonic/gin"
    "go-dynamicKey/handle/health"
    "go-dynamicKey/handle/security"
    "go-dynamicKey/router/middleware"
    "net/http"
)

func Router(g *gin.Engine) {
    g.Use(gin.Recovery())

    g.NoRoute(func(c *gin.Context) {
        c.String(http.StatusNotFound, "The incorrect API route")
    })

    g.GET("/health", health.Health)

    g.Use(middleware.LogMiddleware())

    g.GET("/v1/security/key", security.Key)
}
