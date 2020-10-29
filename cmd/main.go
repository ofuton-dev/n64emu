package main

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"n64emu/pkg/core"
	"n64emu/pkg/util"
	"os"
	"path/filepath"
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

	path := flag.Arg(0)
	_, err := readROM(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to read ROM data: %s\n", err)
		return exitCodeError
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

func readROM(path string) ([]byte, error) {
	if path == "" {
		return []byte{}, errors.New("please enter ROM file path")
	}

	extFilter := []string{".z64", ".n64"}
	if !util.Contains(extFilter, filepath.Ext(path)) {
		return []byte{}, fmt.Errorf("please enter file which has the following extension: %v", extFilter)
	}

	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return []byte{}, errors.New("fail to read file")
	}
	return bytes, nil
}
