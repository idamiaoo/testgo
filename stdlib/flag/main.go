package main

import (
	"flag"
	"os"

	"go.uber.org/zap"
)

var log *zap.Logger

func main() {
	if len(os.Args) < 2 {
		log.Error("need subcommand as first argument")
		return
	}
	switch os.Args[1] {
	case "install":
		doInstall(os.Args[2:])
	case "test":
		//doTest(os.Args[2:])
	default:
		log.Error("unknown command ", zap.String("arg", os.Args[1]))
	}
}

func doInstall(cmdline []string) {
	flag.CommandLine.Parse(cmdline)
	log.Info("NArg", zap.Int("NArg", flag.CommandLine.NArg()))

	//log.Println(flag.CommandLine.Args())
}
