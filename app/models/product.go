package models

import (
	"github.com/Miniand/gorp"
	"github.com/robfig/revel"
	"strconv"
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
	v.Required(p.Name)
	v.Required(p.Description)
}

func (p *Product) ToStringMap() map[string]string {
	activeText := ""
	if p.Active {
		activeText = "true"
	}
	return map[string]string{
		"Id":          strconv.Itoa(int(p.Id)),
		"Name":        p.Name,
		"Active":      activeText,
		"Description": p.Description,
	}
}
