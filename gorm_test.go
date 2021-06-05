package gorm

import (
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"gorm/session"
	"testing"
)

type User struct {
	Name string `gorm:"NOT NULL"`
	Age  int
}

func OpenDB(t *testing.T) *Engine {
	t.Helper()
	engine, err := NewEngine("mysql", "root:123456@tcp(127.0.0.1:3306)/gorm?charset=utf8")
	if err != nil {
		t.Fatal("OpenDB error",err)
	}
	return engine
}

func TestEngine_Transaction(t *testing.T) {
	t.Run("rollBack", func(t *testing.T) {
		transactionRollback(t)
	})
	t.Run("commit", func(t *testing.T) {
		transactionCommit(t)
	})
}

func transactionRollback(t *testing.T)  {
	engine := OpenDB(t)
	defer engine.Close()
	s := engine.NewSession()
	_ = s.Model(&User{}).DropTable()
	_, err := engine.Transaction(func(session *session.Session) (result interface{}, err error) {
		_ = s.Model(&User{}).CreateTable()
		_, err = s.Insert(&User{"jack", 18})
		return nil,errors.New("Error")
	})
	if err == nil || s.HasTable() {
		t.Fatal("fail to rollback")
	}
}

func transactionCommit(t *testing.T)  {
	engine := OpenDB(t)
	defer engine.Close()
	s := engine.NewSession()
	_ = s.Model(&User{}).DropTable()
	_, err := engine.Transaction(func(s *session.Session) (result interface{}, err error) {
		_ = s.Model(&User{}).CreateTable()
		_, err = s.Insert(&User{"tom", 29})
		return
	})
	u := &User{}
	_ = s.First(u)
	if err != nil || u.Name != "tom" {
		t.Fatal("fatal commit")
	}
}
