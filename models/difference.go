//go:generate kallax gen
//go:generate proteus:generate

package models

import (
	"gopkg.in/src-d/go-kallax.v1"
)

type Difference struct {
	kallax.Model                   `table:"differences"`
	ID            kallax.ULID      `pk:"autoincr"`
	kallax.Timestamps
	Pair          *Pair            `fk:"pair_id,inverse"`
	BaseExchange  *Exchange        `fk:"base_id,inverse"`
	QuoteExchange *Exchange        `fk:"quote_id,inverse"`
	Delta         float32          `kallax:"delta"`
	Markets       []*Market        `fk:"exchange_id"`
	Assets        []*ExchangeAsset `fk:"exchange_id"`
}
