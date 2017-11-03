package cli_args

import (
	"flag"
)

var DaemonArgument DaemonArgumentModel

var signal = flag.String("d", "", `Run application in daemon mode.
		start — reloading the configuration file
		stop — fast shutdown`)

type DaemonArgumentModel struct {
	Flag  *string
	Value string
}

func (dc *DaemonArgumentModel) IsUsed() bool {
	return dc.Value == "daemon" || dc.Value == "start" || dc.Value == "stop"
}

func (dc *DaemonArgumentModel) CheckValue(sig string) bool {
	return sig == dc.Value
}

func init() {
	if DaemonArgument == (DaemonArgumentModel{}) {
		flag.Parse()
		da := new(DaemonArgumentModel)
		da.Flag = signal
		da.Value = *signal
		DaemonArgument = *da
	}
}
