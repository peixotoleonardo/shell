package command

import (
	"fmt"
	"os"

	"github.com/peixotoleonardo/shell/internal/register"
)

func init () {
	register.RegisterCommand("type", HandleType)
}

func HandleType(args []string) {
	if len(args) == 0 {
		return
	}

	if register.IsCommandRegistered(args[0]) {
		fmt.Fprintf(os.Stdout, "%s is a shell builtin\n", args[0])
	} else {
		fmt.Fprintf(os.Stdout, "%s not found\n", args[0])
	}
}
