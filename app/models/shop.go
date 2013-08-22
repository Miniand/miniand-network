package models

import (
	"github.com/coopernurse/gorp"
	"github.com/robfig/revel"
	"regexp"
	"time"
)

type Shop struct {
	Id         int64
	Created    int64
	Updated    int64
	Identifier string
	Name       string
	Hue        int
}

func (sh *Shop) PreInsert(s gorp.SqlExecutor) error {
	sh.Created = time.Now().UnixNano()
	sh.Updated = sh.Created
	return nil
}

func (sh *Shop) PreUpdate(s gorp.SqlExecutor) error {
	sh.Updated = time.Now().UnixNano()
	return nil
}

func (sh *Shop) Validate(v *revel.Validation) {
	v.Check(sh.Name, revel.Required{}, revel.MinSize{1})
	v.Check(sh.Hue, revel.Required{})
	v.Check(sh.Identifier, revel.Required{}, revel.MinSize{1}, revel.Match{
		regexp.MustCompile(`^[a-z]+$`),
	})
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
