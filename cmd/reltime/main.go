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
var usage = `usage: reltime [options] <time> [compare time]

A simple tool to parse and display time
`

var versionFlag = flag.Bool("version", false, "Display Version.")
var shortVersionFlag = flag.Bool("v", false, "Display Short Version.")

var displayAgeFlag = flag.String("age", "s", "Display age format [s, h, d, string]. (default:s)")
var displayAgeAbsFlag = flag.Bool("absolute", false, "Display age as absolute number.")

//nolint:forbidigo // CLI tool output.
func main() {
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
			errorAndExitf("unable to parse time: %s\n", arg)
		}
		tslist = append(tslist, ts)
	}

	if len(tslist) < 1 {
		usageAndExitf("must provide at least one datetime")
	}

	td := time.Since(tslist[0])
	if len(tslist) >= 2 { //nolint:gomnd // two times.
		td = tslist[1].Sub(tslist[0])
		if *displayAgeAbsFlag {
			if td < 0 {
				td = -1 * td
			}
		}
	}
	switch {
	case displayAgeFlag != nil:
		switch strings.ToLower(*displayAgeFlag) {
		case "string", "str":
			fmt.Printf("%s", td.String())
		case "d":
			fmt.Printf("%.0f", td.Hours()/24) //nolint:gomnd // 24 hours in day
		case "h":
			fmt.Printf("%.0f", td.Hours())
		case "s":
			fmt.Printf("%.0f", td.Seconds())
		default:
			usageAndExitf("invalid age format specified: %s", *displayAgeFlag)
		}
	default:
		fmt.Printf("%.0f", td.Seconds())
	}
}

func errorAndExitf(format string, args ...interface{}) {
	if format != "" {
		if !strings.HasSuffix(format, "\n") {
			format += "\n"
		}
		format = "error: " + format
		fmt.Fprintf(os.Stderr, format, args...)
	}

	os.Exit(1)
}

func usageAndExitf(format string, args ...interface{}) {
	if format != "" {
		if !strings.HasSuffix(format, "\n") {
			format += "\n"
		}
		format = "error: " + format
		fmt.Fprintf(os.Stderr, format, args...)
	}

	flag.Usage()
	os.Exit(0)
}

func versionAndExit(short bool) {
	if short {
		fmt.Fprintf(os.Stdout, "%s", version)
	} else {
		fmt.Fprintf(os.Stdout, "%s [%s] (%s) <%s>", version, commit, date, builtBy)
	}

	os.Exit(0)
}
