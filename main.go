package main

import (
	"flag"
	"fmt"
	"os"
	"github.com/boris/pretend/cmd/pretend"
)

func main() {
	stepsFile := flag.String("steps", "steps.json", "Path to the steps file. In green you'll see commands to copy/paste.")
	flag.Parse()

	script, err := pretend.ScriptFromFile(*stepsFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading script: %v\n", err)
		os.Exit(1)
	}

	script.Run()
	
}
