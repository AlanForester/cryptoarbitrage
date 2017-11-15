package models

import (
	"gopkg.in/src-d/go-kallax.v1"
)

type Price struct {
	kallax.Model              `table:"prices"`
	kallax.Timestamps
	ID       kallax.NumericID `pk:"autoincr"`
	Pair     *Pair            `fk:"pair_id,inverse"`
	Exchange *Exchange        `fk:"exchange_id,inverse"`
	Market   *Market          `fk:"market_id,inverse"`
	Price    float32          `kallax:"price"`

	PairSymbol     string `kallax:",inline"`
	ExchangeSymbol string `kallax:",inline"`
}
