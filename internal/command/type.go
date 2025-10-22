package command

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/peixotoleonardo/shell/internal/register"
)

func init() {
	register.RegisterCommand("type", HandleType)
}

func HandleType(args []string) {
	if len(args) == 0 {
		return
	}

	if register.IsBuiltinCommand(args[0]) {
		fmt.Fprintf(os.Stdout, "%s is a shell builtin\n", args[0])

		return
	}

	for path := range strings.SplitSeq(os.Getenv("PATH"), ":") {
		cmdPath := filepath.Join(path, args[0])

		info, err := os.Stat(cmdPath)

		if err != nil {
			continue
		}

		if info.Mode()&0111 != 0 {
			fmt.Fprintf(os.Stdout, "%s is %s\n", args[0], cmdPath)

			return
		}
	}

	fmt.Fprintf(os.Stdout, "%s not found\n", args[0])
}
