package db

import (
	"fmt"
	"sync"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/pkg/errors"
	"wncbb.cn/werrors"
)

type DBHandler struct {
	DB        *gorm.DB
	Config    *DBConfig
	Connected bool
	mu        sync.Mutex
}

func NewDBHandler() *DBHandler {
	return &DBHandler{}
}

func NewDBHandlerWithConfig(config *DBConfig) *DBHandler {
	return &DBHandler{
		Config: config,
	}
}

func (p *DBHandler) Close() {
	p.DB.DB().Close()
	p.Connected = false
}

func (p *DBHandler) Connect() error {
	var err error
	if p.Connected {
		return nil
	}

	p.mu.Lock()
	defer p.mu.Unlock()

	if p.Connected {
		return nil
	}

	if nil == p.Config {
		return errors.WithStack(werrors.ERR_DB_CONFIG_NIL)
	}

	p.DB, err = gorm.Open(p.Config.DriverName, p.Config.ToConnString())
	if err != nil {
		return errors.Wrapf(
			err,
			"gorm.Open('%s', '%s') fail",
			p.Config.DriverName,
			p.Config.ToConnString())
	}

	if p.Config.Debug {
		p.DB = p.DB.Debug()
	}

	p.DB.DB().SetMaxIdleConns(p.Config.MaxIdleConns)
	p.DB.DB().SetMaxOpenConns(p.Config.MaxOpenConns)

	return nil
}

func FreeConn() {
	for k, v := range DBMap {
		if v != nil {
			v.Close()
			delete(DBMap, k)
		}
	}
}

func InitConn(dbConfigMap map[string]*DBConfig) error {
	var err error
	for k, v := range dbConfigMap {
		lpDBHandler := NewDBHandlerWithConfig(v)
		err = lpDBHandler.Connect()
		if err != nil {
			FreeConn()
			panic(fmt.Sprintf("Init db conn failed: %v %v", v, err))
			//return errors.WithStack(err)
		}
		DBMap[k] = lpDBHandler

	}
	return err
}

func (p *DBHandler) GetConnection() (*gorm.DB, error) {
	err := p.Connect()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return p.DB, nil
}

func GetHandler(k string) (*DBHandler, error) {
	var err error
	var handler *DBHandler
	var ok bool
	if handler, ok = DBMap[k]; !ok {
		err = werrors.ERR_DB_MAP_KEY_NOT_FIND
		return nil, errors.Wrapf(err, "unknown k:"+k)
	}
	if !handler.Connected {
		err = handler.Connect()
		if err != nil {
			return nil, errors.WithStack(err)
		} else {
			return handler, nil
		}
	}
	return handler, nil
}

var DBMap = make(map[string]*DBHandler, 0)
