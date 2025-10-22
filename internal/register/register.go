package register

import (
	"fmt"
	"os"
)

type CommandHandle = func([]string)

var commandRegistry = map[string]CommandHandle{}

func RegisterCommand(cmd string, fn CommandHandle) {
	commandRegistry[cmd] = fn
}

func IsBuiltinCommand(cmd string) bool {
	_, ok := commandRegistry[cmd]
	
	return ok
}

func Execute(cmd string, args []string) {
	if handle, ok := commandRegistry[cmd]; ok {
		handle(args)
	} else {
		fmt.Fprintf(os.Stdout, "%s: command not found\n", cmd)
	}
}
