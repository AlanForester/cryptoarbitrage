//go:generate kallax gen
//go:generate proteus:generate

package models

import (
	"gopkg.in/src-d/go-kallax.v1"
	"time"
)

type UserBalance struct {
	kallax.Model              `table:"users" pk:"id,autoincr"`
	ID          kallax.ULID   `pk:"autoincr"`

	Exchange  *Exchange
	Asset    string
	Volume   float32

	UserId *User `fk:"user_id"`
}
