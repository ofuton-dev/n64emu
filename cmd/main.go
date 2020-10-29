package main

import (
	"flag"
	"fmt"
	"n64emu/pkg/core"
	"os"
)

var version string

const (
	title = "n64emu"
)

const (
	exitCodeOK int = iota
	exitCodeError
)

func main() {
	os.Exit(Run())
}

func Run() int {
	var (
		showVersion = flag.Bool("v", false, "show version")
	)
	flag.Parse()
	if *showVersion {
		fmt.Println(title+":", getVersion())
		return exitCodeOK
	}

	core.Hello()
	return exitCodeOK
}

func getVersion() string {
	if version == "" {
		return "Develop"
	}
	return version
}
