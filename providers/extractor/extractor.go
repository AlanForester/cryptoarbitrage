package extractor

import (
	"log"
	. "CryptoArbitrage/helpers"
	. "CryptoArbitrage/providers/extractor/internal"
	"encoding/json"
)

var Extractor extractorModel

const (
	getAssetsURL = "https://api.cryptowat.ch/assets"
	getPairsURL = "https://api.cryptowat.ch/pairs"
)

type extractorModel struct {}

func (t *extractorModel) GetAssets() []Asset {
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

func (t *extractorModel) GetPairs() []Asset {
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


func init() {
	Extractor = * &extractorModel{}
}