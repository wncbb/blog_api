package user

import (
	"fmt"
	"testing"

	"blog_api/db"
	"blog_api/log"
)

func init() {
	var err error
	configMap := make(map[string]*db.DBConfig, 0)
	configMap["write"] = &db.DBConfig{
		DriverName:   "postgres",
		Timeout:      "1s",
		ReadTimeout:  "1s",
		WriteTimeout: "1s",
		User:         "postgres",
		Password:     "asd123",
		DBName:       "test1",
		Charset:      "",
		Host:         "127.0.0.1",
		Port:         "5432",
		MaxIdleConns: 5,
		MaxOpenConns: 5,
		SSLMode:      false,
		Debug:        true,
	}
	configMap["read"] = &db.DBConfig{
		DriverName:   "postgres",
		Timeout:      "1s",
		ReadTimeout:  "1s",
		WriteTimeout: "1s",
		User:         "postgres",
		Password:     "asd123",
		DBName:       "test1",
		Charset:      "",
		Host:         "127.0.0.1",
		Port:         "5432",
		MaxIdleConns: 5,
		MaxOpenConns: 5,
		SSLMode:      false,
		Debug:        false,
	}
	err = db.InitConn(configMap)
	if err != nil {
		//panic(errors.Cause(err))
		fmt.Printf("%+v", err)
		panic(err)
	}

}

func init() {
	Init()
}

func Test_Migrate(t *testing.T) {
	Migrate()
}
func Test_Create(t *testing.T) {
	userCore := &UserCore{
		Account:  "todd2",
		Password: "1234",
	}
	err := Create(userCore)
	fmt.Printf("err:%v\n", err)
}

func Test_create(t *testing.T) {
	userCore := &UserCore{
		Account:  "todd2",
		Password: "1234",
	}
	conn, _ := writeDbHandler.GetConnection()
	err := create(conn, userCore)
	fmt.Printf("err:%v\n", err)
}

func Test_ListAll(t *testing.T) {
	userCoreList, err := ListAll()
	if err != nil {
		log.DefaultLog().Errorf("Test_ListAll error:%v", err)
		panic("error")
	}

	for k, v := range userCoreList {
		log.DefaultLog().Infof("id:%d,  %v", k, v)
	}

}
