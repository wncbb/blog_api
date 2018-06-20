package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
	"wncbb.cn/db"
)

var PgReadModule = "read"
var PgWriteModule = "write"
var pgDbModuleList []string = []string{PgReadModule, PgWriteModule}

var pgDBConfMap map[string]*db.DBConfig

func parsePgDbYamlConf(dbConfName string) {
	pgDBConfMap = make(map[string]*db.DBConfig)
	v := viper.New()
	v.SetConfigFile(dbConfName)
	if err := v.ReadInConfig(); err != nil {
		msg := fmt.Sprintf("Failed to load pg db config: %s, %s", dbConfName, err)
		fmt.Fprintf(os.Stderr, "%s\n", msg)
		panic(msg)
	}

	for _, subModeName := range pgDbModuleList {
		subV := v.Sub(subModeName)
		if subV == nil {
			msg := fmt.Sprintf("Failed to parse config sub module: %s", subModeName)
			fmt.Fprintf(os.Stderr, "%s\n", msg)
			panic(msg)
		}
		subDBOptional := db.GetDefaultConfig("postgres")
		pgDBConfMap[subModeName] = subDBOptional
		if err := subV.Unmarshal(pgDBConfMap[subModeName]); err != nil {
			msg := fmt.Sprintf("Failed to unmarshal pg db config: %s", err)
			fmt.Fprintf(os.Stderr, "%s\n", msg)
			panic(msg)
		}
	}
	for k, v := range pgDBConfMap {
		fmt.Printf("LINE36 %v %v %+v\n", k, v, v)
	}
}

func GetPgDBConf() map[string]*db.DBConfig {
	return pgDBConfMap
}
