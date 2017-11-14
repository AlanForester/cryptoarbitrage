//go:generate kallax gen
//go:generate proteus:generate

package models

import (
	"gopkg.in/src-d/go-kallax.v1"
)

type Price struct {
	kallax.Model           `table:"prices"`
	kallax.Timestamps
	ID         kallax.ULID `pk:"autoincr"`
	PairId     *Pair       `fk:"pair_id,inverse"`
	ExchangeId *Exchange   `fk:"exchange_id,inverse"`
	MarketId   *Market     `fk:"market_id,inverse"`
	Price      float32     `kallax:"price"`

	PairSymbol     string `kallax:",inline"`
	ExchangeSymbol string `kallax:",inline"`
}
