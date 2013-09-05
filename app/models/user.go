package models

import (
	"github.com/coopernurse/gorp"
	"time"
)

type User struct {
	Id        int64
	CreatedAt int64
	UpdatedAt int64
	Email     string
	Password  string
	Name      string
}

func (u *User) PreInsert(s gorp.SqlExecutor) error {
	u.CreatedAt = time.Now().UnixNano()
	u.UpdatedAt = u.CreatedAt
	return nil
}

func (u *User) PreUpdate(s gorp.SqlExecutor) error {
	u.UpdatedAt = time.Now().UnixNano()
	return nil
}
