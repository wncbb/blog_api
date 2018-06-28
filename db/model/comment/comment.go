package comment

import (
	"blog_api/db"
	"blog_api/define"
	"blog_api/log"
	"context"
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

type Comment struct {
	//gorm.Model
	Id        int64 `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
	ArticleId int64      `sql:"index"`
	// 该条评论的UserId
	UserId  int64
	Content string `gorm:"type:varchar(255)"`
	ReplyId int64
}

var readDbHandler *db.DBHandler
var writeDbHandler *db.DBHandler

func (Comment) TableName() string {
	return "comment"
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
	if err = conn.AutoMigrate(&Comment{}).Error; err != nil {
		err = errors.WithStack(err)
		return
	}
	return
}

func Create(ctx context.Context, comment *Comment) (err error) {
	if nil == comment {
		err = errors.New("comment is nil")
		log.DefaultLogError("create comment failed, comment is nil", err)
		return
	}
	var conn *gorm.DB
	conn, err = writeDbHandler.GetConnection()
	if err != nil {
		err = errors.WithStack(err)
		log.DefaultLogError()
		return
	}
}
