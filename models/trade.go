package models

import (
	"gopkg.in/src-d/go-kallax.v1"
)

type Trade struct {
	kallax.Model         `table:"trades"`
	ID       kallax.ULID `pk:"autoincr"`
	kallax.Timestamps
	User     *User       `fk:"user_id,inverse"`
	Exchange *Exchange   `fk:"exchange_id,inverse"`
	Pair     *Pair       `fk:"pair_id,inverse"`
	Market   *Market     `fk:"market_id,inverse"`
	Order    *Order      `fk:"order_id,inverse"`
	Type     string      `kallax:"type"`
	Volume   float32     `kallax:"volume"`
	Price    float32     `kallax:"price"`
}
