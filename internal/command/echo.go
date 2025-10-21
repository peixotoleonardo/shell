package command

import (
	"fmt"
	"os"
	"strings"

	"github.com/peixotoleonardo/shell/internal/register"
)

func init() {
	register.RegisterCommand("echo", HandleEcho)
}

func HandleEcho(args []string) {
	fmt.Fprintf(os.Stdout, "%s\n", strings.Join(args, " "))
}
