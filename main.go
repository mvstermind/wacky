package main

import (
	"os"

	execCommand "github.com/mvstermind/wacky/execCommand"
	file "github.com/mvstermind/wacky/file"
)

func main() {

	cmd := execCommand.Run(os.Args)
	file.Watch(cmd)

}
