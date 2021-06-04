package session

import (
	"gorm/log"
	"testing"
)

func (a *Account) AfterQuery(s *Session) error  {
	log.Info("after query", a)
	a.PassWord = "******"
	return nil
}

func TestSession_CallMethod(t *testing.T) {
	s := NewSession().Model(&Account{})
	_ = s.DropTable()
	_ = s.CreateTable()
	_,_ = s.Insert(&Account{1,"12345"})
	u := &Account{}
	err := s.First(&u)
	if err != nil || u.PassWord != "******" {
		t.Fatal("hooks error")
	}
}