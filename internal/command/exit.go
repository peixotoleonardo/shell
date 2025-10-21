package command

import (
	"os"
	"strconv"

	"github.com/peixotoleonardo/shell/internal/register"
)

func init() {
	register.RegisterCommand("exit", HandleExit)
}

func HandleExit(args []string) {
	if len(args) > 0 {
		code, err := strconv.Atoi(args[0])

		if err != nil {
			os.Exit(0)
		}

		os.Exit(code)
	}

	os.Exit(0)
}
