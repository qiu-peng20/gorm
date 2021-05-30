package main

import (
	"gorm"
	"gorm/log"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Name string `gorm:"NOT NULL"`
	Age  int
}

func main() {
	engine, _ := gorm.NewEngine("mysql", "root:123456@tcp(127.0.0.1:3306)/gorm?charset=utf8")
	defer engine.Close()

	s := engine.NewSession().Model(&User{})
	_ = s.DropTable()
	_ = s.CreateTable()
	if !s.HasTable() {
		log.Errorf("this table is error")
	}
}
