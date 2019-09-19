package main

import (
    "go-dynamicKey-client/api"
    "go-dynamicKey-client/config"
)

func main() {
    config.Init()

    if key, err := api.GetSecurityKeyV1(); err == nil {
        //logrus.Infof("id:%s key:%s iv:%s", key.ID, key.Key, key.Iv)
        api.UserList(key.Key, key.Iv)
    }
}

