package execcommand

import (
	"errors"
	"fmt"
	"strings"
)

func commandToExecute(args []string) (bool, []string, error) {
	var command []string
	for i := 0; i < len(args); i++ {
		if args[i] == "-e" || args[i] == "--execute" {
			if i+1 < len(args) {
				command = args[i+1:]
			}
			return true, command, nil
		}

	}
	err := errors.New("no -e flag found")
	return false, command, err
}

func Run(arguments []string) string {

	var cmdStr string

	ok, cmd, err := commandToExecute(arguments)
	if !ok {
		return ""
	}
	if err != nil {
		panic(err)
	}

	for i := 0; i < len(cmd); i++ {

		cmdStr += fmt.Sprintf("%v ", cmd[i])
	}
	cmdStr = strings.TrimSpace(cmdStr)
	return cmdStr

}
