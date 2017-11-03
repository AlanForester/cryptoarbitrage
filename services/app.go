package services

import (
	"log"
	. "CryptoArbitrage/helpers/arg-parser"
	. "CryptoArbitrage/helpers"
)

var Application *ApplicationModel

type ApplicationModel struct {}

func (a *ApplicationModel) Loader() {
	log.Println(ArgumentParser.Daemon.IsUsed())
	if ArgumentParser.Daemon.IsUsed() {
		log.Println("Program running in daemon mode...")
		Daemon.Start(a.start)
	} else {
		log.Println("Program running in foreground mode...")
		a.start()
	}
}

func (a *ApplicationModel) start() {
	log.Println("Application started!")
	//exchanges.TestVar = "546"
	//log.Println(exchanges.TickerModel.GetAssets())
}

func init() {
	Application = new(ApplicationModel)
}