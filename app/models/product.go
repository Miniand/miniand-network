package models

import (
	"github.com/coopernurse/gorp"
	"time"
)

type Product struct {
	Id          int64
	Created     int64
	Updated     int64
	Name        string
	Active      bool
	Description string
	SKU         string
}

func (p *Product) PreInsert(s gorp.SqlExecutor) error {
	p.Created = time.Now().UnixNano()
	p.Updated = p.Created
	return nil
}

func (p *Product) PreUpdate(s gorp.SqlExecutor) error {
	p.Updated = time.Now().UnixNano()
	return nil
}
