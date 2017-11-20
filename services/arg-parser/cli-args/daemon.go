package cli_args

import (
	"flag"
)

var DaemonArgument DaemonArgumentModel

type DaemonArgumentModel struct {
	Name  string
	Value string
	Flag  *string
}

func (dc *DaemonArgumentModel) IsUsed() bool {
	return dc.Value == "daemon" || dc.Value == "start" || dc.Value == "stop"
}

func (dc *DaemonArgumentModel) CheckValue(sig string) bool {
	return sig == dc.Value
}

func init() {
	if DaemonArgument == (DaemonArgumentModel{}) {
		da := new(DaemonArgumentModel)
		da.Name = "d"
		da.Flag = flag.String(da.Name, "", `Run application in daemon mode.
		start — reloading the configuration file
		stop — fast shutdown`)
		flag.Parse()
		da.Value = *da.Flag
		DaemonArgument = *da
	}
}
