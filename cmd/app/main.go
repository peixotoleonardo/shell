package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parse(raw string) (string, []string) {
	parts := strings.Split(strings.TrimSpace(raw), " ")

	return parts[0], parts[1:]
}

func main() {
	for {
		fmt.Print("$ ")

		raw, err := bufio.NewReader(os.Stdin).ReadString('\n')

		if err != nil {
			fmt.Fprintf(os.Stderr, "there was an error: %v\n", err)

			os.Exit(1)
		}

		cmd, args := parse(raw)

		switch cmd {
		case "exit":
			handleExitCommand(args)
		case "echo":
			handleEchoCommand(args)
		default:
			fmt.Fprintf(os.Stdout, "%s: command not found\n", cmd)
		}
	}
}

func handleExitCommand(args []string) {
	if len(args) > 0 {
		code, err := strconv.Atoi(args[0])

		if err != nil {
			os.Exit(0)
		}

		os.Exit(code)
	}

	os.Exit(0)
}

func handleEchoCommand(args []string) {
	fmt.Fprintf(os.Stdout, "%s\n", strings.Join(args, " "))
}
