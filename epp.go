package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/blendle/epp/epp"
)

var (
	// Version of the application
	Version string

	// GitCommit of the application
	GitCommit string

	output      = flag.String("o", "", "output file")
	version     = flag.Bool("version", false, "print epp version")
	partialsDir = flag.String("partials-dir", epp.DefaultPartialsPath, "pass the path to your partials directory")
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

	out, err := epp.Parse(fileContents, *partialsDir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "templating error: %s\n", err)
		os.Exit(1)
	}

	if *output == "" {
		fmt.Print(string(out))
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

	inputFileNames, err := filepath.Glob(input)
	if err != nil {
		return nil, fmt.Errorf("we couldn't glob the provided path: %v", err)
	}

	allContent := bytes.Buffer{}

	for _, fileName := range inputFileNames {
		content, err := ioutil.ReadFile(fileName)
		if err != nil {
			return nil, fmt.Errorf("unable to read input file: %s - %v", fileName, err)
		}

		allContent.Write(content)
	}

	return allContent.Bytes(), nil
}
