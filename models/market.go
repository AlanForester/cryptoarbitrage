//go:generate kallax gen
//go:generate proteus:generate

package models

import (
	"gopkg.in/src-d/go-kallax.v1"
)

type Market struct {
	kallax.Model           `table:"pairs"`
	ID         kallax.ULID `pk:"autoincr"`
	PairId     *Pair       `fk:"pair_id,inverse"`
	ExchangeId *Exchange   `fk:"exchange_id,inverse"`
	IsActive   bool        `kallax:"is_active"`
}
