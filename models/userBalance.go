package models

import (
	"gopkg.in/src-d/go-kallax.v1"
)

type UserBalance struct {
	kallax.Model              `table:"user_balances"`
	ID       kallax.NumericID `pk:"autoincr"`
	User     *User            `fk:"user_id,inverse"`
	Exchange *Exchange        `fk:"exchange_id,inverse"`
	Asset    *Asset           `fk:"asset_id,inverse"`
	Volume   float32          `kallax:"volume"`
}
