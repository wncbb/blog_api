package session

import (
	"fmt"
	"testing"

	"blog_api/db/pg"
)

var repo *SessionRepository

func init() {
	var err error
	configMap := make(map[string]*pg.DBConfig, 0)
	configMap["test"] = &pg.DBConfig{
		DriverName:   "postgres",
		Timeout:      "1s",
		ReadTimeout:  "1s",
		WriteTimeout: "1s",
		User:         "test2",
		Password:     "test1",
		DBName:       "test1",
		Charset:      "",
		Host:         "127.0.0.1",
		Port:         "5432",
		MaxIdleConns: 5,
		MaxOpenConns: 5,
		SSLMode:      false,
	}
	err = pg.InitConn(configMap)
	if err != nil {
		//panic(errors.Cause(err))
		fmt.Printf("%+v", err)
		panic(err)
	}

	handler, err := pg.GetHandler("test")
	if err != nil {
		panic(err)
	}
	repo = NewSessionRepository(handler)
}

func Test_Migrate(t *testing.T) {
	err := repo.Migrate()
	if err != nil {
		t.Errorf("err: %v", err)
	}
}
