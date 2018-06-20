package dal

import (
	"wncbb.cn/db"
)

func InitMysql(dbConfigMap map[string]*db.DBConfig) {
	db.InitConn(dbConfigMap)
}
