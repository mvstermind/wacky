package execcommand

import "fmt"

func commandToExecute(args []string) (bool, []string) {
	var command []string
	for i := 0; i < len(args); i++ {
		if args[i] == "-e" || args[i] == "--execute" {
			if i+1 < len(args) {
				command = args[i+1:]
			}
			return true, command
		}
	}
	return false, nil
}

func Run(arguments []string) string {

	var cmdStr string

	ok, cmd := commandToExecute(arguments)
	if !ok {
		fmt.Println("Cannot find argument to execute")
		return ""
	}

	for i := 0; i < len(cmd); i++ {

		cmdStr += fmt.Sprintf("%v ", cmd[i])
	}
	return cmdStr

}

