package models

import (
	"github.com/coopernurse/gorp"
	"github.com/robfig/revel"
	"time"
)

type Product struct {
	Id          int64
	CreatedAt   int64
	UpdatedAt   int64
	Name        string
	Active      bool
	Description string
}

func (p *Product) PreInsert(s gorp.SqlExecutor) error {
	p.CreatedAt = time.Now().UnixNano()
	p.UpdatedAt = p.CreatedAt
	return nil
}

func (p *Product) PreUpdate(s gorp.SqlExecutor) error {
	p.UpdatedAt = time.Now().UnixNano()
	return nil
}

func (p *Product) Validate(v *revel.Validation) {
	v.Check(p.Name, revel.Required{}, revel.MinSize{1})
}
