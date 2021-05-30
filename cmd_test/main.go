package main

import "gorm"

func main()  {
	engine ,_ := gorm.NewEngine("mysql", "root:123456@tcp(127.0.0.1:3306)/gorm?charset=utf8")
	defer engine.Close()

	s := engine.NewSession()
	s.Raw("")
}
