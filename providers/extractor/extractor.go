package extractor

import (
	"log"
	. "crypto-arbitrage/helpers"
	. "crypto-arbitrage/providers/extractor/internal"
	"encoding/json"
	. "strings"
)

var Extractor *extractorModel

const (
	getAssetsURL = "https://api.cryptowat.ch/assets"
	getPairsURL  = "https://api.cryptowat.ch/pairs"
	getPricesURL = "https://api.cryptowat.ch/markets/prices"
	getMarketsURL = "https://api.cryptowat.ch/markets"
)

type extractorModel struct{}

func (t *extractorModel) GetAssets() []*Asset {
	response := HTTPClient.Get(getAssetsURL)
	if response.StatusCode != 200 {
		log.Fatalln(response.Error)
		return nil
	}

	var data AssetSet
	err := json.Unmarshal(response.Body, &data)
	if err != nil {
		log.Fatalln(err)
		return nil
	}
	return data.Assets
}

func (t *extractorModel) GetPairs() []*Pair {
	response := HTTPClient.Get(getPairsURL)
	if response.StatusCode != 200 {
		log.Fatalln(response.Error)
		return nil
	}

	var data PairSet
	err := json.Unmarshal(response.Body, &data)
	if err != nil {
		log.Fatalln(err)
		return nil
	}
	return data.Pairs
}

func (t *extractorModel) GetMarkets() []*Market {
	response := HTTPClient.Get(getMarketsURL)
	if response.StatusCode != 200 {
		log.Fatalln(response.Error)
		return nil
	}

	var data MarketSet
	err := json.Unmarshal(response.Body, &data)
	if err != nil {
		log.Fatalln(err)
		return nil
	}
	return data.Markets
}

func (t *extractorModel) GetPrices() Prices {
	response := HTTPClient.Get(getPricesURL)
	if response.StatusCode != 200 {
		log.Fatalln(response.Error)
		return nil
	}

	var data PricesSet
	err := json.Unmarshal(response.Body, &data)
	if err != nil {
		log.Fatalln(err)
		return nil
	}

	var prices Prices
	for key, price := range data.Prices {
		s := Split(key,":")
		exchange, symbol := s[0], s[1]
		priceModel := &Price{Exchange: exchange, Symbol: symbol, Price: float32(price)}
		prices = append(prices, priceModel)
	}
	return prices
}

func init() {
	Extractor = &extractorModel{}
}
