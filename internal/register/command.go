package register

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type Command struct {
	Name      string
	Path      string
	IsBuiltin bool
	Args      []string
}

func NewCommand(name string, args []string) (*Command, error) {
	_, isBuiltin := commandRegistry[name]

	path := findExecutablePath(name)

	if !isBuiltin && path == "" {
		return nil, fmt.Errorf("%s: command not found", name)
	}

	cmd := &Command{
		Name:      name,
		IsBuiltin: isBuiltin,
		Args:      args,
		Path:      findExecutablePath(name),
	}

	return cmd, nil
}

func (c *Command) String() string {
	return c.Name
}

func findExecutablePath(cmd string) string {
	for path := range strings.SplitSeq(os.Getenv("PATH"), ":") {
		cmdPath := filepath.Join(path, cmd)

		info, err := os.Stat(cmdPath)

		if err != nil {
			continue
		}

		if info.Mode()&0111 != 0 {
			return cmdPath
		}
	}

	return ""
}

func (c *Command) Execute() error {
	if c.IsBuiltin {
		return commandRegistry[c.Name](c.Args)
	}

	cmd := exec.Command(c.Path, c.Args...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
