package models

import (
	"gopkg.in/src-d/go-kallax.v1"
)

type Exchange struct {
	kallax.Model         `table:"exchanges"`
	ID       kallax.ULID `pk:"autoincr"`
	Symbol   string      `kallax:"symbol"`
	Name     string      `kallax:"name"`
	IsActive bool        `kallax:"is_active"`

	Markets          []*Market        `fk:"exchange_id"`
	Assets           []*ExchangeAsset `fk:"exchange_id"`
	BaseDifferences  []*Difference    `fk:"base_id"`
	QuoteDifferences []*Difference    `fk:"quote_id"`
	Orders           []*Order         `fk:"exchange_id"`
	Prices           []*Price         `fk:"exchange_id"`
	Trades           []*Trade         `fk:"exchange_id"`
	Balances         []*UserBalance   `fk:"exchange_id"`
}
