package config

import (
	"flag"
	"fmt"
	"os"

	"blog_api/log"

	"github.com/spf13/viper"
)

type Conf struct {
	ConfName     string
	Ip           string
	Port         string
	PgDb         string
	SessionRedis string
	TempRedis    string
	Cors         string
}

var conf *Conf

func init() {
	conf = &Conf{
		ConfName:     "./conf/conf.yaml",
		Ip:           "",
		Port:         "",
		PgDb:         "",
		SessionRedis: "",
		TempRedis:    "",
		Cors:         "",
	}
}

func parseCmdArgs() {
	confName := flag.String("confame", "./conf/conf.yaml", "conf location")
	ip := flag.String("ip", "", "ip")
	port := flag.String("port", "", "port")
	flag.Parse()
	if *confName != "" {
		conf.ConfName = *confName
	}
	if *ip != "" {
		conf.Ip = *ip
	}
	if *port != "" {
		conf.Port = *port
	}
}

func parseConfFile() {
	v := viper.New()
	v.SetConfigFile(conf.ConfName)
	if err := v.ReadInConfig(); err != nil {
		msg := fmt.Sprintf("Failed to read conf: %s", conf.ConfName)
		fmt.Fprintf(os.Stderr, "%s\n", msg)
		panic(err)
	}
	if err := v.Unmarshal(conf); err != nil {
		msg := fmt.Sprintf("Failed to parse conf: %s", conf.ConfName)
		fmt.Fprintf(os.Stderr, "%s\n", msg)
		panic(err)
	}
	log.DefaultLog().Debug("conf: %v", conf)
}

func Init() {
	parseConfFile()
	parseCmdArgs()
	parsePgDbYamlConf(conf.PgDb)
}

func RunAddr() string {

	return fmt.Sprintf("%s:%s", conf.Ip, conf.Port)
}

func GetSessionRedisConfFile() string {
	return conf.SessionRedis
}

func GetTempRedisConfFile() string {
	return conf.TempRedis
}

func GetCorsConfFile() string {
	return conf.Cors
}
