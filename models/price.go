//go:generate kallax gen
//go:generate proteus:generate

package models

import (
	"gopkg.in/src-d/go-kallax.v1"
)

type Price struct {
	kallax.Model            `table:"prices" pk:"id,autoincr"`
	kallax.Timestamps
	ID             kallax.ULID
	PairId         Pair     `fk:"pair_id,inverse"`
	ExchangeId     Exchange `fk:"exchange_id,inverse"`
	Price          float32

	PairSymbol     string   `kallax:",inline"`
	ExchangeSymbol string   `kallax:",inline"`
}
