package builtin

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/peixotoleonardo/shell/internal/register"
)

func init() {
	register.RegisterCommandHandle("cd", HandleCd)
}

func HandleCd(args []string) error {
	dir, err := getDir(args)

	if err != nil {
		return err
	}

	if err = os.Chdir(dir); err != nil {
		return err
	}

	return nil
}

func getDir(args []string) (string, error) {
	if len(args) == 0 || args[0] == "~" {
		dir, err := os.UserHomeDir()

		if err != nil {
			return "", err
		}

		return dir, nil
	}

	if strings.HasPrefix(args[0], "~") {
		dir, err := os.UserHomeDir()

		if err != nil {
			return "", err
		}

		return filepath.Join(dir, args[0][1:]), nil
	}

	return args[0], nil
}
