package arg_parser

import (
	"flag"
	. "CryptoArbitrage/helpers/arg-parser/cli-args"
)

type ArgumentParser struct {
	Daemon DaemonArgument
}

func NewArgumentParser() *ArgumentParser {
	flag.Parse()
	ap := new(ArgumentParser)
	ap.Daemon = *NewDaemonArgument()
	return ap
}



