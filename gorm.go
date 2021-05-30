package gorm

import (
	"database/sql"
	"gorm/log"
	"gorm/session"
)

type Engine struct {
	db *sql.DB
}

func NewEngine(driver, source string) (e *Engine, err error) {
	db, err := sql.Open(driver, source)
	if err != nil {
		log.Error(err)
		return
	}
	if err = db.Ping(); err != nil {
		log.Error(err)
		return
	}

	e = &Engine{db}
	log.Info("成功与数据库建立链接")
	return
}

func (e *Engine) Close() {
	if err := e.db.Close(); err != nil {
		log.Error(err)
	}
	log.Info("成功关闭数据库连接")
}

func (e *Engine) NewSession() *session.Session {
	return session.New(e.db)
}
