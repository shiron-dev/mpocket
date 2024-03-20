package exec

import (
	"fmt"
	"os"

	"golang.org/x/sys/execabs"
)

func CheckAllCommands() {
	ok := true
	for i := 0; i < CommandLast; i++ {
		if !checkCommand(i) {
			ok = false
			fmt.Fprintf(os.Stderr, "Command %s is not found\n", GetCommandName(i))
		}
	}

	if !ok {
		os.Exit(1)
	}
}

func checkCommand(command int) bool {
	cmd := execabs.Command(GetCommandName(command), "--version")
	if err := cmd.Run(); err != nil {
		return false
	}
	return true
}
