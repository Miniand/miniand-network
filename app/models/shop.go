package models

import (
	"github.com/coopernurse/gorp"
	"github.com/robfig/revel"
	"time"
)

type Shop struct {
	Id         int64
	Created    int64
	Updated    int64
	Identifier string
	Name       string
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
}

func FindShopByName(name string, exe gorp.SqlExecutor) *Shop {
	var shops []*Shop
	_, err := exe.Select(&shops, "SELECT * FROM Shop WHERE name=? LIMIT 1", name)
	if err != nil {
		panic(err)
	}
	if len(shops) > 0 {
		return shops[0]
	}
	return nil
}
