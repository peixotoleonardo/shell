package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	_ "github.com/peixotoleonardo/shell/internal/builtin"
	"github.com/peixotoleonardo/shell/internal/register"
)

func main() {
	for {
		fmt.Print("$ ")

		raw, err := bufio.NewReader(os.Stdin).ReadString('\n')

		if err != nil {
			panic(fmt.Sprintf("there was an error during input reading: %v\n", err))
		}

		parts := strings.Split(strings.TrimSpace(raw), " ")

		args := []string{}

		if len(parts) > 1 {
			args = parts[1:]
		}

		cmd, err := register.NewCommand(parts[0], args)

		if err != nil {
			fmt.Fprintf(os.Stdout, "%s: command not found\n", parts[0])

			continue
		}

		cmd.Execute()
	}
}
