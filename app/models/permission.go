package models

import (
	"github.com/coopernurse/gorp"
	"time"
)

type Permission struct {
	Id        int64
	CreatedAt int64
	UpdatedAt int64
	Name      string
	UserId    int64
	ShopId    int64
}

func (p *Permission) PreInsert(s gorp.SqlExecutor) error {
	p.CreatedAt = time.Now().UnixNano()
	p.UpdatedAt = p.CreatedAt
	return nil
}

func (p *Permission) PreUpdate(s gorp.SqlExecutor) error {
	p.UpdatedAt = time.Now().UnixNano()
	return nil
}
