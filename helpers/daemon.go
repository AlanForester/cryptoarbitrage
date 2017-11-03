package helpers

import (
	goDaemon "github.com/sevlyar/go-daemon"
	. "CryptoArbitrage/helpers/arg-parser/cli-args"
	"log"
	"syscall"
	"os"
	. "CryptoArbitrage/helpers/arg-parser"
)

var Daemon DaemonModel

type DaemonModel struct {
	arg DaemonArgumentModel
}

func (dc *DaemonModel) getContext() *goDaemon.Context {
	return &goDaemon.Context{
		PidFileName: "pids/application.pid",
		PidFilePerm: 0644,
		LogFileName: "logs/application.log",
		LogFilePerm: 0640,
		WorkDir:     "./",
		Umask:       027,
		Args:        []string{"[crypto-arbitrage]", "-d", "daemon"},
	}
}

func (dc *DaemonModel) Start(f func()) {

	goDaemon.AddCommand(goDaemon.StringFlag(dc.arg.Flag, "stop"), syscall.SIGQUIT, dc.termHandler)

	cntxt := dc.getContext()

	if len(goDaemon.ActiveFlags()) > 0 {
		d, err := cntxt.Search()
		if err != nil {
			if err.Error() == "open pids/application.pid: no such file or directory" {
				log.Fatalln("Daemon not running.")
			} else {
				log.Fatalln("Unable send signal to the daemon:", err)
			}
		}
		goDaemon.SendCommands(d)
		if dc.arg.CheckValue("stop") {
			log.Println("Daemon has been stoped.")
		}
		return
	}

	d, err := cntxt.Reborn()
	if err != nil {
		log.Fatalln(err)
	}
	if d != nil {
		return
	}
	defer cntxt.Release()

	log.Println("Daemon started.")

	go f()

	err = goDaemon.ServeSignals()
	if err != nil {
		log.Println("Error:", err)
	}
	log.Println("Daemon terminated.")
}

func (dc *DaemonModel) termHandler(sig os.Signal) error {
	log.Println("Terminating...")
	return goDaemon.ErrStop
}

func init() {
	if Daemon == (DaemonModel{}) {
		dh := new(DaemonModel)
		dh.arg = ArgumentParser.Daemon
		Daemon = *dh
	}
}
