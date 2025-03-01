package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/Forever-Young/otto"
	"github.com/Forever-Young/otto/parser"
	// "github.com/Forever-Young/otto/underscore"
)

// var flag_underscore *bool = flag.Bool("underscore", true, "Load underscore into the runtime environment")

func readSource(filename string) ([]byte, error) {
	if filename == "" || filename == "-" {
		return ioutil.ReadAll(os.Stdin)
	}
	return ioutil.ReadFile(filename)
}

func main() {
	flag.Parse()

	// if !*flag_underscore {
	// 	underscore.Disable()
	// }

	err := func() error {
		src, err := readSource(flag.Arg(0))
		if err != nil {
			return err
		}

		// vm := otto.New()
		// _, err = vm.Run(src)
		_, err = parser.ParseFile(nil, flag.Arg(0), src, 0)
		return err
	}()
	if err != nil {
		switch err := err.(type) {
		case *otto.Error:
			fmt.Print(err.String())
		default:
			fmt.Println(err)
		}
		os.Exit(64)
	}
}
