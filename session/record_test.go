package session

import (
	"testing"
)

var (
	user1 = &User{"Tom",11}
	user2 = &User{"Jack",12}
	user3 = &User{"Baby",13}
)

func testRecordInit(t *testing.T) *Session {
	t.Helper()
	s := NewSession().Model(&User{})
	err1 := s.DropTable()
	err2 := s.CreateTable()
	_, err3 := s.Insert(user1, user2)
	if err1 != nil || err2 != nil || err3 != nil {
		t.Fatal("failded to create record")
	}
	return s
}

func TestSession_Insert(t *testing.T) {
	s := testRecordInit(t)
	result, err := s.Insert(user3)
	if err != nil || result != 1 {
		t.Fatal("fatal to create data")
	}
}

func TestSession_Find(t *testing.T) {
	s := testRecordInit(t)
	var users []User
	if err := s.Find(&users); err != nil || len(users) != 2 {
		t.Fatal("fatal to query all")
	}
}

func TestSession_Limit(t *testing.T) {
	s := testRecordInit(t)
	var users []User
	err := s.Limit(1).Find(&users)
	if err != nil || len(users) != 1 {
		t.Fatal("这里是limit的错误")
	}
 }




