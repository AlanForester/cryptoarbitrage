package exchanges

var Ticker = NewTicker()

type ticker struct {}

func (t *ticker) GetAssets() string {
	return ""
}

func NewTicker() *ticker {
	return new(ticker)
}