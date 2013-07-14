package models

import (
	"database/sql"
	"github.com/coopernurse/gorp"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"github.com/robfig/revel"
	_ "github.com/ziutek/mymysql/mysql"
	"path"
)

var dbInstance *gorp.DbMap

func Db() (*gorp.DbMap, error) {
	if dbInstance == nil {
		driverName, _ := revel.Config.String("db.driverName")
		dataSourceName, _ := revel.Config.String("db.dataSourceName")
		// For SQLite, make the source relative to the app path
		if driverName == "sqlite3" {
			dataSourceName = path.Join(revel.AppPath, dataSourceName)
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
