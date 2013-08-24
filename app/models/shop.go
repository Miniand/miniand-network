package models

import (
	"github.com/coopernurse/gorp"
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

func (sh *Shop) Validate(v *revel.Validation) {
	v.Check(sh.Name, revel.Required{}, revel.MinSize{1})
	v.Check(sh.Hue, revel.Required{})
	v.Check(sh.Identifier, revel.Required{}, revel.MinSize{1}, revel.Match{
		regexp.MustCompile(`^[a-z]+$`),
	})
}

func (sh *Shop) ToStringMap() map[string]string {
	return map[string]string{
		"Id":               strconv.Itoa(int(sh.Id)),
		"Identifier":       sh.Identifier,
		"Name":             sh.Name,
		"Hue":              strconv.Itoa(sh.Hue),
		"ShortDescription": sh.ShortDescription,
		"LongDescription":  sh.LongDescription,
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
