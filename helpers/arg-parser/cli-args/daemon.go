package cli_args

import (
	"flag"
)

var (
	signal = flag.String("d", "", `Run application in daemon mode.
		start — reloading the configuration file
		stop — fast shutdown
	`)
)

type DaemonArgument struct {
	Flag *string
	Value string
}

func (dc *DaemonArgument) IsUsed() bool {
	return dc.Value == "daemon" || dc.Value == "start" || dc.Value == "stop"
}

func (dc *DaemonArgument) CheckValue(sig string) bool {
	return sig == dc.Value
}

func NewDaemonArgument() *DaemonArgument {
	dc := new(DaemonArgument)
	dc.Flag = signal
	dc.Value = *signal
	return dc
}