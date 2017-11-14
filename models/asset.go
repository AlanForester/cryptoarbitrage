//go:generate kallax gen
//go:generate proteus:generate
package models

import (
	"gopkg.in/src-d/go-kallax.v1"
)

type Asset struct {
	kallax.Model       `table:"assets"`
	ID     kallax.ULID `pk:"autoincr"`
	Symbol string      `kallax:"symbol"`
	Name   string      `kallax:"name"`
	IsFiat bool        `kallax:"is_fiat"`

	Markets    []*Market `fk:"exchange_id"`
	BasePairs  []*Pair   `fk:"base_id"`
	QuotePairs []*Pair   `fk:"quote_id"`
}
