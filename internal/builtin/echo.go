package command

import (
	"fmt"
	"os"
	"strings"

	"github.com/peixotoleonardo/shell/internal/register"
)

func init() {
	register.RegisterCommandHandle("echo", HandleEcho)
}

func HandleEcho(args []string) error {
	fmt.Fprintf(os.Stdout, "%s\n", strings.Join(args, " "))

	return nil
}
