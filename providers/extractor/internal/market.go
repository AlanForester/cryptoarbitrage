package internal

type MarketSet struct {
	Markets []*Market `json:"result"`
	Allowance      `json:"allowance"`
}

type Market struct {
	Exchange string
	Pair string
	Active bool
}