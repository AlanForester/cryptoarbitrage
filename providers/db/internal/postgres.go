package internal

import (
	_ "github.com/lib/pq"
	"database/sql"
	"fmt"
	. "cryptoarbitrage/helpers"
	. "cryptoarbitrage/providers/config"
)

var Postgres *postgresModel

type postgresModel struct {
	Instance *sql.DB
	Error    error
}

func (s *postgresModel) connect() {
	s.Instance, s.Error = sql.Open("postgres", s.getDSN())
	Error.Check(s.Error)
	s.ping()
}

func (s *postgresModel) ping() {
	if s.Error = s.Instance.Ping(); s.Error != nil {
		Error.Check(s.Error)
	}
}

func (s *postgresModel) getDSN() string {
	c := Config.Storage.Postgres
	return fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", c.User, c.Pass, c.Host, c.Name)
}

func init() {
	Postgres = &postgresModel{}
	Postgres.connect()
}
