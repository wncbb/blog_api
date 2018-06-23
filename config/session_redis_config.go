package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type SessionRedisConfig struct {
	Ip        string
	Port      string
	MaxIdlNum int
	Secret    string
	Password  string
	Db        string
}

func (p *SessionRedisConfig) GetConnString() string {
	return fmt.Sprintf("%s:%s", p.Ip, p.Port)
}

func GetDefaultSessionRedisConfig() *SessionRedisConfig {
	return &SessionRedisConfig{
		Ip:        "127.0.0.1",
		Port:      "6379",
		MaxIdlNum: 10,
		Secret:    "secret",
		Password:  "",
		Db:        "0",
	}
}

var sessionRedisConfig *SessionRedisConfig

func GetSessionRedisConfig() *SessionRedisConfig {
	configFile := GetSessionRedisConfFile()
	v := viper.New()
	v.SetConfigFile(configFile)
	if err := v.ReadInConfig(); err != nil {
		msg := fmt.Sprintf("Failed to load session redis config: %s, %s", configFile, err)
		fmt.Fprintf(os.Stderr, "%s\n", msg)
		panic(msg)
	}

	subV := v.Sub("session")
	if subV == nil {
		msg := fmt.Sprintf("Failed to parse session redis ")
		fmt.Fprintf(os.Stderr, "%s\n", msg)
		panic(msg)
	}
	sessionRedisConfig = GetDefaultSessionRedisConfig()
	if err := subV.Unmarshal(sessionRedisConfig); err != nil {
		msg := fmt.Sprintf("Failed to unmarshal session session config: %s", err)
		fmt.Fprintf(os.Stderr, "%s\n", msg)
		panic(msg)
	}

	return sessionRedisConfig
}
