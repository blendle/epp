package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/blendle/epp/epp"
)

var (
	// Version of the application
	Version string

	// GitCommit of the application
	GitCommit string

	output  = flag.String("o", "", "output file")
	version = flag.Bool("version", false, "print epp version")
)

func main() {
	flag.Parse()

	if *version {
		fmt.Fprintf(os.Stderr, "epp %s (%s)\n", Version, GitCommit)
		os.Exit(0)
	}

	if len(flag.Args()) == 0 {
		fmt.Fprintln(os.Stderr, "error: an input file is required")
		os.Exit(1)
	}

	fileContents, err := readInput(flag.Arg(0))
	if err != nil {
		fmt.Fprintf(os.Stderr, "IO error: %s\n", err)
		os.Exit(1)
	}

	out, err := epp.Parse(fileContents)
	if err != nil {
		fmt.Fprintf(os.Stderr, "templating error: %s\n", err)
		os.Exit(1)
	}

	if *output == "" {
		fmt.Printf(string(out))
		return
	}

	err = ioutil.WriteFile(*output, out, 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "IO error: %s\n", err)
		os.Exit(1)
	}
}

func readInput(input string) ([]byte, error) {
	if inputFile := flag.Arg(0); inputFile == "-" {
		return ioutil.ReadAll(os.Stdin)
	}

	return ioutil.ReadFile(input)
}
