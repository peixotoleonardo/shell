package builtin

import (
	"fmt"
	"os"

	"github.com/peixotoleonardo/shell/internal/register"
)

func init() {
	register.RegisterCommandHandle("pwd", HandlePwd)
}

func HandlePwd(args []string) error {
	pwd, err := os.Getwd()

	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stdout, "%s\n", pwd)

	return nil
}
