package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Fprint(os.Stdout, "$ ")

		command, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
			os.Exit(1)
		}

		// Trim the newline character from the command
		command = strings.TrimSpace(command)

		// extract the first argument from the command
		parts := strings.Fields(command)
		if len(parts) == 0 {
			continue
		}
		command = parts[0]

		switch command {
		case "exit":
			if len(parts) > 1 {
				exitCode, err := strconv.Atoi(parts[1])
				if err != nil {
					fmt.Fprintln(os.Stderr, "Invalid exit code:", parts[1])
					continue
				}

				// Exit with the specified code
				os.Exit(exitCode)
			} else {
				// Exit with code 0
				os.Exit(0)
			}
		case "echo":
			fmt.Println(strings.Join(parts[1:], " "))
		case "type":
			switch parts[1] {
			case "echo", "exit", "type":
				fmt.Println(parts[1] + " is a shell builtin")
			default:
				path, exists := os.LookupEnv("PATH")
				if !exists {
					fmt.Fprintln(os.Stderr, "PATH environment variable not set")
				} else {
					found := false
					paths := strings.Split(path, ":")
					for _, p := range paths {
						_, err := os.Stat(p + "/" + parts[1])
						if err == nil {
							fmt.Println(parts[1] + " is " + p + "/" + parts[1])
							found = true
							break
						}
					}
					if !found {
						fmt.Println(parts[1] + ": not found")
					}
				}
			}
		default:
			fmt.Println(command + ": command not found")
		}

	}
}
