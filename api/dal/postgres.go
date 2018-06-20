package dal

import (
	"wncbb.cn/db"
)

func InitPostgres(dbConfigMap map[string]*db.DBConfig) {
	db.InitConn(dbConfigMap)
}
