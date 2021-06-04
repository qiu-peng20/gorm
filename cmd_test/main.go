package main

import (
	_ "github.com/go-sql-driver/mysql"
	"gorm"
	"gorm/log"
)

type User struct {
	Name string `gorm:"NOT NULL"`
	Age  int
}

var (
	user1 = &User{
		Name: "tom",
		Age:  11,
	}
	user2 = &User{
		Name: "jack",
		Age:  10,
	}
)

func main() {
	engine, _ := gorm.NewEngine("mysql", "root:123456@tcp(127.0.0.1:3306)/gorm?charset=utf8")
	defer engine.Close()

	s := engine.NewSession().Model(&User{})
	_ = s.DropTable()
	_ = s.CreateTable()
	_, err := s.Insert(user1, user2)
	log.Error(err)
}
