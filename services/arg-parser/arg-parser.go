package arg_parser

import (
	. "crypto-arbitrage/services/arg-parser/cli-args"
)

var ArgumentParser ArgumentParserModel

type ArgumentParserModel struct {
	Daemon DaemonArgumentModel
	Config ConfigArgumentModel
}

func init() {
	if ArgumentParser == (ArgumentParserModel{}) {
		ap := new(ArgumentParserModel)
		ap.Daemon = DaemonArgument
		ap.Config = ConfigArgument
		ArgumentParser = *ap
	}
}



