package goutils_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/mekramy/goutils"
)

func TestNormalizePath(t *testing.T) {
	expected := filepath.Clean("/a//b\\c")
	result := goutils.NormalizePath("/a", "b", "c")
	if result != expected {
		t.Errorf("expected %s, got %s", expected, result)
	}
}

func TestCreateDirectory(t *testing.T) {
	dir := "testdir/1/2"
	err := goutils.CreateDirectory(dir)
	if err != nil {
		t.Errorf("failed to create directory: %v", err)
	}
	defer os.RemoveAll(dir)

	info, err := os.Stat(dir)
	if os.IsNotExist(err) {
		t.Errorf("directory does not exist")
	}
	if !info.IsDir() {
		t.Errorf("expected a directory, got a file")
	}
}

func TestIsDirectory(t *testing.T) {
	dir := "testdir"
	os.Mkdir(dir, os.ModePerm)
	defer os.RemoveAll(dir)

	isDir, err := goutils.IsDirectory(dir)
	if err != nil {
		t.Errorf("failed to check if path is directory: %v", err)
	}
	if !isDir {
		t.Errorf("expected a directory, got a file")
	}
}

func TestGetSubDirectory(t *testing.T) {
	dir := "testdir"
	subDir := filepath.Join(dir, "subdir")
	os.MkdirAll(subDir, os.ModePerm)
	defer os.RemoveAll(dir)

	subDirs, err := goutils.GetSubDirectory(dir)
	if err != nil {
		t.Errorf("failed to get sub directories: %v", err)
	}
	if len(subDirs) != 1 || subDirs[0] != "subdir" {
		t.Errorf("expected [subdir], got %v", subDirs)
	}
}

func TestClearDirectory(t *testing.T) {
	dir := "testdir"
	os.MkdirAll(filepath.Join(dir, "subdir"), os.ModePerm)
	defer os.RemoveAll(dir)

	err := goutils.ClearDirectory(dir)
	if err != nil {
		t.Errorf("failed to clear directory: %v", err)
	}

	files, err := os.ReadDir(dir)
	if err != nil {
		t.Errorf("failed to read directory: %v", err)
	}
	if len(files) != 0 {
		t.Errorf("expected empty directory, got %v", files)
	}
}

func TestFileExists(t *testing.T) {
	file := "testfile.txt"
	os.WriteFile(file, []byte("content"), os.ModePerm)
	defer os.Remove(file)

	exists, err := goutils.FileExists(file)
	if err != nil {
		t.Errorf("failed to check if file exists: %v", err)
	}
	if !exists {
		t.Errorf("expected file to exist")
	}
}

func TestFindFile(t *testing.T) {
	dir := "testdir"
	file := filepath.Join(dir, "testfile.txt")
	os.MkdirAll(dir, os.ModePerm)
	os.WriteFile(file, []byte("content"), os.ModePerm)
	defer os.RemoveAll(dir)

	result := goutils.FindFile(dir, "testfile.txt")
	if result == nil || *result != file {
		t.Errorf("expected %s, got %v", file, result)
	}
}

func TestFindFiles(t *testing.T) {
	dir := "testdir"
	file1 := filepath.Join(dir, "testfile1.txt")
	file2 := filepath.Join(dir, "testfile2.txt")
	os.MkdirAll(dir, os.ModePerm)
	os.WriteFile(file1, []byte("content"), os.ModePerm)
	os.WriteFile(file2, []byte("content"), os.ModePerm)
	defer os.RemoveAll(dir)

	result := goutils.FindFiles(dir, "testfile.*.txt")
	if len(result) != 2 || result[0] != file1 || result[1] != file2 {
		t.Errorf("expected [%s, %s], got %v", file1, file2, result)
	}
}

func TestGetMime(t *testing.T) {
	data := []byte("test content")
	expected := "text/plain; charset=utf-8"
	mime := goutils.GetMime(data)
	if mime.String() != expected {
		t.Errorf("expected %s, got %s", expected, mime.String())
	}
}

func TestGetExtension(t *testing.T) {
	file := "testfile.txt"
	expected := "txt"
	result := goutils.GetExtension(file)
	if result != expected {
		t.Errorf("expected %s, got %s", expected, result)
	}
}

func TestGetFilename(t *testing.T) {
	file := "testfile.txt"
	expected := "testfile"
	result := goutils.GetFilename(file)
	if result != expected {
		t.Errorf("expected %s, got %s", expected, result)
	}
}

func TestTimestampedFile(t *testing.T) {
	file := "testfile.txt"
	result := goutils.TimestampedFile(file)
	if !strings.HasSuffix(result, "-testfile.txt") {
		t.Errorf("expected suffix -testfile.txt, got %s", result)
	}
}

func TestNumberedFile(t *testing.T) {
	dir := "testdir"
	file := "testfile.txt"
	os.MkdirAll(dir, os.ModePerm)
	os.WriteFile(dir+"/"+file, []byte("TEST"), 0644)
	defer os.RemoveAll(dir)

	result, err := goutils.NumberedFile(dir, file)
	if err != nil {
		t.Errorf("failed to generate numbered file: %v", err)
	}
	if result != "testfile-1.txt" {
		t.Errorf("expected testfile-1.txt, got %s", result)
	}
}
