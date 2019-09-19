package health

import (
    "github.com/gin-gonic/gin"
    "go-dynamicKey/handle"
)

func Health(c *gin.Context) {
    heartbeat := c.Query("heartbeat")
    if heartbeat == "ok" {
        handle.SendResult(c, nil, "ok")
        return
    }
    handle.SendNotFound(c)
}
