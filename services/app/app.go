package app

import (
	"log"
	. "CryptoArbitrage/services/arg-parser"
	. "CryptoArbitrage/services"
	. "CryptoArbitrage/providers/extractor"
)

var Application applicationModel

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
	log.Println(Extractor.GetAssets())
}

func init() {
	Application = *new(applicationModel)
}