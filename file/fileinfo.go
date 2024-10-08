package file

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"reflect"
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
	// if p.CheckFileDeletions() == true {
	// 	return true
	// }

	return false
}

// TODO: make it so it will know if file/dir was deleted/moved etc
func (p *ProjectFilesInfo) CheckFileDeletions() bool {

	files := p.FileName

	existingFiles := GetFilesInProject()
	return reflect.DeepEqual(files, existingFiles)
}

func GetFilesInProject() []string {
	files, err := os.ReadDir("./")
	if err != nil {
		panic("cannot read file dir")
	}

	var fileSlice []string
	for _, v := range files {
		// skip .git, .gitignore, etc.
		if strings.HasPrefix(v.Name(), ".") {
			continue
		}

		if v.IsDir() {
			fileSlice = append(fileSlice, v.Name())
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

	var out bytes.Buffer
	for {
		fileChanged := fileProjectInfo.CheckIfChanged()
		if fileChanged {
			log.Println("FOUND CHANGE")
			splitCmd, args := splitUserCommand(command)

			cmd := exec.Command(splitCmd, args...)
			cmd.Dir = "./"
			cmd.Stdout = &out

			err := cmd.Run()

			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(out.String())

			fileChanged = false
		}
		goto watcherUpdate
	}
}

func splitUserCommand(command string) (string, []string) {

	cmd := stringToSlice(command)

	cmdName := cmd[0]

	cmdArgs := cmd[1:]

	return cmdName, cmdArgs

}

func stringToSlice(str string) []string {

	return strings.Fields(str)
}
