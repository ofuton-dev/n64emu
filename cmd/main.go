package main

import (
	"flag"
	"fmt"
	"n64emu/pkg/core"
	"n64emu/pkg/core/cart"
	"os"
)

var (
	version     string
	showVersion = flag.Bool("v", false, "show version")
)

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
	flag.Parse()

	if *showVersion {
		fmt.Println(title+":", getVersion())
		return exitCodeOK
	}

	romPath := flag.Arg(0)
	eepromPath := flag.Arg(1)
	nvsramPath := flag.Arg(2)
	c, err := cart.NewCart(romPath, eepromPath, nvsramPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to read ROM data: %s\n", err)
		return exitCodeError
	}

	// test code
	fmt.Printf("ROM ImageName='%s'\n", c.ROM.ImageName)
	core.Hello()

	return exitCodeOK
}

func getVersion() string {
	if version == "" {
		return "Develop"
	}
	return version
}
