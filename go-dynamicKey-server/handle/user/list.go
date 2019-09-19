package user

import (
    "encoding/json"
    "github.com/gin-gonic/gin"
    "github.com/sirupsen/logrus"
    "go-dynamicKey-server/pkg"
    "net/http"
    "strconv"
)

func List(c *gin.Context) {
    info := c.Query("info")

    out, err := pkg.DecodeSecurity(info)
    if err != nil {
        logrus.Error(err)
        c.JSON(http.StatusOK, gin.H{
            "Code":-1,
            "Msg":"decode info error",
        })
        return
    }

    var comment ListUserReq
    if err = json.Unmarshal(out, &comment); err != nil {
        logrus.Error(err)
        c.JSON(http.StatusOK, gin.H{
            "Code":-1,
            "Msg":"json info error",
        })
        return
    }

    users := make([]User, 0)
    for i := 1; i < 4; i++ {
        users = append(users, User{
            Name:"test"+strconv.Itoa(i),
            Age:20+i,
        })
    }

    mUser := make(map[string][]User)
    mUser["users"] = users

    jUsers, err := json.Marshal(mUser)
    if err != nil {
        logrus.Error(err)
        c.JSON(http.StatusOK, gin.H{
            "Code":-1,
            "Msg":"json info error",
        })
        return
    }

    var RespOut string
    if RespOut, err = pkg.EncryptSecurity(pkg.RandString(pkg.CodeSaltLen)+string(jUsers)); err != nil {
        logrus.Error("RespEncrypt error : ", err)
        c.JSON(http.StatusOK, gin.H{
            "Code":-1,
            "Msg":"encrypt error",
        })
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "Code":0,
        "Msg":"ok",
        "Data":RespOut,
    })
}
