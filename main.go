package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

type ProjectFilesInfo struct {
	FileName []string
	ModTime  []time.Time
}

func main() {

fileCheck:
	projectFiles := getFilesInProject()

	fileProjectInfo := NewProjectFileInfo(getFileStatus(projectFiles))

	for {
		fileChanged := fileProjectInfo.checkIfChanged()
		if fileChanged {
			fmt.Println("restarting project")
			goto fileCheck
		}
		time.Sleep(200 * time.Millisecond)
	}

}

func NewProjectFileInfo(files []string, modTime []time.Time) *ProjectFilesInfo {

	return &ProjectFilesInfo{
		FileName: files,
		ModTime:  modTime,
	}
}

func (p *ProjectFilesInfo) checkIfChanged() bool {

	projectFiles := getFilesInProject()

	_, newTime := getFileStatus(projectFiles)
	for i := 0; i < len(p.FileName); i++ {

		if newTime[i] != p.ModTime[i] {
			return true
		}

	}

	return false

}

func getFilesInProject() []string {
	files, err := os.ReadDir("./")
	if err != nil {
		panic("cannot read file dir")
	}

	var fileSlice []string
	for _, v := range files {

		// skip .git .gitignore etc
		if v.IsDir() || strings.HasPrefix(v.Name(), ".") {
			continue
		}
		fileSlice = append(fileSlice, v.Name())
	}
	return fileSlice
}

func getFileStatus(fileNames []string) ([]string, []time.Time) {

	var (
		fName   []string
		modTime []time.Time
	)

	for i := 0; i < len(fileNames); i++ {
		currFile, err := os.Stat(fileNames[i])

		if err != nil {
			fmt.Println("cannot get file info")
			return nil, nil
		}

		fName = append(fName, currFile.Name())
		modTime = append(modTime, currFile.ModTime())
	}
	return fName, modTime

}

func watch() {

}
