//go:generate kallax gen
//go:generate proteus:generate

package models

import (
	"gopkg.in/src-d/go-kallax.v1"
)

type Trade struct {
	kallax.Model           `table:"trades"`
	ID         kallax.ULID `pk:"autoincr"`
	kallax.Timestamps
	UserId     User        `fk:"user_id,inverse"`
	ExchangeId Exchange    `fk:"exchange_id,inverse"`
	PairId     Pair        `fk:"pair_id,inverse"`
	OrderId    Order       `kallax:"order_id"`
	Type       string      `kallax:"type"`
	Volume     float32     `kallax:"volume"`
	Price      float32     `kallax:"price"`
}
