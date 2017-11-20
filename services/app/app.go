package app

import (
	"log"
	. "crypto-arbitrage/services/arg-parser"
	. "crypto-arbitrage/services"
)

var Application *applicationModel

type applicationModel struct {}

func (a *applicationModel) Loader() {
	if ArgumentParser.Daemon.IsUsed() {
		log.Println("Program running in daemon mode...")
		Daemon.Run(a.start)
	} else {
		log.Println("Program running in foreground mode...")
		defer a.start()
	}
}

func (a *applicationModel) start() {
	log.Println("Application started!")
	//cities := NewCityStore(DB.SQL)
	//cities1, _ := cities.FindAll(NewCityQuery())
	//
	//for _, city1 := range cities1 {
	//	log.Printf("%w", city1)
	//}
}

func init() {
	Application = new(applicationModel)
}