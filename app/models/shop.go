package models

import (
	"github.com/coopernurse/gorp"
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
