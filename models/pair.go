package models

import (
	"gopkg.in/src-d/go-kallax.v1"
)

type Pair struct {
	kallax.Model       `table:"pairs"`
	ID     kallax.ULID `pk:"autoincr"`
	Symbol string      `kallax:"symbol"`
	Base   *Asset      `fk:"base_id,inverse"`
	Quote  *Asset      `fk:"quote_id,inverse"`

	Markets     []*Market     `fk:"pair_id"`
	Differences []*Difference `fk:"pair_id"`
	Orders      []*Order      `fk:"pair_id"`
	Prices      []*Price      `fk:"pair_id"`
	Trades      []*Trade      `fk:"pair_id"`
}
