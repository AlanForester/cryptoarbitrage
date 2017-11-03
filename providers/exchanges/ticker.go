package exchanges

import (
	. "CryptoArbitrage/helpers"

	"log"
)

var Ticker = NewTicker()

const (
	GetAssetsURL = "https://api.cryptowat.ch/assets"
)

type TickerModel struct {}

type assetsModel struct {

}

func (t *TickerModel) GetAssets() string {
	assetsRes := HTTPClient.Get(GetAssetsURL)
	log.Println(assetsRes)
	//response, _ := ioutil.ReadAll(assetsRes.Response)
	//log.Println(response)
	//var data map[string]interface{}
	//err := json.Unmarshal(response, &data)
	//if err != nil {
	//	panic(err)
	//}
	//log.Println(data)
	//json.NewDecoder(assetsRes.Body).Decode(assetsModel)

	return ""
}

func NewTicker() *TickerModel {
	return new(TickerModel)
}