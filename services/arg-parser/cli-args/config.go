package cli_args

import (
	"fmt"
	"os"
	"flag"
)

var ConfigArgument ConfigArgumentModel

type ConfigArgumentModel struct {
	Name  string
	Value string
	Flag   *string
}

func (dc *ConfigArgumentModel) CheckValue(sig string) bool {
	return sig == dc.Value
}

func (dc *ConfigArgumentModel) Required() {
	fmt.Fprintf(os.Stderr, "Missing required -%s flag!\n", dc.Name)
	os.Exit(0) // the same exit code flag.Parse uses
}

func init() {
	if ConfigArgument == (ConfigArgumentModel{}) {
		da := new(ConfigArgumentModel)
		da.Name = "e"
		da.Flag = flag.String(da.Name, "", `Set environment`)
		flag.Parse()
		da.Value = *da.Flag
		if da.CheckValue("") {
			flag.PrintDefaults()
			da.Required()
		}
		ConfigArgument = *da
	}
}
