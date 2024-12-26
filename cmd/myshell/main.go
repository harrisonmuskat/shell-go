package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Regular expression to match "exit X" where X is a number
	exitPattern := regexp.MustCompile(`^exit (\d+)$`)

	for {
		fmt.Fprint(os.Stdout, "$ ")

		command, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
			os.Exit(1)
		}

		// Trim the newline character from the command
		command = strings.TrimSpace(command)

		if matches := exitPattern.FindStringSubmatch(command); matches != nil {
			// Convert the exit code to an integer
			exitCode, err := strconv.Atoi(matches[1])
			if err != nil {
				fmt.Fprintln(os.Stderr, "Invalid exit code:", matches[1])
				continue
			}

			// Exit with the specified code
			os.Exit(exitCode)
		}

		// Handle all other commands
		fmt.Println(command + ": command not found")
	}
}
