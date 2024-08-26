package main

import (
	"fmt"
	"os"

	execCommand "github.com/mvstermind/file-watcher/execCommand"
	"github.com/mvstermind/file-watcher/filedata"
)

func main() {

	cmd := execCommand.Run(os.Args)
	fmt.Println(cmd)
	filedata.Watch(cmd)

}
