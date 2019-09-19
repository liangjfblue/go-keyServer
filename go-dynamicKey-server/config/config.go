package config

import (
    "github.com/fsnotify/fsnotify"
    "github.com/sirupsen/logrus"
    "github.com/spf13/viper"
    "os"
    "strings"
)

// init 初始化config
func Init() {
    if err := initConfig(); err != nil {
        panic(err)
    }
    initLog()
    watchConfig()
}

// initConfig 初始化config
func initConfig() error {
    viper.AddConfigPath(".")
    viper.SetConfigName("config")

    viper.SetConfigType("yaml")
    viper.AutomaticEnv()
    viper.SetEnvPrefix("GO_DYNAMICKEY_SERVER")
    replacer := strings.NewReplacer(".", "_")
    viper.SetEnvKeyReplacer(replacer)
    if err := viper.ReadInConfig(); err != nil {
        return err
    }

    return nil
}

// watchConfig 监听config文件
func watchConfig() {
    viper.WatchConfig()
    viper.OnConfigChange(func(e fsnotify.Event) {
        logrus.Info("Config file changed: %s", e.Name)
    })
}

// initLog 初始化log
func initLog() {
    logrus.SetFormatter(&logrus.JSONFormatter{})
    logrus.SetOutput(os.Stdout)
    logrus.SetLevel(logrus.Level(viper.GetInt("log.level")))
    logrus.SetReportCaller(viper.GetBool("log.reportCaller"))
}
