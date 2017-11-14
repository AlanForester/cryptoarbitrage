//go:generate kallax gen
//go:generate proteus:generate

package models

import (
	"gopkg.in/src-d/go-kallax.v1"
	"time"
)

type User struct {
	kallax.Model              `table:"users" pk:"id,autoincr"`
	ID          kallax.ULID   `pk:"autoincr"`
	Email       string        `kallax:"email"`
	Password    string        `kallax:"password"`
	LastLogin   time.Time     `kallax:"last_login"`
	SubscribeTo time.Time     `kallax:"subscribe_to"`
	Role        string        `kallax:"role"`
	kallax.Timestamps

	Balances    *[]UserBalance `kallax:"balances"`

	Orders *[]Order `fk:"user_id"`
}