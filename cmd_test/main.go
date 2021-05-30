package main

import (
	"fmt"
	"gorm"

	_ "github.com/go-sql-driver/mysql"
)

func main()  {
	engine ,_ := gorm.NewEngine("mysql", "root:123456@tcp(127.0.0.1:3306)/gorm?charset=utf8")
	defer engine.Close()

	s := engine.NewSession()
	_, _ = s.Raw("DROP TABLE IF EXISTS User").Exec()
	_, _ = s.Raw("CREATE TABLE User(name CHAR(10))").Exec()
	_, _ = s.Raw("CREATE TABLE User(name CHAR(10))").Exec()
	result, _ := s.Raw("INSERT INTO User(name) value (?), (?)", "tom", "Sam").Exec()
	count, _ := result.RowsAffected()
	fmt.Print(count)
}
