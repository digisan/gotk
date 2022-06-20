package io

import (
	"bufio"
	"bytes"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	fd "github.com/digisan/gotk/filedir"
)

const (
	// FilePerm :
	FilePerm = 0666
	// DirPerm :
	DirPerm = 0777
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

// MustCreateDir :
func MustCreateDir(path string) {

	mtx4crtdir.Lock()
	defer mtx4crtdir.Unlock()

	path, _ = fd.AbsPath(path, false)
	fpath := path + "/MustCreateDir.temp"
	MustWriteFile(fpath, []byte{})
	if err := os.Remove(fpath); err != nil {
		log.Fatalf("%v", err)
	}
}

func MustCreateDirs(paths ...string) {
	for _, path := range paths {
		MustCreateDir(path)
	}
}

// MustWriteFile :
func MustWriteFile(path string, data []byte) {

	dir, _ := fd.AbsPath(filepath.Dir(path), false)
	_, err := os.Stat(dir)
	if err != nil && os.IsNotExist(err) {
		if err := os.MkdirAll(dir, DirPerm); err != nil { // dir must be 0777 to put writes in
			log.Fatalf("Could NOT Create File to Write: %v", err)
		}
		goto WRITE
	}
	if err != nil {
		log.Fatalf("Could NOT Get file Status: %v", err)
	}

	path = filepath.Join(dir, filepath.Base(path))
WRITE:
	if err := os.WriteFile(path, data, FilePerm); err != nil {
		log.Fatalf("Could NOT Write File: %v", err)
	}
}

// MustAppendFile :
func MustAppendFile(path string, data []byte, newline bool) {

	path, _ = fd.AbsPath(path, false)
	_, err := os.Stat(path)
	if err != nil && os.IsNotExist(err) {
		MustWriteFile(path, data)
		return
	}
	if err != nil {
		log.Fatalf("Could NOT Get file Status: %v", err)
	}

	file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, FilePerm)
	if err != nil {
		log.Fatalf("Could NOT Open File to Append: %v", err)
	}
	defer file.Close()

	if newline {
		data = append([]byte{'\n'}, data...)
	}
	if _, err = file.Write(data); err != nil {
		log.Fatalf("Could NOT Append File: %v", err)
	}
}

// scanLine :
func scanLine(r io.Reader, f func(line string) (bool, string), outfile string) (string, error) {
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
func FileLineScan(path string, f func(line string) (bool, string), outfile string) (string, error) {
	path, err := fd.AbsPath(path, true)
	if err != nil {
		return "", err
	}
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer func() {
		if file != nil {
			file.Close()
		}
	}()
	return scanLine(file, f, outfile)
}

// StrLineScan :
func StrLineScan(str string, f func(line string) (bool, string), outfile string) (string, error) {
	return scanLine(strings.NewReader(str), f, outfile)
}

func StreamToBytes(stream io.Reader) []byte {
	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(stream)
	if err != nil {
		panic(err)
	}
	return buf.Bytes()
}

func StreamToString(stream io.Reader) string {
	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(stream)
	if err != nil {
		panic(err)
	}
	return buf.String()
}
