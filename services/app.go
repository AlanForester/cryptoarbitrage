package services

import (
	"log"
	"CryptoArbitrage/helpers/arg-parser"
	"CryptoArbitrage/helpers"
)

type Application struct {}

func (a *Application) Loader() {
	args := arg_parser.NewArgumentParser()
	if args.Daemon.IsUsed() {
		log.Println("Program running in daemon mode...")
		helpers.StartDaemon(args.Daemon, a.start)
	} else {
		log.Println("Program running in foreground mode...")
		a.start()
	}
}

func (a *Application) start() {
	log.Println("Application started!")
}