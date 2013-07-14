package models

import (
	"github.com/coopernurse/gorp"
	"time"
)

type ProductVariant struct {
	Id          int64
	Created     int64
	Updated     int64
	ProductId   int64
	Name        string
	Active      bool
	Price       int64
	Description string
	SKU         string
}

func (pv *ProductVariant) PreInsert(s gorp.SqlExecutor) error {
	pv.Created = time.Now().UnixNano()
	pv.Updated = pv.Created
	return nil
}

func (pv *ProductVariant) PreUpdate(s gorp.SqlExecutor) error {
	pv.Updated = time.Now().UnixNano()
	return nil
}
