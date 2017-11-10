//go:generate kallax gen
//proteus:generate
package models

import (
	"gopkg.in/src-d/go-kallax.v1"
)

type Pair struct {
	kallax.Model `table:"pairs" pk:"id"`
	ID     kallax.ULID
	Symbol string
	BaseId   string
	QuoteId  string
}
