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

type Key struct {
    ID  string `json:"id"`
    Key string `json:"key"`
    Iv  string `json:"iv"`
}

func GetSecurityKeyV1() (*Key, error) {
    info := struct {
        Mid string `json:"mid"`
        Sn  string `json:"sn"`
    }{
        "mid",
        "sn",
    }

    jInfo, err := json.Marshal(info)
    if err != nil {
        logrus.Error(err)
        return nil, err
    }

    sInfo, err := pkg.EncryptSecurity(pkg.RandString(pkg.CodeSaltLen)+string(jInfo))
    if err != nil {
        logrus.Error(err)
        return nil, err
    }

    scurityServer := viper.GetString("scurityServer.Addr")
    URL := fmt.Sprintf("http://%s/v1/security/key?info=%s", scurityServer, sInfo)
    resp, err := http.Get(URL)
    if err != nil {
        logrus.Error(err)
        return nil, err
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        logrus.Error(err)
        return nil, err
    }

    out, err := pkg.DecodeSecurity(string(body))
    if err != nil {
        logrus.Error(err)
        return nil, err
    }

    var key Key
    if err = json.Unmarshal(out, &key); err != nil {
        logrus.Error(err)
        return nil, err
    }

    if key.ID == "" || key.Key == "" || key.Iv == "" {
        err = errors.New(fmt.Sprintf("id/key/iv is empty : %s %s %s", key.ID, key.Key, key.Iv))
        logrus.Error(err)
        return nil, err
    }

    return &key, nil
}
