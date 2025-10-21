package main

import (
	"fmt"
)

func main() {
	fmt.Print("$ ")

	var cmd string

	for {
		fmt.Scan(&cmd)

		fmt.Printf("%s: command not found\n", cmd)
	}
}
