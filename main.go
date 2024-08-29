package main

import (
	"os"

	execCommand "github.com/mvstermind/file-watcher/execCommand"
	file "github.com/mvstermind/file-watcher/file"
)

func main() {

	cmd := execCommand.Run(os.Args)
	file.Watch(cmd)

}
