//go:generate kallax gen
//go:generate proteus:generate

package models

import (
	"gopkg.in/src-d/go-kallax.v1"
)

type Pair struct {
	kallax.Model `table:"pairs" pk:"id,autoincr"`
	ID     kallax.ULID
	Symbol string
	BaseId   *Asset `fk:"base_id,inverse"`
	QuoteId  *Asset `fk:"quote_id,inverse"`
}
