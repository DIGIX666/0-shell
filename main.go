package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	// Get the initial working directory.
	initialDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(os.Stdin)

	for {
		// Display current directory.
		currentDir, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("\033[32m%s\033[0m $ ", currentDir)

		// Read command.
		scanner.Scan()
		line := scanner.Text()
		args := strings.Fields(line)

		// Handle exit command.
		if len(args) > 0 && args[0] == "exit" {
			break
		}

		// Handle cd command.
		if len(args) > 0 && args[0] == "cd" {
			newDir := initialDir
			if len(args) > 1 {
				newDir = args[1]
			}

			err := os.Chdir(newDir)
			if err != nil {
				log.Printf("cd: %s\n", err)
			}

			continue
		}

		// Run command and wait for it to finish.
		cmd := exec.Command(args[0], args[1:]...)
		cmd.Stdin = os.Stdin
		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout

		err = cmd.Run()

		if err != nil {
			log.Printf("failed to execute process: %s\n", err)
		}
	}

	if scanner.Err() != nil {
		log.Fatal(err)
	}
}
