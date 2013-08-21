package models

import (
	"github.com/coopernurse/gorp"
	"time"
)

type User struct {
	Id       int64
	Created  int64
	Updated  int64
	Email    string
	Password string
	Name     string
}

func (u *User) PreInsert(s gorp.SqlExecutor) error {
	u.Created = time.Now().UnixNano()
	u.Updated = u.Created
	return nil
}

func (u *User) PreUpdate(s gorp.SqlExecutor) error {
	u.Updated = time.Now().UnixNano()
	return nil
}
