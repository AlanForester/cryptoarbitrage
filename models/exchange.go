//go:generate kallax gen
//go:generate proteus:generate

package models

import (
	"gopkg.in/src-d/go-kallax.v1"
)

type Exchange struct {
	kallax.Model         `table:"exchanges"`
	ID       kallax.ULID `pk:"autoincr"`
	Code     string      `kallax:"code"`
	IsActive bool        `kallax:"is_active"`

	Markets []*Market        `fk:"exchange_id"`
	Assets  []*ExchangeAsset `fk:"exchange_id"`
}
