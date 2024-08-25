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
		newTime := fileInfo.ModTime()

		if newTime.After(oldTime) {
			fmt.Println("This file has changed")
			oldTime = newTime
		}
		time.Sleep(200 * time.Millisecond)

	}
}

