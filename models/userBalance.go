//go:generate kallax gen
//go:generate proteus:generate

package models

import (
	"gopkg.in/src-d/go-kallax.v1"
)

type UserBalance struct {
	kallax.Model         `table:"users" pk:"id,autoincr"`
	ID       kallax.ULID `pk:"autoincr"`
	User     *User       `fk:"user_id,inverse"`
	Exchange *Exchange   `fk:"exchange_id,inverse"`
	Asset    string
	Volume   float32
}
