package session

import "testing"

var (
	user1 = &User{
		"tom",
		11,
	}
	user2 = &User{
		"jack",
		12,
	}
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


