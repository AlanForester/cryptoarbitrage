package models

import (
	"gopkg.in/src-d/go-kallax.v1"
)

type Asset struct {
	kallax.Model            `table:"assets"`
	ID     kallax.NumericID `pk:"autoincr"`
	Symbol string           `kallax:"symbol"`
	Name   string           `kallax:"name"`
	IsFiat bool             `kallax:"is_fiat"`

	BasePairs  []*Pair        `fk:"base_id"`
	QuotePairs []*Pair        `fk:"quote_id"`
	Balances   []*UserBalance `fk:"asset_id"`
}
