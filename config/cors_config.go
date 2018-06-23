package config

import (
	middleware_cors "blog_api/api/middleware/cors"
	"fmt"
	"os"

	"github.com/spf13/viper"
)

func GetCorsConfig() *middleware_cors.Config {
	var err error
	configFile := GetCorsConfFile()
	v := viper.New()
	v.SetConfigFile(configFile)
	if err = v.ReadInConfig(); err != nil {
		msg := fmt.Sprintf("Failed to load cors config file:%s, err:%v", configFile, err)
		fmt.Fprintf(os.Stderr, "%s\n", msg)
		panic(msg)
	}

	corsConfig := &middleware_cors.Config{}
	if err = v.Unmarshal(corsConfig); err != nil {
		msg := fmt.Sprintf("Faild to parse cors config file:%s, err:%v", configFile, err)
		panic(msg)
	}
	return corsConfig
}
