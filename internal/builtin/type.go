package builtin

import (
	"fmt"
	"os"

	"github.com/peixotoleonardo/shell/internal/register"
)

func init() {
	register.RegisterCommandHandle("type", HandleType)
}

func HandleType(args []string) error {
	if len(args) == 0 {
		return nil
	}

	cmd, err := register.NewCommand(args[0], []string{})

	if err != nil {
		fmt.Fprintf(os.Stdout, "%s not found\n", args[0])

		return nil
	}

	if cmd.IsBuiltin {
		fmt.Fprintf(os.Stdout, "%s is a shell builtin\n", cmd)

		return nil
	}

	fmt.Fprintf(os.Stdout, "%s is %s\n", cmd, cmd.Path)

	return nil
}
