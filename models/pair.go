//go:generate kallax gen
//go:generate proteus:generate

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

	Markets []*Market `fk:"pair_id"`
}
