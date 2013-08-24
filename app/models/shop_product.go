package models

import (
	"github.com/coopernurse/gorp"
	"time"
)

type ShopProduct struct {
	Id        int64
	CreatedAt int64
	UpdatedAt int64
	ShopId    int64
	ProductId int64
}

func (sp *ShopProduct) PreInsert(s gorp.SqlExecutor) error {
	sp.CreatedAt = time.Now().UnixNano()
	sp.UpdatedAt = sp.CreatedAt
	return nil
}

func (sp *ShopProduct) PreUpdate(s gorp.SqlExecutor) error {
	sp.UpdatedAt = time.Now().UnixNano()
	return nil
}
