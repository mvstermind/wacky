package file_test

import (
	"os"
	"testing"
	"time"

	"github.com/mvstermind/file-watcher/file"
)

func TestNewProjectFileInfo(t *testing.T) {
	t.Run("Basic Initialization", func(t *testing.T) {
		files := []string{"file1.txt", "file2.txt"}
		modTimes := []time.Time{time.Now(), time.Now().Add(1 * time.Hour)}

		info := file.NewProjectFileInfo(files, modTimes)

		if len(info.FileName) != 2 {
			t.Fatalf("Expected 2 files, got %d", len(info.FileName))
		}

		if !info.ModTime[0].Equal(modTimes[0]) || !info.ModTime[1].Equal(modTimes[1]) {
			t.Fatalf("Modification times do not match")
		}
	})

	t.Run("Empty Initialization", func(t *testing.T) {
		info := file.NewProjectFileInfo(nil, nil)

		if len(info.FileName) != 0 || len(info.ModTime) != 0 {
			t.Fatalf("Expected no files, got %d", len(info.FileName))
		}
	})
}

func TestCheckIfChanged(t *testing.T) {
	t.Run("No Change Detected", func(t *testing.T) {
		tempDir := t.TempDir()
		file1 := tempDir + "/file1.txt"
		file2 := tempDir + "/file2.txt"

		os.WriteFile(file1, []byte("file1 content"), 0644)
		os.WriteFile(file2, []byte("file2 content"), 0644)

		files := []string{file1, file2}
		_, modTimes := file.GetFileStatus(files)
		projectInfo := file.NewProjectFileInfo(files, modTimes)

		if projectInfo.CheckIfChanged() {
			t.Fatalf("Expected no changes, but detected changes")
		}
	})

	t.Run("Change Detected", func(t *testing.T) {
		tempDir := t.TempDir()
		file1 := tempDir + "/file1.txt"
		file2 := tempDir + "/file2.txt"

		os.WriteFile(file1, []byte("file1 content"), 0644)
		os.WriteFile(file2, []byte("file2 content"), 0644)

		files := []string{file1, file2}
		_, modTimes := file.GetFileStatus(files)
		projectInfo := file.NewProjectFileInfo(files, modTimes)

		time.Sleep(1 * time.Second)
		os.WriteFile(file1, []byte("updated file1 content"), 0644)

		if !projectInfo.CheckIfChanged() {
			t.Fatalf("Expected changes, but detected no changes")
		}
	})

	t.Run("Identical Modification Times", func(t *testing.T) {
		tempDir := t.TempDir()
		file1 := tempDir + "/file1.txt"
		file2 := tempDir + "/file2.txt"

		os.WriteFile(file1, []byte("file1 content"), 0644)
		os.WriteFile(file2, []byte("file2 content"), 0644)

		// Ensure we pass the actual mod times
		_, modTimes := file.GetFileStatus([]string{file1, file2})
		projectInfo := file.NewProjectFileInfo([]string{file1, file2}, modTimes)

		if projectInfo.CheckIfChanged() {
			t.Fatalf("Expected no changes, but detected changes")
		}
	})

	t.Run("File Removed", func(t *testing.T) {
		tempDir := t.TempDir()
		file1 := tempDir + "/file1.txt"
		file2 := tempDir + "/file2.txt"

		os.WriteFile(file1, []byte("file1 content"), 0644)
		os.WriteFile(file2, []byte("file2 content"), 0644)

		files := []string{file1, file2}
		_, modTimes := file.GetFileStatus(files)
		projectInfo := file.NewProjectFileInfo(files, modTimes)

		os.Remove(file2)

		if !projectInfo.CheckIfChanged() {
			t.Fatalf("Expected changes due to file removal, but detected no changes")
		}
	})
}

func TestGetFilesInProject(t *testing.T) {
	t.Run("Standard Files", func(t *testing.T) {
		tempDir := t.TempDir()
		file1 := tempDir + "/file1.txt"
		file2 := tempDir + "/file2.txt"

		os.WriteFile(file1, []byte("file1 content"), 0644)
		os.WriteFile(file2, []byte("file2 content"), 0644)

		os.Chdir(tempDir)

		files := file.GetFilesInProject()

		if len(files) != 2 {
			t.Fatalf("Expected 2 files, got %d", len(files))
		}
	})

	t.Run("Hidden Files Ignored", func(t *testing.T) {
		tempDir := t.TempDir()
		file1 := tempDir + "/file1.txt"
		file2 := tempDir + "/file2.txt"
		hiddenFile := tempDir + "/.hiddenfile"

		os.WriteFile(file1, []byte("file1 content"), 0644)
		os.WriteFile(file2, []byte("file2 content"), 0644)
		os.WriteFile(hiddenFile, []byte("hidden content"), 0644)

		os.Chdir(tempDir)

		files := file.GetFilesInProject()

		if len(files) != 2 {
			t.Fatalf("Expected 2 files (ignoring hidden), got %d", len(files))
		}
	})

	t.Run("No Files", func(t *testing.T) {
		tempDir := t.TempDir()
		os.Chdir(tempDir)

		files := file.GetFilesInProject()

		if len(files) != 0 {
			t.Fatalf("Expected no files, got %d", len(files))
		}
	})
}

func TestGetFileStatus(t *testing.T) {
	t.Run("Valid Files", func(t *testing.T) {
		tempDir := t.TempDir()
		file1 := tempDir + "/file1.txt"
		file2 := tempDir + "/file2.txt"

		os.WriteFile(file1, []byte("file1 content"), 0644)
		os.WriteFile(file2, []byte("file2 content"), 0644)

		files := []string{file1, file2}
		fNames, modTimes := file.GetFileStatus(files)

		if len(fNames) != 2 || len(modTimes) != 2 {
			t.Fatalf("Expected 2 files with mod times, got %d files and %d mod times", len(fNames), len(modTimes))
		}
	})

	t.Run("Non-existent Files", func(t *testing.T) {
		files := []string{"nonexistent1.txt", "nonexistent2.txt"}
		fNames, modTimes := file.GetFileStatus(files)

		if fNames != nil || modTimes != nil {
			t.Fatalf("Expected nil for both filenames and mod times for non-existent files")
		}
	})

	t.Run("Empty File List", func(t *testing.T) {
		fNames, modTimes := file.GetFileStatus([]string{})

		if len(fNames) != 0 || len(modTimes) != 0 {
			t.Fatalf("Expected empty results for empty file list")
		}
	})
}
