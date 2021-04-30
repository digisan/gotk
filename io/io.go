package io

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"strings"
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

// DirExists :
func DirExists(dirname string) bool {
	dirname, _ = AbsPath(dirname, false)
	info, err := os.Stat(dirname)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}

// FileExists :
func FileExists(filename string) bool {
	filename, _ = AbsPath(filename, false)
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// FilesAllExist :
func FilesAllExist(filenames []string) bool {
	for _, filename := range filenames {
		if !FileExists(filename) {
			return false
		}
	}
	return len(filenames) > 0
}

// DirsAllExist :
func DirsAllExist(dirnames []string) bool {
	for _, dirname := range dirnames {
		if !DirExists(dirname) {
			return false
		}
	}
	return len(dirnames) > 0
}

// AbsPath : if check(false), error always nil
func AbsPath(path string, check bool) (string, error) {
	if sHasPrefix(path, "~/") {
		user, err := user.Current()
		if err != nil {
			log.Fatalf("%v", err)
		}
		path = user.HomeDir + path[1:]
	}
	abspath, err := filepath.Abs(path)
	if err != nil {
		log.Fatalf("%v", err)
	}

	if check && (!DirExists(abspath) && !FileExists(abspath)) {
		return abspath, fmt.Errorf("%s doesn't exist", abspath)
	}
	return abspath, nil
}

// MustCreateDir :
func MustCreateDir(dir string) {
	dir, _ = AbsPath(dir, false)
	filename := dir + "/MustCreateDir.temp"
	MustWriteFile(filename, []byte{})
	if err := os.Remove(filename); err != nil {
		log.Fatalf("%v", err)
	}
}

// MustWriteFile :
func MustWriteFile(filename string, data []byte) {

	dir, _ := AbsPath(filepath.Dir(filename), false)
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

WRITE:
	if err := os.WriteFile(filename, data, FilePerm); err != nil {
		log.Fatalf("Could NOT Write File: %v", err)
	}
}

// MustAppendFile :
func MustAppendFile(filename string, data []byte, newline bool) {

	filename, _ = AbsPath(filename, false)
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

// FileIsEmpty :
func FileIsEmpty(filename string) (bool, error) {
	filename, err := AbsPath(filename, true)
	if err != nil {
		return true, err
	}

	info, err := os.Stat(filename)
	if err != nil {
		log.Fatalf("Could NOT Get file Status: %v", err)
	}
	return info.Size() == 0, nil
}

// DirIsEmpty :
func DirIsEmpty(dirname string) (bool, error) {
	dirname, err := AbsPath(dirname, true)
	if err != nil {
		return true, err
	}

	fs, err := os.ReadDir(dirname)
	if err != nil {
		log.Fatalf("Could NOT ReadDir: %v", err)
	}
	return len(fs) == 0, nil
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
	filepath, err := AbsPath(filepath, true)
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

// FileDirCount : ignore hidden file or directory
func FileDirCount(dirname string, recursive bool, exctypes ...string) (fileCount, dirCount int, err error) {

	dirname, err = AbsPath(dirname, true)
	if err != nil {
		return -1, -1, err
	}

	if !recursive {

		files, err := os.ReadDir(dirname)
		if err != nil {
			return -1, -1, err
		}
	NEXT_FILE:
		for _, file := range files {
			filename := file.Name()
			// ignore hidden file or directory
			if sHasPrefix(filename, ".") {
				continue
			}
			// ignore excluded files
			for _, exc := range exctypes {
				if sHasSuffix(filename, "."+exc) {
					continue NEXT_FILE
				}
			}
			path := dirname + filename
			if FileExists(path) {
				fileCount++
			} else {
				dirCount++
			}
		}

	} else {

		if err = filepath.Walk(dirname,
			func(path string, info os.FileInfo, err error) error {
				if err != nil {
					return err
				}
				filename := info.Name()
				// ignore hidden file or directory
				if sHasPrefix(filename, ".") {
					return nil
				}
				// ignore excluded files
				for _, exc := range exctypes {
					if sHasSuffix(filename, "."+exc) {
						return nil
					}
				}
				if FileExists(path) {
					fileCount++
				} else {
					dirCount++
				}
				return nil
			}); err != nil {
			return -1, -1, err
		}

	}

	return
}
