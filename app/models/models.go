package models

import (
	"database/sql"
	"github.com/coopernurse/gorp"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"github.com/robfig/revel"
	_ "github.com/ziutek/mymysql/mysql"
	"os"
	"path"
)

var dbInstance *gorp.DbMap

func Db() (*gorp.DbMap, error) {
	if dbInstance == nil {
		driverName, _ := revel.Config.String("db.driverName")
		dataSourceName, _ := revel.Config.String("db.dataSourceName")
		if driverName == "sqlite3" {
			// For SQLite, make the source relative to the app path and make
			// sure the dir exists
			dataSourceName = path.Join(revel.AppPath, dataSourceName)
			err := os.MkdirAll(path.Dir(dataSourceName), 0775)
			if err != nil {
				return nil, err
			}
		}
		db, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		var dialect gorp.Dialect
		switch driverName {
		case "sqlite3":
			dialect = gorp.SqliteDialect{}
		case "mysql":
			dialect = gorp.MySQLDialect{}
		case "postgres":
			dialect = gorp.PostgresDialect{}
		}
		dbInstance = &gorp.DbMap{
			Db:      db,
			Dialect: dialect,
		}
		dbInstance.AddTable(Product{}).SetKeys(true, "Id")
		dbInstance.AddTable(ProductVariant{}).SetKeys(true, "Id")
	}
	return dbInstance, nil
}
