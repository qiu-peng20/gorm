package session

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"gorm/dialect"
	"os"
	"testing"
)

var (
	TestDB      *sql.DB
	TestDial, _ = dialect.GetDialect("mysql")
)

func TestMain(m *testing.M)  {
	dest := "root:123456@tcp(127.0.0.1:3306)/gorm"
	TestDB, _ = sql.Open("mysql", dest)
	code := m.Run()
	_ = TestDB.Close()
	os.Exit(code)
}

func NewSession() *Session {
	return New(TestDB, TestDial)
}
