//go:generate kallax gen
//go:generate proteus:generate
package models

import (
	"gopkg.in/src-d/go-kallax.v1"
)

type Asset struct {
	kallax.Model `table:"assets" pk:"id,autoincr"`
	ID     kallax.ULID
	Symbol string
	Name   string
	IsFiat bool
}