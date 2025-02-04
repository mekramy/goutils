package goutils

import (
	"fmt"
	"math"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/gabriel-vasile/mimetype"
)

// NormalizePath join and normalize file path.
func NormalizePath(path ...string) string {
	return filepath.Clean(filepath.Join(path...))
}

// CreateDirectory create nested directory.
func CreateDirectory(path string) error {
	return os.MkdirAll(path, os.ModePerm)
}

// IsDirectory check if path is directory
func IsDirectory(path string) (bool, error) {
	stat, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false, err
	} else if err != nil {
		return false, err
	}
	return stat.IsDir(), nil
}

// GetSubDirectory returns list of sub directories.
func GetSubDirectory(dir string) ([]string, error) {
	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var res []string
	for _, file := range files {
		if file.IsDir() {
			res = append(res, file.Name())
		}
	}
	return res, nil
}

// ClearDirectory delete all files and sub-directory in directory.
func ClearDirectory(dir string) error {
	d, err := os.Open(dir)
	if err != nil {
		return err
	}
	defer d.Close()

	names, err := d.Readdirnames(-1)
	if err != nil {
		return err
	}

	for _, name := range names {
		err = os.RemoveAll(filepath.Join(dir, name))
		if err != nil {
			return err
		}
	}
	return nil
}

// FileExists check if file exists.
func FileExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	} else if os.IsNotExist(err) {
		return false, nil
	} else {
		return false, err
	}
}

// FindFile search directory for file with pattern and returns first file.
func FindFile(dir string, pattern string) *string {
	var result string

	// Create regex
	rx, err := regexp.Compile(pattern)
	if err != nil {
		return nil
	}

	// Search for file
	err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && rx.MatchString(info.Name()) {
			result = path
			return filepath.SkipAll
		}

		return nil
	})

	// Handle result
	if err != nil || result == "" {
		return nil
	}
	return &result
}

// FindFiles search directory for files with pattern.
func FindFiles(dir string, pattern string) []string {
	var result []string

	// Create regex
	rx, err := regexp.Compile(pattern)
	if err != nil {
		return nil
	}

	// Search for file
	err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && rx.MatchString(info.Name()) {
			result = append(result, path)
		}

		return nil
	})

	// Handle result
	if err != nil || len(result) == 0 {
		return nil
	}
	return result
}

// GetMime returns file mime info from content
func GetMime(data []byte) *mimetype.MIME {
	return mimetype.Detect(data)
}

// GetExtension returns file extension.
func GetExtension(file string) string {
	return strings.TrimLeft(
		strings.ToLower(filepath.Ext(file)),
		".",
	)
}

// GetFilename returns file name without extension.
func GetFilename(file string) string {
	return strings.TrimSuffix(
		filepath.Base(file),
		filepath.Ext(file),
	)
}

// TimestampedFile returns file name with timestamp prefix.
func TimestampedFile(file string) string {
	name := GetFilename(file)
	ext := filepath.Ext(file)
	return name + "-" + strconv.FormatInt(time.Now().UnixMilli(), 10) + ext
}

// NumberedFile generate unique numbered file name (e.g. file.txt file-1.txt, file-2.txt).
func NumberedFile(dir, file string) (string, error) {
	name := GetFilename(file)
	ext := filepath.Ext(file)

	// Return current name if available
	exists, err := FileExists(NormalizePath(dir, name+ext))
	if err != nil {
		return "", err
	} else if !exists {
		return name + ext, nil
	}

	// Generate filename
	for i := 1; i < math.MaxUint32; i++ {
		counter := strconv.Itoa(i)
		res := name + "-" + counter + ext
		exists, err := FileExists(NormalizePath(dir, res))
		if err != nil {
			return "", err
		} else if !exists {
			return res, nil
		}
	}

	return "", fmt.Errorf("try %d name failed", math.MaxUint32)
}
