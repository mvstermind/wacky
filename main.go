package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	lastModTimes := make(map[string]time.Time)

	for {
		files := allFilesInCurrDir()
		for _, file := range files {
			fileInfo, err := os.Stat(file)
			if err != nil {
				log.Println(err)
				continue
			}

			newTime := fileInfo.ModTime()
			if lastModTime, exists := lastModTimes[file]; !exists || newTime.After(lastModTime) {
				fmt.Printf("This file has changed: %s\n", file)
				lastModTimes[file] = newTime
			}
		}

		time.Sleep(200 * time.Millisecond)
	}
}

func allFilesInCurrDir() []string {
	filesInDir, err := os.ReadDir("./")
	if err != nil {
		panic(err)
	}

	var filesSlice []string

	for _, entry := range filesInDir {
		if !entry.IsDir() { // Only consider regular files, not directories.
			filesSlice = append(filesSlice, entry.Name())
		}
	}
	return filesSlice
}

