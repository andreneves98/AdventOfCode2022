package utils_test

import (
	"advent-of-code/utils"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestReadFile(t *testing.T) {
	var (
		filePath  = "./test.txt"
		wantLines = []string{
			"Hello",
			"World!",
			"Advent of Code",
			"2022",
		}
	)
	f, err := os.Create(filePath)
	if err != nil {
		t.Errorf("didn't expect error; got %v", err)
		return
	}

	defer f.Close()
	_, err = f.WriteString("Hello\nWorld!\nAdvent of Code\n2022")
	if err != nil {
		t.Errorf("didn't expect error; got %v", err)
		return
	}

	// Flush writes to stable storage
	f.Sync()

	// Read file
	lines, err := utils.ReadFile(filePath)
	if err != nil {
		t.Errorf("didn't expect error; got %v", err)
		return
	}

	if diff := cmp.Diff(lines, wantLines); diff != "" {
		t.Errorf("want %v, got %v", wantLines, lines)
		return
	}

	os.Remove(filePath)
}

func TestReadFileWithEmptyLines(t *testing.T) {
	var (
		filePath  = "./test.txt"
		wantLines = []string{
			"Hello",
			"World!",
			"",
			"Advent of Code",
			"2022",
		}
	)
	f, err := os.Create(filePath)
	if err != nil {
		t.Errorf("didn't expect error; got %v", err)
		return
	}

	defer f.Close()
	_, err = f.WriteString("Hello\nWorld!\n\nAdvent of Code\n2022\n")
	if err != nil {
		t.Errorf("didn't expect error; got %v", err)
		return
	}

	// Flush writes to stable storage
	f.Sync()

	// Read file
	lines, err := utils.ReadFile(filePath)
	if err != nil {
		t.Errorf("didn't expect error; got %v", err)
		return
	}

	if diff := cmp.Diff(lines, wantLines); diff != "" {
		t.Errorf("want %v, got %v", wantLines, lines)
		if err := os.Remove(filePath); err != nil {
			t.Errorf("didn't expect error; got %v", err)
		}
		return
	}

	if err := os.Remove(filePath); err != nil {
		t.Errorf("didn't expect error; got %v", err)
		return
	}
}

func TestReadFile_FileDoesntExist(t *testing.T) {
	filePath := "./test.txt"

	// Read file that doesn't exist
	_, err := utils.ReadFile(filePath)
	if err == nil {
		t.Errorf("expected error; got %v", err)
		return
	}
}
