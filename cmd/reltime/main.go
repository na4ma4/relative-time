package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/na4ma4/relative-time/timeparser"
)

//nolint:gochecknoglobals // usage for command.
var usage = `usage: reltime [options] <time> [time1..timeN]

A simple tool to parse and display time
`

var versionFlag = flag.Bool("version", false, "Display Version")
var shortVersionFlag = flag.Bool("v", false, "Display Short Version")

// var displayAgeFlag = flag.Bool("age", false, "Display Age")

func init() {
}

//nolint:forbidigo // CLI tool output.
func main() {
	// var cmd *Command

	flag.Usage = func() {
		fmt.Fprint(flag.CommandLine.Output(), usage)
		flag.PrintDefaults()
	}

	flag.Parse()

	if *versionFlag {
		versionAndExit(false)
	}
	if *shortVersionFlag {
		versionAndExit(true)
	}

	args := flag.Args()

	tslist := []time.Time{}
	for _, arg := range args {
		ts, err := timeparser.Parse(arg)
		if err != nil {
			errorAndExit("unable to parse time: %s\n", arg)
		}
		tslist = append(tslist, ts)
	}

	if len(tslist) < 1 {
		usageAndExit("must provide at least one datetime")
	}

	switch {
	default:
		fmt.Printf("%.0f", time.Since(tslist[0]).Seconds())
	}
	// var versionFlag = flag.Bool("version", false, "Display Version")

	// flag.Usage = func() {
	// 	fmt.Fprint(os.Stderr, fmt.Sprint(usage))
	// }

	// if len(os.Args) <= 1 {
	// 	usageAndExit("")
	// }

	// switch os.Args[1] {
	// case "--version":
	// 	cmd = NewVersionCommand()
	// // case "info":
	// // 	cmd = NewInfoCommand()
	// // case "cmp":
	// // 	cmd = NewCompareCommand()
	// // case "diff":
	// // 	cmd = NewDiffCommand()
	// default:
	// 	usageAndExit(fmt.Sprintf("%s: '%s' is not a command.\n", os.Args[0], os.Args[1]))
	// }

	// if err := cmd.Init(os.Args[2:]); err != nil {
	// 	fmt.Printf("error: command init failed: %s", err)
	// 	os.Exit(1)
	// }

	// cmd.Run()
}

func errorAndExit(format string, args ...interface{}) {
	if format != "" {
		if !strings.HasSuffix(format, "\n") {
			format += "\n"
		}
		format = "error: " + format
		fmt.Fprintf(os.Stderr, format, args...)
	}

	os.Exit(1)
}

func usageAndExit(msg string) {
	if msg != "" {
		fmt.Fprint(os.Stderr, msg)
		fmt.Fprintf(os.Stderr, "\n")
	}

	flag.Usage()
	os.Exit(0)
}

func versionAndExit(short bool) {
	if short {
		fmt.Printf("%s", version)
	} else {
		fmt.Printf("%s [%s] (%s) <%s>", version, commit, date, builtBy)
	}
	os.Exit(0)
}
