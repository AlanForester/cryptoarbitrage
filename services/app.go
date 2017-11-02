package services
import (
	"github.com/sevlyar/go-daemon"
	"time"
	"log"
	"flag"
	"syscall"
	"os"
)

var (
	signal = flag.String("s", "", `send signal to the daemon
		quit — graceful shutdown
		stop — fast shutdown
		reload — reloading the configuration file`)
	daemonFlag = flag.String("r", "", `Run app as daemon`)
)

var (
	stop = make(chan struct{})
	done = make(chan struct{})
)

type Application struct {}

func (a *Application) Run() {
	flag.Parse()
	if *daemonFlag == "daemon" {
		print("Daemon mode\n")
		daemon.AddCommand(daemon.StringFlag(signal, "stop"), syscall.SIGTERM, termHandler)

		cntxt := &daemon.Context{
			PidFileName: "pids/application.pid",
			PidFilePerm: 0644,
			LogFileName: "logs/application.log",
			LogFilePerm: 0640,
			WorkDir:     "./",
			Umask:       027,
			Args:        []string{"[crypto-arbitrage]"},
		}

		if len(daemon.ActiveFlags()) > 0 {
			d, err := cntxt.Search()
			log.Printf("%+v %+v",d, err)
			if err != nil {
				log.Fatalln("Unable send signal to the daemon:", err)
			}
			daemon.SendCommands(d)
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

		log.Println("- - - - - - - - - - - - - - -")
		log.Println("daemon started")

		go RunServices()

		err = daemon.ServeSignals()
		if err != nil {
			log.Println("Error:", err)
		}
		log.Println("daemon terminated")
	} else {
		print("Foreground mode\n")
		RunServices()
	}

}

func RunServices() {
	worker()
}

func worker() {
	for {
		time.Sleep(time.Second)
		log.Println("work...")
		if _, ok := <-stop; ok {
			break
		}
	}
	done <- struct{}{}
}

func termHandler(sig os.Signal) error {
	log.Println("terminating...")
	stop <- struct{}{}
	if sig == syscall.SIGQUIT {
		<-done
	}
	return daemon.ErrStop
}
