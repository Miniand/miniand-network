package models

import (
	"github.com/Miniand/gorp"
	"github.com/robfig/revel"
	"regexp"
	"strconv"
	"time"
)

type Shop struct {
	Id               int64
	CreatedAt        int64
	UpdatedAt        int64
	Identifier       string
	Name             string
	Active           bool
	Hue              int
	ShortDescription string
	LongDescription  string
}

func (sh *Shop) PreInsert(s gorp.SqlExecutor) error {
	sh.CreatedAt = time.Now().UnixNano()
	sh.UpdatedAt = sh.CreatedAt
	return nil
}

func (sh *Shop) PreUpdate(s gorp.SqlExecutor) error {
	sh.UpdatedAt = time.Now().UnixNano()
	return nil
}

func (s *Shop) Validate(v *revel.Validation) {
	v.Required(s.Name)
	v.Required(s.Identifier)
	v.Match(s.Identifier, regexp.MustCompile(`^[a-z\-]+$`))
	v.Required(s.Hue)
}

func (s *Shop) ToStringMap() map[string]string {
	activeText := ""
	if s.Active {
		activeText = "true"
	}
	return map[string]string{
		"Id":               strconv.Itoa(int(s.Id)),
		"Identifier":       s.Identifier,
		"Name":             s.Name,
		"Active":           activeText,
		"Hue":              strconv.Itoa(s.Hue),
		"ShortDescription": s.ShortDescription,
		"LongDescription":  s.LongDescription,
	}
}

func FindShopByName(name string, exe gorp.SqlExecutor) *Shop {
	var shops []*Shop
	_, err := exe.Select(&shops, "SELECT * FROM Shop WHERE name=? LIMIT 1",
		name)
	if err != nil {
		panic(err)
	}
	if len(shops) > 0 {
		return shops[0]
	}
	return nil
}

func FindShopByIdentifier(identifier string, exe gorp.SqlExecutor) *Shop {
	var shops []*Shop
	_, err := exe.Select(&shops,
		"SELECT * FROM Shop WHERE identifier=? LIMIT 1", identifier)
	if err != nil {
		panic(err)
	}
	if len(shops) > 0 {
		return shops[0]
	}
	return nil
}
