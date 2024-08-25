package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {

	fileInfo, err := os.Stat("test.txt")
	if err != nil {
		log.Println(err)
		return
	}

	oldTime := fileInfo.ModTime()

	for {

		fileInfo, err = os.Stat("test.txt")
		if err != nil {
			log.Println(err)
			return
		}

		fileState, err := os.ReadFile("test.txt")
		if err != nil {
			log.Println(err)
		}

		newTime := fileInfo.ModTime()

		if newTime.After(oldTime) {
			fmt.Println("This file has changed")
			oldTime = newTime

			newFileState, err := os.ReadFile("test.txt")
			if err != nil {
				fmt.Println("bruh this can't even happen")
			}

			for _, v := range newFileState {

				if fileState[v] != newFileState[v] {
					fmt.Printf("%v", v)
				}

			}
		}
		time.Sleep(200 * time.Millisecond)

	}
}
