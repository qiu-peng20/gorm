package session

type User struct {
	Name string `gorm:"NOT NULL"`
	Age  int
}

type Account struct {
	ID       int
	PassWord string
}
