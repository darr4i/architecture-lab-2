package main

import (
	"flag"
	"fmt"
	"io"
	//"log"
	"os"
	"strings"

	lab2 "github.com/darr4i/architecture-lab-2"
)

var (
	inputFromStdin = flag.String("e", "", "Expression from stdin")
	inputFromFile = flag.String("f", "", "Expression from file")
	outputToFile = flag.String("o", "", "Result to file")
)

func main() {
	flag.Parse()

	var input io.Reader = nil
	var output = os.Stdout

	if *inputFromStdin != "" {
		input = strings.NewReader(*inputFromStdin)
	}

	if *inputFromFile != "" {
		file, err := os.Open(*inputFromFile)
		if err != nil {
			os.Stderr.WriteString("Error with input from file")
			return
		}
		defer file.Close()
		input = file
	}

	if *outputToFile != "" {
		file, err := os.Create(*outputToFile)
		if err != nil {
			os.Stderr.WriteString("Error with output to file")
			return
		}
		defer file.Close()
		output = file
	}

	if input == nil {
		os.Stderr.WriteString("Error with input, got <nil>")
		return
	}

	handler := &lab2.ComputeHandler{
		Input: input,
		Output: output,
	}

	err := handler.Compute()
	
	if err != nil {
		fmt.Println(err) 
		return
	}
}
