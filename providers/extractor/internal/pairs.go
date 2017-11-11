package internal

type PairSet struct {
	Pairs []Pair `json:"result"`
	Allowance      `json:"allowance"`
}

type Pair struct {
	Symbol string    `json:"symbol"`
	Base   PairBase  `json:"base"`
	Quote  PairQuote `json:"quote"`
}

type PairBase struct {
	Symbol string `json:"symbol"`
}

type PairQuote struct {
	Symbol string `json:"symbol"`
}
