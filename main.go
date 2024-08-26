package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	for {
		filez := readFilesInDir()
		for i := 0; i < len(filez); i++ {
			currFile, err := os.Stat(filez[i])

			if err != nil {
				fmt.Println("cannot get file info")
				return
			}
			// we get time of modification for all fies in curr directory
			fmt.Println(currFile.Name(), currFile.ModTime())

		}
	}

}

func readFilesInDir() []string {
	files, err := os.ReadDir("./")
	if err != nil {
		panic("cannot read file dir")
	}

	var fileSlice []string
	for _, v := range files {

		// avoid .git .gitignore etc
		if strings.HasPrefix(v.Name(), ".") {
			continue
		}
		fileSlice = append(fileSlice, v.Name())
	}
	return fileSlice
}
