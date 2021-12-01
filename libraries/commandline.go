package commandline

import (
	"fmt"
)

type CommandLine struct {
	Command string
	// TODO: Dodaj zmienną "CommandSplit" po naprawieniu błędu z tablicami
}

func (cmd CommandLine) Input() string {
	fmt.Print("> ")
	fmt.Scanln(&cmd.Command)
	return cmd.Command
}
