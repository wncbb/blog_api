package session

import (
	"errors"
	"time"

	"wncbb.cn/db/pg"
)

type Session struct {
	//gorm.Model
	ID        int64 `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
	UserID    int64
	Type      int8 //pc, mobile, h5
	Session   string
}

func (Session) TableName() string {
	return "session"
}

type SessionRepository struct {
	handler *pg.DBHandler
}

func NewSessionRepository(handler *pg.DBHandler) *SessionRepository {
	return &SessionRepository{
		handler: handler,
	}
}

func (p *SessionRepository) Migrate() error {
	conn, err := p.handler.GetConnection()
	if err != nil {
		return err
	}
	/*
		//for mysql
		tableOptions := "ENGINE=InnoDB DEFAULT CHARSET=utf8"
		if err := conn.Set("gorm:table_options", tableOptions).AutoMigrate(Session{}).Error; err != nil {
			return err
		}
	*/
	if err := conn.AutoMigrate(Session{}).Error; err != nil {
		return err
	}
	if err := conn.Model(Session{}).AddIndex("idx_session_type", "session", "type").Error; err != nil {
		return err
	}
	/*
	 * pg 看索引
	 * select * from pg_indexes;
	 * seelct * from pg_indexes where tablename='session'; --必须用单引号
	 * CREATE INDEX idx_session_type ON public.session USING btree (session, type)
	 */
	if err := conn.Model(Session{}).AddIndex("idx_session_type", "session", "type").Error; err != nil {
		return err
	}

	return nil
}

func (p *SessionRepository) Create(inPtr *Session) error {
	if nil == inPtr {
		return errors.New("user is nil")
	}

	conn, err := p.handler.GetConnection()
	if err != nil {
		return err
	}

	return conn.Create(inPtr).Error
}

//func (p *SessionRepository) GetBy
