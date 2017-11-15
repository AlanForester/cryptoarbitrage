package models

import (
	"gopkg.in/src-d/go-kallax.v1"
)

type ExchangeAsset struct {
	kallax.Model               `table:"exchange_assets"`
	ID kallax.NumericID `pk:"autoincr"`
	Asset          *Asset      `fk:"asset_id,inverse"`
	Exchange       *Exchange   `fk:"exchange_id,inverse"`
	TransactionFee float32     `kallax:"transaction_fee"`
}
