//go:generate kallax gen
//go:generate proteus:generate

package models

import (
	"gopkg.in/src-d/go-kallax.v1"
)

type ExchangeAsset struct {
	kallax.Model               `table:"exchange_assets"`
	ID             kallax.ULID `pk:"autoincr"`
	AssetId        *Asset      `fk:"asset_id,inverse"`
	ExchangeId     *Exchange   `fk:"exchange_id,inverse"`
	TransactionFee float32     `kallax:"transaction_fee"`
}
