//go:generate kallax gen
//go:generate proteus:generate

package models

import (
	"gopkg.in/src-d/go-kallax.v1"
)

type Order struct {
	kallax.Model              `table:"orders"`
	ID            kallax.ULID `pk:"autoincr"`
	kallax.Timestamps
	UserId        *User       `fk:"user_id,inverse"`
	ExchangeId    *Exchange   `fk:"exchange_id,inverse"`
	PairId        *Pair       `fk:"pair_id,inverse"`
	MarketId      *Market     `fk:"market_id,inverse"`
	OrderType     string      `kallax:"order_type"`
	OpenPrice     float32     `kallax:"open_price"`
	ClosePrice    float32     `kallax:"close_price"`
	OrderedVolume float32     `kallax:"ordered_volume"`
	SwappedVolume float32     `kallax:"swapped_volume"`
	IsClosed      bool        `kallax:"is_closed"`
	StopLoss      float32     `kallax:"stop_loss"`
	TakeProfit    float32     `kallax:"take_profit"`
	BuyFee        float32     `kallax:"buy_fee"`
	SellFee       float32     `kallax:"sell_fee"`
	Delta         float32     `kallax:"delta"`
}
