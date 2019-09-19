package api

import (
    "encoding/json"
    "errors"
    "fmt"
    "github.com/sirupsen/logrus"
    "github.com/spf13/viper"
    "go-dynamicKey-client/pkg"
    "io/ioutil"
    "net/http"
)

type ListUserReq struct {
    Mid string  `json:"mid"`
    Sn  string  `json:"sn"`
}

type User struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}

func UserList(key, iv string) {
    req := ListUserReq{
        Mid:"mid",
        Sn:"sn",
    }

    jReq, err := json.Marshal(req)
    if err != nil {
        logrus.Error(err)
        return
    }

    sReq, err := pkg.EncryptSecurity(pkg.RandString(pkg.CodeSaltLen)+string(jReq))
    if err != nil {
        logrus.Error(err)
        return
    }

    appServer := viper.GetString("appServer.Addr")
    URL := fmt.Sprintf("http://%s/v1/user/list?info=%s", appServer, sReq)
    resp, err := http.Get(URL)
    if err != nil {
        logrus.Error(err)
        return
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        logrus.Error(err)
        return
    }

    result := struct {
        Code    int     `json:"Code"`
        Msg     string  `json:"Msg"`
        Data    string  `json:"Data"`
    }{}
    if err = json.Unmarshal(body, &result); err != nil {
        logrus.Error(err)
        return
    }
    if result.Data == "" {
        logrus.Error(errors.New("data empty"))
        return
    }

    out, err := pkg.DecodeByDynamics(result.Data, key, iv)
    if err != nil {
        logrus.Error(err)
        return
    }

    mUser := make(map[string][]User)
    if err = json.Unmarshal(out, &mUser); err != nil {
        logrus.Error(err)
        return
    }

    for _, v := range mUser["users"] {
         logrus.Info(v)
    }
}
