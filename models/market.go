package models

import (
	"gopkg.in/src-d/go-kallax.v1"
)

type Market struct {
	kallax.Model         `table:"markets"`
	ID       kallax.ULID `pk:"autoincr"`
	Pair     *Pair       `fk:"pair_id,inverse"`
	Exchange *Exchange   `fk:"exchange_id,inverse"`
	IsActive bool        `kallax:"is_active"`

	Orders []*Order `fk:"market_id"`
	Prices []*Price `fk:"market_id"`
	Trades []*Trade `fk:"market_id"`
}
