package models

import (
	"gopkg.in/src-d/go-kallax.v1"
)

type Difference struct {
	kallax.Model                   `table:"differences"`
	ID            kallax.NumericID `pk:"autoincr"`
	kallax.Timestamps
	Pair          *Pair            `fk:"pair_id,inverse"`
	BaseExchange  *Exchange        `fk:"base_id,inverse"`
	QuoteExchange *Exchange        `fk:"quote_id,inverse"`
	Delta         float32          `kallax:"delta"`
}
