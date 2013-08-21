package models

import (
	"github.com/coopernurse/gorp"
	"time"
)

type Permission struct {
	Id      int64
	Created int64
	Updated int64
	Name    string
	UserId  int64
	ShopId  int64
}

func (p *Permission) PreInsert(s gorp.SqlExecutor) error {
	p.Created = time.Now().UnixNano()
	p.Updated = p.Created
	return nil
}

func (p *Permission) PreUpdate(s gorp.SqlExecutor) error {
	p.Updated = time.Now().UnixNano()
	return nil
}
