package user

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"wncbb.cn/db"
	"wncbb.cn/define"
	"wncbb.cn/werrors"
)

type UserCore struct {
	//gorm.Model
	Id        int64 `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
	Account   string     `gorm:"type:varchar(100);unique_index"`
	Password  string     `gorm:"type:varchar(100)"`
	Secret    string     `gorm:"type:varchar(10);"`
}

var readDbHandler *db.DBHandler
var writeDbHandler *db.DBHandler

func (UserCore) TableName() string {
	return "user_core"
}

func Init() (err error) {
	readDbHandler, err = db.GetHandler(define.READ)
	if err != nil {
		panic(fmt.Sprintf("Init database %s faild, error:%v", define.READ, err))
	}
	writeDbHandler, err = db.GetHandler(define.WRITE)
	if err != nil {
		panic(fmt.Sprintf("Init database %s faild, error:%v", define.WRITE, err))
	}
	return
}

func Migrate() (err error) {
	var conn *gorm.DB
	conn, err = writeDbHandler.GetConnection()
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	/*
		//for mysql
		tableOptions := "ENGINE=InnoDB DEFAULT CHARSET=utf8"
		if err := conn.Set("gorm:table_options", tableOptions).AutoMigrate(User{}).Error; err != nil {
			return err
		}
	*/
	if err = conn.AutoMigrate(UserCore{}).Error; err != nil {
		err = errors.WithStack(err)
		return
	}

	return
}

func Create(userCore *UserCore) (err error) {
	if nil == userCore {
		err = errors.New("userCore is nil")
		return
	}
	var conn *gorm.DB
	conn, err = writeDbHandler.GetConnection()
	if err != nil {
		err = errors.WithStack(err)
		return
	}

	return create(conn, userCore)
}

func create(conn *gorm.DB, userCore *UserCore) (err error) {
	if nil == userCore {
		err = errors.New("userCore is nil")
		return
	}

	query := conn.Create(userCore)
	err = query.Error
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	return
}

func GetById(id int64) (userCore *UserCore, err error) {
	var conn *gorm.DB
	conn, err = readDbHandler.GetConnection()
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	return getById(conn, id)
}

func getById(conn *gorm.DB, id int64) (userCore *UserCore, err error) {
	query := conn.Where("id = ?", id).First(userCore)
	err = query.Error
	if err != nil {
		err = errors.WithStack(err)
		return
	}

	return
}

func GetByAccount(account string) (userCore *UserCore, err error) {
	var conn *gorm.DB
	conn, err = readDbHandler.GetConnection()
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	return getByAccount(conn, account)
}

func getByAccount(conn *gorm.DB, account string) (userCore *UserCore, err error) {
	userCore = &UserCore{}
	query := conn.Where("account = ?", account).First(userCore)
	err = query.Error
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	return
}

func UpdateByMap(userCore *UserCore) (err error) {
	if nil == userCore {
		err = werrors.ERR_FUNC_PARAMS_POINT_NIL
		err = errors.Wrapf(err, "UpdateViaMap parameter user is nil")
		err = errors.WithStack(err)
		return err
	}
	var conn *gorm.DB
	conn, err = writeDbHandler.GetConnection()
	if err != nil {
		err = errors.WithStack(err)
		return
	}

	updateInfo := make(map[string]interface{}, 0)
	updateInfo["account"] = userCore.Account
	updateInfo["password"] = userCore.Password
	updateInfo["secret"] = userCore.Secret

	err = conn.Model(&UserCore{}).Updates(updateInfo).Error
	err = errors.WithStack(err)
	return
}

func UpdateBySave(userCore *UserCore) (err error) {
	if nil == userCore {
		err = errors.New("userCore is nil")
		return
	}

	var conn *gorm.DB

	conn, err = writeDbHandler.GetConnection()
	if err != nil {
		err = errors.WithStack(err)
		return
	}

	err = conn.Save(userCore).Error
	return
}

func IsExistAccount(account string) (isExist bool, err error) {
	var conn *gorm.DB

	conn, err = writeDbHandler.GetConnection()
	if err != nil {
		isExist = false
		err = errors.WithStack(err)
		return
	}

	return isExistAccount(conn, account)
}

func isExistAccount(conn *gorm.DB, account string) (isExist bool, err error) {
	var num int64
	err = conn.Model(&UserCore{}).Where("account=?", account).Count(&num).Error
	if err != nil {
		isExist = false
		err = errors.WithStack(err)
		return
	}
	if num != 1 {
		isExist = false
		return
	}
	isExist = true

	return
}

func Delete(userCore *UserCore) (err error) {
	if nil == userCore {
		err = errors.New("user is nil")
		return
	}

	var conn *gorm.DB
	conn, err = writeDbHandler.GetConnection()
	if err != nil {
		err = errors.WithStack(err)
		return
	}

	if userCore.Id > 0 {
		err = conn.Where("id = ?", userCore.Id).Delete(&UserCore{}).Error
		err = errors.WithStack(err)
		return
	}

	if userCore.Account != "" {
		err = conn.Where("account = ?", userCore.Account).Delete(&UserCore{}).Error
		err = errors.WithStack(err)
		return
	}

	return
}

// func ListAll(offset, count int64) (userCoreList []*UserCore, err error) {
// 	userCoreList = make([]*UserCore, 0, count)
// 	var conn *gorm.DB
// 	conn, err = writeDbHandler.GetConnection()
// 	if err != nil {
// 		err = errors.WithStack(err)
// 		return
// 	}

// 	err = conn.Order("id").Offset(offset).Count(count).Find(&userCoreList).Error
// 	err = errors.WithStack(err)
// 	return
// }

func ListAll() (userCoreList []*UserCore, err error) {
	userCoreList = make([]*UserCore, 0)
	var conn *gorm.DB
	conn, err = writeDbHandler.GetConnection()
	if err != nil {
		err = errors.WithStack(err)
		return
	}

	err = conn.Order("id").Find(&userCoreList).Error
	err = errors.WithStack(err)
	return
}

func deleteById(conn *gorm.DB, id int64) (err error) {
	err = conn.Delete(&UserCore{Id: id}).Error
	return errors.WithStack(err)
}
