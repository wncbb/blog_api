package db

import "fmt"

type DBConfig struct {
	DriverName   string //"postgres, mysql"
	Timeout      string //ns, us, ms, s, m, h
	ReadTimeout  string
	WriteTimeout string
	User         string
	Password     string
	DBName       string
	Charset      string
	Host         string
	Port         string
	MaxIdleConns int
	MaxOpenConns int
	SSLMode      bool
	Debug        bool
}

func GetDefaultConfig(driverName string) *DBConfig {
	var ret *DBConfig
	switch driverName {
	case "postgres":
		ret = &DBConfig{
			DriverName:   "postgres",
			Timeout:      "1s",
			ReadTimeout:  "1s",
			WriteTimeout: "1s",
			User:         "",
			Password:     "",
			DBName:       "",
			Charset:      "",
			Host:         "127.0.0.1",
			Port:         "5432",
			MaxIdleConns: 10,
			MaxOpenConns: 10,
			SSLMode:      false,
		}
	case "mysql":
		ret = &DBConfig{
			DriverName:   "mysql",
			Timeout:      "1s",
			ReadTimeout:  "1s",
			WriteTimeout: "1s",
			User:         "",
			Password:     "",
			DBName:       "",
			Charset:      "",
			Host:         "127.0.0.1",
			Port:         "3306",
			MaxIdleConns: 10,
			MaxOpenConns: 10,
			SSLMode:      false,
		}
	}
	return ret
}

func (p *DBConfig) GetDefaultConfig() *DBConfig {
	var ret *DBConfig
	switch p.DriverName {
	case "postgres":
		ret = &DBConfig{
			DriverName:   "postgres",
			Timeout:      "1s",
			ReadTimeout:  "1s",
			WriteTimeout: "1s",
			User:         "",
			Password:     "",
			DBName:       "",
			Charset:      "",
			Host:         "127.0.0.1",
			Port:         "5432",
			MaxIdleConns: 10,
			MaxOpenConns: 10,
			SSLMode:      false,
		}
	case "mysql":
		ret = &DBConfig{
			DriverName:   "mysql",
			Timeout:      "1s",
			ReadTimeout:  "1s",
			WriteTimeout: "1s",
			User:         "",
			Password:     "",
			DBName:       "",
			Charset:      "",
			Host:         "127.0.0.1",
			Port:         "3306",
			MaxIdleConns: 10,
			MaxOpenConns: 10,
			SSLMode:      false,
		}
	}
	return ret
}

func (p *DBConfig) ToConnString() string {
	if p.Charset == "" {
		p.Charset = "utf8"
	}
	var configStr string
	switch p.DriverName {
	case "mysql":
		format := "%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local&timeout=%s&readTimeout=%s&writeTimeout=%s"
		configStr = fmt.Sprintf(
			format,
			p.User,
			p.Password,
			p.Host,
			p.Port,
			p.DBName,
			p.Charset,
			p.Timeout,
			p.ReadTimeout,
			p.WriteTimeout,
		)
	case "postgres":
		format := "user=%s password=%s host=%s port=%s dbname=%s"
		if !p.SSLMode {
			format += " sslmode=disable"
		}
		configStr = fmt.Sprintf(format, p.User, p.Password, p.Host, p.Port, p.DBName)
	}

	return configStr
}
