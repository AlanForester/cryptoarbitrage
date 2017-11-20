package db

import (
	_ "github.com/lib/pq"
	"database/sql"
	. "crypto-arbitrage/providers/db/internal"
)

var DB *dbModel

type dbModel struct {
	SQL *sql.DB
	SQLError      error
}

func init() {
	DB = &dbModel{}
	DB.SQL = Postgres.Instance
	DB.SQLError = Postgres.Error
}


