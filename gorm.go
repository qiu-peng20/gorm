package gorm

import (
	"database/sql"
	"gorm/dialect"
	"gorm/log"
	"gorm/session"
)

type Engine struct {
	db *sql.DB
	dialect dialect.Dialect
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
	dialect2, ok := dialect.GetDialect(driver)
	if !ok {
		log.Errorf("dialect %s is not found",driver)
		return
	}
	e = &Engine{db, dialect2}
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
	return session.New(e.db,e.dialect)
}

type TxFunc func(*session.Session) (interface{}, error)

func (e *Engine) Transaction(f TxFunc) (result interface{}, err error)  {
	s := e.NewSession()
	if err = s.Begin(); err != nil {
		return nil, err
	}
	defer func() {
		if p := recover(); p != nil {
			_ = s.Rollback()
			panic(p)
		}else if err != nil {
			_ = s.Rollback()
		}else {
			err = s.Commit()
			defer func() {
				if err != nil {
					_ = s.Rollback()
				}
			}()
		}
	}()
	return f(s)
}
