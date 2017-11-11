package internal

type PricesMap map[string] float32

type PricesSet struct {
	Prices PricesMap `json:"result"`
	Allowance      `json:"allowance"`
}

type Prices [] Price

type Price struct {
	Exchange string
	Symbol string
	Price float32
}