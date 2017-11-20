package internal

type AssetSet struct {
	Assets []*Asset 		`json:"result"`
	Allowance			`json:"allowance"`
}

type Asset struct {
	Symbol string 		`json:"symbol"`
	Name string			`json:"name"`
	Fiat bool			`json:"fiat"`
}