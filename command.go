package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

func processLine(line string, initialDir string) {
	if strings.HasPrefix(line, "cd ") {
		handleCdCommand(line, initialDir)
		return
	}
	executeCommand(line)
}

func handleCdCommand(line string, initialDir string) {
	newDir := strings.TrimSpace(strings.TrimPrefix(line, "cd "))
	if newDir == "" {
		newDir = initialDir
	}
	if err := os.Chdir(newDir); err != nil {
		log.Printf("cd: %s\n", err)
	}
}

func executeCommand(line string) {
	pattern := regexp.MustCompile(`^(\w+)="?(.*?)"?\s+(.*)`)
	matches := pattern.FindStringSubmatch(line)
	var cmd *exec.Cmd

	if len(matches) > 0 {
		varName := matches[1]
		varValue := strings.Trim(matches[2], `"`)
		command := matches[3]
		cmd = exec.Command("bash", "-c", command)
		cmd.Env = append(os.Environ(), fmt.Sprintf("%s=%s", varName, varValue))
	} else {
		cmd = exec.Command("bash", "-c", line)
	}

	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	if err := cmd.Run(); err != nil {
		log.Printf("failed to execute process: %s\n", err)
	}
}
