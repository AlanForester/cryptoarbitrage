//go:generate kallax gen
//go:generate proteus:generate

package models

import (
	"gopkg.in/src-d/go-kallax.v1"
)

type Exchange struct {
	kallax.Model `table:"exchanges" pk:"id,autoincr"`
	ID   kallax.ULID
	Code string
}
