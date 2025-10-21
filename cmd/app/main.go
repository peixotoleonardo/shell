package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/peixotoleonardo/shell/internal/register"

	_ "github.com/peixotoleonardo/shell/internal/command"
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
			panic(fmt.Sprintf("there was an error during input reading: %v\n", err))
		}

		cmd, args := parse(raw)

		register.Execute(cmd, args)
	}
}
