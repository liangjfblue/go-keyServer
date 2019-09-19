package handle

import (
    "github.com/gin-gonic/gin"
    "github.com/sirupsen/logrus"
    "go-dynamicKey/pkg/errno"
    "net/http"
)

type Result struct {
    Code    int         `json:"code"`
    Message string      `json:"message"`
    Data    interface{} `json:"data"`
}

func SendBadRequest(c *gin.Context) {
    c.Writer.WriteHeader(http.StatusBadRequest)
}

func SendNotFound(c *gin.Context) {
    c.Writer.WriteHeader(http.StatusNotFound)
}

func SendInternalServerError(c *gin.Context) {
    c.Writer.WriteHeader(http.StatusInternalServerError)
}

func SendContent(c *gin.Context, data []byte) {
    _, err := c.Writer.Write(data)
    if err != nil {
        logrus.Error("write error : ", err)
    }
}

func SendResult(c *gin.Context, err error, data interface{}) {
    code, message := errno.DecodeErr(err)

    // always return http.StatusOK
    c.JSON(http.StatusOK, Result{
        Code:    code,
        Message: message,
        Data:    data,
    })
}

