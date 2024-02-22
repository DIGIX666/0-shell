package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	initialDir, err := getInitialDir()
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(os.Stdin)
	for {
		displayPrompt()
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		if line == "exit" {
			break
		}
		processLine(line, initialDir)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
