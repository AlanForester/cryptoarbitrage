package arg_parser

import (
	. "CryptoArbitrage/services/arg-parser/cli-args"
)

var ArgumentParser ArgumentParserModel

type ArgumentParserModel struct {
	Daemon DaemonArgumentModel
}

func init() {
	if ArgumentParser == (ArgumentParserModel{}) {
		ap := new(ArgumentParserModel)
		ap.Daemon = DaemonArgument
		ArgumentParser = *ap
	}
}



