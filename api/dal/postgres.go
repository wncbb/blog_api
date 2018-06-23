package dal

import (
	"blog_api/db"
)

func InitPostgres(dbConfigMap map[string]*db.DBConfig) {
	db.InitConn(dbConfigMap)
}
