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
func MustCreateDir(dir string) {

	mtx4crtdir.Lock()
	defer mtx4crtdir.Unlock()

	dir, _ = fd.AbsPath(dir, false)
	filename := dir + "/MustCreateDir.temp"
	MustWriteFile(filename, []byte{})
	if err := os.Remove(filename); err != nil {
		log.Fatalf("%v", err)
	}
}

// MustWriteFile :
func MustWriteFile(filename string, data []byte) {

	dir, _ := fd.AbsPath(filepath.Dir(filename), false)
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

	filename = filepath.Join(dir, filepath.Base(filename))
WRITE:
	if err := os.WriteFile(filename, data, FilePerm); err != nil {
		log.Fatalf("Could NOT Write File: %v", err)
	}
}

// MustAppendFile :
func MustAppendFile(filename string, data []byte, newline bool) {

	filename, _ = fd.AbsPath(filename, false)
	_, err := os.Stat(filename)
	if err != nil && os.IsNotExist(err) {
		MustWriteFile(filename, data)
		return
	}
	if err != nil {
		log.Fatalf("Could NOT Get file Status: %v", err)
	}

	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, FilePerm)
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

// RmFileAndEmptyDir :
func RmFileAndEmptyDir(path string) error {
	if err := os.RemoveAll(path); err != nil {
		return err
	}
	ls := fd.AncestorList(path)
	for i := len(ls); i > 0; i-- {
		p := "/" + filepath.Join(ls[:i]...)
		empty, err := fd.DirIsEmpty(p)
		if err != nil {
			log.Println(err)
			return err
		}
		if empty {
			if err := os.RemoveAll(p); err != nil {
				return err
			}
		} else {
			break
		}
	}
	return nil
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
	filepath, err := fd.AbsPath(filepath, true)
	if err != nil {
		return "", err
	}
	file, err := os.Open(filepath)
	if err != nil {
		return "", err
	}
	defer func() {
		if file != nil {
			file.Close()
		}
	}()
	return readByLine(file, f, outfile)
}

// StrLineScan :
func StrLineScan(str string, f func(line string) (bool, string), outfile string) (string, error) {
	return readByLine(strings.NewReader(str), f, outfile)
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
