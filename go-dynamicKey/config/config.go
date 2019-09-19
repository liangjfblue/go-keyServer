package config

import (
    "github.com/fsnotify/fsnotify"
    "github.com/sirupsen/logrus"
    "github.com/spf13/viper"
    "os"
    "strings"
    "time"
)

// ServerConfig server config
type ServerConfig struct {
    HTTPConf	*HTTPConfig
    RedisConf	*RedisConfig
}

// HTTPConfig http config
type HTTPConfig struct {
    RunMode	string
    Addr	string
    Name	string
    PingMax int
}

// MysqlConfig mysql config
type RedisConfig struct {
    Proto           string
    Addr            string
    Auth            string
    MaxIdle         int
    MaxActive       int
    IdleTimeout     time.Duration
}

// NewServerConf new server config
func NewServerConf() *ServerConfig {
    return &ServerConfig{
        HTTPConf: &HTTPConfig{
            RunMode:	viper.GetString("runmode"),
            Addr:		viper.GetString("addr"),
            Name:		viper.GetString("name"),
            PingMax:	viper.GetInt("PingMax"),
        },
        RedisConf: &RedisConfig{
            Proto:          viper.GetString("redis.Proto"),
            Addr:           viper.GetString("redis.Addr"),
            Auth:           viper.GetString("redis.Auth"),
            MaxIdle:        viper.GetInt("redis.MaxIdle"),
            MaxActive:      viper.GetInt("redis.MaxActive"),
            IdleTimeout:    time.Duration(viper.GetInt("redis.IdleTimeout")),
        },
    }
}

// init init config
func init() {
    if err := initConfig(); err != nil {
        panic(err)
    }
    initLog()
    watchConfig()
}

// initConfig init config
func initConfig() error {
    viper.AddConfigPath(".")
    viper.SetConfigName("config")

    viper.SetConfigType("yaml")
    viper.AutomaticEnv()
    viper.SetEnvPrefix("GO_DYNAMICKEY")
    replacer := strings.NewReplacer(".", "_")
    viper.SetEnvKeyReplacer(replacer)
    if err := viper.ReadInConfig(); err != nil {
        return err
    }

    return nil
}

// watchConfig watch config file
func watchConfig() {
    viper.WatchConfig()
    viper.OnConfigChange(func(e fsnotify.Event) {
        logrus.Info("Config file changed: %s", e.Name)
    })
}

// initLog init log
func initLog() {
    logrus.SetFormatter(&logrus.JSONFormatter{})
    logrus.SetOutput(os.Stdout)
    logrus.SetLevel(logrus.Level(viper.GetInt("log.level")))
    logrus.SetReportCaller(viper.GetBool("log.reportCaller"))
}
