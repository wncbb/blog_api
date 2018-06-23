package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type TempRedisConfig struct {
	Ip        string
	Port      string
	MaxIdlNum int
	Secret    string
	Password  string
	Db        string
}

func (p *TempRedisConfig) GetConnString() string {
	return fmt.Sprintf("%s:%s", p.Ip, p.Port)
}

func GetDefaultTempRedisConfig() *TempRedisConfig {
	return &TempRedisConfig{
		Ip:        "127.0.0.1",
		Port:      "6379",
		MaxIdlNum: 10,
		Secret:    "secret",
		Password:  "",
		Db:        "0",
	}
}

var tempRedisConfig *TempRedisConfig

func GetTempRedisConfig() *TempRedisConfig {
	configFile := GetTempRedisConfFile()
	v := viper.New()
	v.SetConfigFile(configFile)
	if err := v.ReadInConfig(); err != nil {
		msg := fmt.Sprintf("Failed to load temp redis config: %s, %s", configFile, err)
		fmt.Fprintf(os.Stderr, "%s\n", msg)
		panic(msg)
	}

	subV := v.Sub("temp")
	if subV == nil {
		msg := fmt.Sprintf("Failed to parse temp redis ")
		fmt.Fprintf(os.Stderr, "%s\n", msg)
		panic(msg)
	}
	tempRedisConfig = GetDefaultTempRedisConfig()
	if err := subV.Unmarshal(tempRedisConfig); err != nil {
		msg := fmt.Sprintf("Failed to unmarshal temp redis config: %s", err)
		fmt.Fprintf(os.Stderr, "%s\n", msg)
		panic(msg)
	}

	return tempRedisConfig
}
