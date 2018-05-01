package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	flags "github.com/jessevdk/go-flags"
	"github.com/mackerelio/checkers"
)

var opts struct {
	Base        string `short:"b" long:"base" required:"true" description:"the base directory(required)"`
	WarningNum  int64  `short:"w" long:"warning-num" default:"10" description:"warning if more number of files(count)"`
	CriticalNum int64  `short:"c" long:"critical-num" default:"100" description:"critical if more number of files(count)"`
	Debug       bool   `short:"d" long:"debug" description:"debug print"`
}

func main() {
	ckr := run(os.Args[1:])
	ckr.Name = "FileNum"
	ckr.Exit()
}

func run(args []string) *checkers.Checker {
	_, err := flags.ParseArgs(&opts, args)
	if err != nil {
		log.Printf("parse args: %v", err)
		os.Exit(1)
	}

	if opts.Debug {
		log.Printf("Base: %v\n", opts.Base)
		log.Printf("WarningAge: %v\n", opts.WarningNum)
		log.Printf("CriticalAge: %v\n", opts.CriticalNum)
	}

	files, err := ioutil.ReadDir(opts.Base)
	if err != nil {
		log.Printf("list file: %v", err)
		os.Exit(1)
	}

	if len(files) < 1 {
		return checkers.Ok("Directory is empty")
	}

	if opts.Debug {
		for _, f := range files {
			log.Println(f)
		}
	}

	current := int64(len(files))

	result := checkers.OK
	if current > opts.CriticalNum {
		result = checkers.CRITICAL
	}

	if current > opts.WarningNum {
		result = checkers.WARNING
	}

	msg := fmt.Sprintf("count=%d", current)
	return checkers.NewChecker(result, msg)
}
