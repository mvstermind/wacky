package main

import (
	"fmt"
	"os"

	execCommand "github.com/mvstermind/file-watcher/execCommand"
	file "github.com/mvstermind/file-watcher/file"
)

func main() {

	cmd := execCommand.Run(os.Args)
	fmt.Println(cmd)
	file.Watch(cmd)

}
