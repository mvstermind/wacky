package filedata

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

type ProjectFilesInfo struct {
	FileName []string
	ModTime  []time.Time
}

func NewProjectFileInfo(files []string, modTime []time.Time) *ProjectFilesInfo {
	return &ProjectFilesInfo{
		FileName: files,
		ModTime:  modTime,
	}
}

func (p *ProjectFilesInfo) CheckIfChanged() bool {
	currentFiles, newTime := GetFileStatus(p.FileName)

	if len(currentFiles) != len(p.FileName) {
		return true
	}

	for i := 0; i < len(p.FileName); i++ {
		if !newTime[i].Equal(p.ModTime[i]) {
			return true
		}
	}

	return false
}

func GetFilesInProject() []string {
	files, err := os.ReadDir("./")
	if err != nil {
		panic("cannot read file dir")
	}

	var fileSlice []string
	for _, v := range files {
		// skip .git, .gitignore, etc.
		if v.IsDir() || strings.HasPrefix(v.Name(), ".") {
			continue
		}
		fileSlice = append(fileSlice, v.Name())
	}
	return fileSlice
}

func GetFileStatus(fileNames []string) ([]string, []time.Time) {
	var (
		fName   []string
		modTime []time.Time
	)

	for i := 0; i < len(fileNames); i++ {
		currFile, err := os.Stat(fileNames[i])

		if os.IsNotExist(err) {
			// handle missing file (indicate a change)
			continue
		} else if err != nil {
			fmt.Println("cannot get file info: ", err)
			return nil, nil
		}

		fName = append(fName, currFile.Name())
		modTime = append(modTime, currFile.ModTime())
	}
	return fName, modTime
}

func Watch(command string) {
watcherUpdate:
	projectFiles := GetFilesInProject()

	fileProjectInfo := NewProjectFileInfo(GetFileStatus(projectFiles))

	for {
		fileChanged := fileProjectInfo.CheckIfChanged()
		if fileChanged {
			fmt.Println("change found")
			cmd := exec.Command(command)
			cmd.Run()
			fileChanged = false
		}
		goto watcherUpdate
	}
}
