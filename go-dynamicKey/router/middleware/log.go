package middleware

import (
    "github.com/gin-gonic/gin"
    "github.com/sirupsen/logrus"
    "time"
)

func LogMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        startTm := time.Now().UnixNano()

        c.Next()

        useTime := float64(time.Now().UnixNano() - startTm)/1e6
        logrus.WithFields(
            logrus.Fields{
                "method":c.Request.Method,
                "uri":c.Request.URL,
                "code":c.Writer.Status(),
                "useTime":useTime,
                "path":c.Request.URL.Path,
            })
    }
}
