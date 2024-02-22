package main

import (
	"fmt"
	"log"
	"os"
)

func getInitialDir() (string, error) {
	return os.Getwd()
}

func displayPrompt() {
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\033[32m%s\033[0m $ ", currentDir)
}
