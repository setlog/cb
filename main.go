package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/atotto/clipboard"
)

func main() {
	flagTruncNewLine := flag.Bool("n", false, "Set to truncate one trailing line break if it exists.")
	flag.Parse()

	var cb string
	var err error
	if flag.NArg() == 0 {
		data, ioErr := ioutil.ReadAll(os.Stdin)
		cb, err = string(data), ioErr
	} else if flag.NArg() == 1 {
		filePath := flag.Arg(0)
		cb, err = readTextFile(filePath)
	} else {
		fatalf("provide file path or input on stdin")
	}
	if err != nil {
		fatalf("%v", err)
	}

	if *flagTruncNewLine && len(cb) > 0 && cb[len(cb)-1] == '\n' {
		cb = cb[:len(cb)-1]
	}

	err = clipboard.WriteAll(cb)
	if err != nil {
		fatalf("%v", err)
	}
}

func readTextFile(filePath string) (string, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(data), err
}

func fatalf(formatMessage string, args ...interface{}) {
	fmt.Printf(formatMessage+"\n", args...)
	os.Exit(1)
}
