package models

import (
	"github.com/coopernurse/gorp"
	"time"
)

type ShopProduct struct {
	Id        int64
	Created   int64
	Updated   int64
	ShopId    int64
	ProductId int64
}

func (sp *ShopProduct) PreInsert(s gorp.SqlExecutor) error {
	sp.Created = time.Now().UnixNano()
	sp.Updated = sp.Created
	return nil
}

func (sp *ShopProduct) PreUpdate(s gorp.SqlExecutor) error {
	sp.Updated = time.Now().UnixNano()
	return nil
}
