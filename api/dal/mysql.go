package dal

import (
	"blog_api/db"
)

func InitMysql(dbConfigMap map[string]*db.DBConfig) {
	db.InitConn(dbConfigMap)
}
