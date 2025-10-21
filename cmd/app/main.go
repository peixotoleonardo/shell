package main

import (
	"bufio"
	"fmt"
	"os"
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

		cmd, _ := parse(raw)

		switch cmd {
		case "exit":
			os.Exit(0)
		}

		fmt.Printf("%s: command not found\n", cmd)
	}
}
