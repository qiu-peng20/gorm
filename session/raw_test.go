package session

import (
	"database/sql"
	"gorm/dialect"
)

type User struct {
	Name string
	Age  int
}

var (
	TestDB *sql.DB
	TestDial,_ = dialect.GetDialect("mysql")
)

func NewSession() *Session  {
	return New(TestDB, TestDial)
}
