package io

import (
	"bufio"
	"io"
	"os"
	"strings"

	"github.com/cdutwhu/debog/base"
)

var (
	// MustWriteFile : from debog/base
	MustWriteFile = base.MustWriteFile

	// MustAppendFile : from debog/base
	MustAppendFile = base.MustAppendFile
)

// FileIsEmpty :
func FileIsEmpty(filename string) bool {
	info, err := os.Stat(filename)
	if err != nil {
		panic(err)
	}
	return info.Size() == 0
}

// DirIsEmpty :
func DirIsEmpty(dirname string) bool {
	fs, err := os.ReadDir(dirname)
	if err != nil {
		panic(err)
	}
	return len(fs) == 0
}

// FileExists :
func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// DirExists :
func DirExists(dirname string) bool {
	info, err := os.Stat(dirname)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}

// readByLine :
func readByLine(r io.Reader, f func(line string) (bool, string), outfile string) (string, error) {
	lines := []string{}
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		if ok, line := f(scanner.Text()); ok {
			lines = append(lines, line)
		}
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}
	content := sJoin(lines, "\n")
	if outfile != "" {
		MustWriteFile(outfile, []byte(content))
	}
	return content, nil
}

// FileLineScan :
func FileLineScan(filepath string, f func(line string) (bool, string), outfile string) (string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return "", err
	}
	defer file.Close()
	return readByLine(file, f, outfile)
}

// StrLineScan :
func StrLineScan(str string, f func(line string) (bool, string), outfile string) (string, error) {
	return readByLine(strings.NewReader(str), f, outfile)
}
