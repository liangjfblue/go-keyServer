package security

import (
    "encoding/json"
    "github.com/gin-gonic/gin"
    "github.com/sirupsen/logrus"
    "go-dynamicKey/handle"
    "go-dynamicKey/pkg/rand"
    "go-dynamicKey/service/security"
)

func Key(c *gin.Context) {
    var (
        err error
        reqOut  []byte
        RespOut string
    )

    req := c.Query("info")
    if reqOut, err = security.DecodeSecurity(req); err != nil {
        logrus.Error("DecodeSecurity error : ", err)
        handle.SendBadRequest(c)
        return
    }

    mReqOut := make(map[string]interface{})
    if err = json.Unmarshal(reqOut, &mReqOut); err != nil {
        logrus.Error("json.Unmarshal error : ", err)
        handle.SendBadRequest(c)
        return
    }

    if mReqOut["mid"].(string) == "" || mReqOut["sn"].(string) == "" {
        logrus.Error("mid/sn is empty")
        handle.SendBadRequest(c)
        return
    }

    secretKey := security.GetSecretKey()
    if len(secretKey) == 0 {
        logrus.Error("")
        handle.SendInternalServerError(c)
        return
    }

    if RespOut, err = security.EncryptSecurity(rand.RandString(security.CodeSaltLen)+secretKey); err != nil {
        logrus.Error("EncryptSecurity error : ", err)
        handle.SendInternalServerError(c)
        return
    }

    handle.SendContent(c, []byte(RespOut))
}
