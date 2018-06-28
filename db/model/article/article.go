package article

import (
	"context"
	"fmt"
	"time"

	"blog_api/db"
	"blog_api/define"

	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
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

type ArticleModel struct {
	ctx context.Context
}

func NewModel(ctx context.Context) *ArticleModel {
	return &ArticleModel{
		ctx: ctx,
	}
}

func (p *ArticleModel) GetList(offset, limit int64) ([]*Article, error) {
	conn, err := readDbHandler.GetConnection()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	list, err := p.getList(conn, offset, limit)
	return list, errors.WithStack(err)
}

func (p *ArticleModel) getList(conn *gorm.DB, offset, limit int64) ([]*Article, error) {
	articles := make([]*Article, 0, limit)
	err := conn.Model(&Article{}).Order("created_at").Offset(offset).Limit(limit).Find(&articles).Error
	if err != nil {
		err = errors.WithStack(err)
	}
	return articles, err
}

func (p *ArticleModel) GetNum() (int64, error) {
	conn, err := readDbHandler.GetConnection()
	if err != nil {
		return 0, errors.WithStack(err)
	}
	num, err := p.getNum(conn)
	return num, errors.WithStack(err)
}

func (p *ArticleModel) getNum(conn *gorm.DB) (int64, error) {
	var num int64
	err := conn.Model(&Article{}).Count(&num).Error
	return num, errors.WithStack(err)
}

func (p *ArticleModel) GetById(id int64) (article *Article, err error) {
	conn, err := readDbHandler.GetConnection()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	article, err = p.getById(conn, id)
	err = errors.WithStack(err)
	return
}

func (p *ArticleModel) getById(conn *gorm.DB, id int64) (article *Article, err error) {
	article = &Article{}
	err = conn.Where("id = ?", id).First(article).Error
	err = errors.WithStack(err)
	return
}

func (p *ArticleModel) Create(article *Article) (err error) {
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

	err = p.create(conn, article)
	err = errors.WithStack(err)
	return
}

func (p *ArticleModel) create(conn *gorm.DB, article *Article) (err error) {
	conn, err = writeDbHandler.GetConnection()
	if err != nil {
		err = errors.WithStack(err)
		return
	}

	err = conn.Create(article).Error
	err = errors.WithStack(err)
	return
}

func (p *ArticleModel) UpdateByMap(article *Article) (err error) {
	if article == nil {
		panic("model article UpdateByMap article=nil")
	}
	var conn *gorm.DB
	conn, err = writeDbHandler.GetConnection()
	if err != nil {
		err = errors.WithStack(err)
	}
	err = p.updateByMap(conn, article)
	err = errors.WithStack(err)
	return err
}

func (p *ArticleModel) updateByMap(conn *gorm.DB, article *Article) (err error) {
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
