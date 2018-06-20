package article

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"wncbb.cn/db"
	"wncbb.cn/define"
)

type Article struct {
	//gorm.Model
	Id        int64 `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
	Title     string
	Content   string
	UserId    int64
}

var readDbHandler *db.DBHandler
var writeDbHandler *db.DBHandler

func (Article) TableName() string {
	return "article"
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

func Migrate() error {
	conn, err := writeDbHandler.GetConnection()
	if err != nil {
		return err
	}

	if err := conn.AutoMigrate(&Article{}).Error; err != nil {
		return err
	}

	/*
	 * pg 看索引
	 * select * from pg_indexes;
	 * seelct * from pg_indexes where tablename='session'; --必须用单引号
	 * CREATE INDEX idx_session_type ON public.session USING btree (session, type)
	 */

	return nil
}

func GetById(id int64) (article *Article, err error) {
	conn, err := readDbHandler.GetConnection()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	article, err = getById(conn, id)
	err = errors.WithStack(err)
	return
}

func getById(conn *gorm.DB, id int64) (article *Article, err error) {
	article = &Article{}
	err = conn.Where("id = ?", id).First(article).Error
	err = errors.WithStack(err)
	return
}

func Create(article *Article) (err error) {
	if nil == article {
		err = errors.New("article is nil")
		return
	}
	var conn *gorm.DB
	conn, err = writeDbHandler.GetConnection()
	if err != nil {
		err = errors.WithStack(err)
		return
	}

	err = create(conn, article)
	err = errors.WithStack(err)
	return
}

func create(conn *gorm.DB, article *Article) (err error) {
	conn, err = writeDbHandler.GetConnection()
	if err != nil {
		err = errors.WithStack(err)
		return
	}

	err = conn.Create(article).Error
	err = errors.WithStack(err)
	return
}

func UpdateByMap(article *Article) (err error) {
	if article == nil {
		panic("model article UpdateByMap article=nil")
	}
	var conn *gorm.DB
	conn, err = writeDbHandler.GetConnection()
	if err != nil {
		err = errors.WithStack(err)
	}
	err = updateByMap(conn, article)
	err = errors.WithStack(err)
	return err
}

func updateByMap(conn *gorm.DB, article *Article) (err error) {
	if article == nil {
		panic("model article updateByMap article=nil")
	}

	updateInfo := make(map[string]interface{}, 0)
	updateInfo["title"] = article.Title
	updateInfo["content"] = article.Content
	err = conn.Model(&Article{}).Updates(updateInfo).Error
	err = errors.WithStack(err)
	return err
}

// func (p *ArticleRepository) Update(inData *Article) error {
// 	if nil == inData {
// 		return errors.New("article is nil")
// 	}

// 	conn, err := p.handler.GetConnection()
// 	if err != nil {
// 		return errors.WithStack(err)
// 	}

// 	return conn.Update(inData).Error
// }
