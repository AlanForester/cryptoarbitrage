package helpers

import (
	"github.com/sevlyar/go-daemon"
	. "CryptoArbitrage/helpers/arg-parser/cli-args"
	"log"
	"syscall"
	"os"
)

type Daemon struct {
	arg DaemonArgument
}

func (dc *Daemon) getContext() *daemon.Context {
	return &daemon.Context{
		PidFileName: "pids/application.pid",
		PidFilePerm: 0644,
		LogFileName: "logs/application.log",
		LogFilePerm: 0640,
		WorkDir:     "./",
		Umask:       027,
		Args:        []string{"[crypto-arbitrage]", "-d", "daemon"},
	}
}

func (dc *Daemon) start(f func()) {

	daemon.AddCommand(daemon.StringFlag(dc.arg.Flag, "stop"), syscall.SIGQUIT, dc.termHandler)

	cntxt := dc.getContext()

	if len(daemon.ActiveFlags()) > 0 {
		d, err := cntxt.Search()
		if err != nil {
			if err.Error() == "open pids/application.pid: no such file or directory" {
				log.Fatalln("Daemon not running.")
			} else {
				log.Fatalln("Unable send signal to the daemon:", err)
			}
		}
		daemon.SendCommands(d)
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

	err = daemon.ServeSignals()
	if err != nil {
		log.Println("Error:", err)
	}
	log.Println("Daemon terminated.")
}

func (dc *Daemon) termHandler(sig os.Signal) error {
	log.Println("Terminating...")
	return daemon.ErrStop
}

func StartDaemon(flag DaemonArgument, f func()) {
	dh := new(Daemon)
	dh.arg = flag
	dh.start(f)
}
