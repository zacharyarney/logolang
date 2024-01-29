package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 2 {
		fmt.Println("Usage: lox [script]")
		os.Exit(64)
	} else if len(os.Args) == 2 {
		runFile(os.Args[1])
	} else {
		runPrompt()
	}
}

// runFile runs the lox interpreter on a file when a file path is provided as an argument
func runFile(path string) {
	byes, err := os.ReadFile(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't read from file: %v \n %v\n", path, err)
		os.Exit(74)
	}
	run(string(byes))
}

// runPrompt runs the lox interpreter in a REPL when no arguments are provided
func runPrompt() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		run(line)
	}
	fmt.Println("Exiting program.")
	os.Exit(0)
}

// run runs the lox interpreter on the source code string
func run(source string) {
	fmt.Println(source)
}
