package register

type CommandHandle = func([]string) error

var commandRegistry = map[string]CommandHandle{}

func RegisterCommandHandle(cmd string, fn CommandHandle) {
	commandRegistry[cmd] = fn
}
