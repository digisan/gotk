package filedir

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"strings"

	"github.com/digisan/gotk/slice/ts"
)

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

// RelPath : if check(false), error always nil
func RelPath(path string, check bool) (string, error) {
	basepath, err := AbsPath(filepath.Dir(os.Args[0]), check)
	if err != nil {
		return "", err
	}
	targpath, err := AbsPath(path, check)
	if err != nil {
		return "", err
	}
	// fmt.Println("basepath: ", basepath)
	// fmt.Println("target:   ", targpath)
	return filepath.Rel(basepath, targpath)
}

func AncestorList(path string) (ancestors []string) {
	abspath, _ := AbsPath(path, false)
	for {
		// fmt.Println(abspath)
		abspath = filepath.Dir(abspath)
		if abspath == "/" || abspath[1:] == ":\\" {
			break
		}
		ancestors = append(ancestors, abspath)
	}
	ancestors = ts.Reverse(ancestors)
	// fmt.Println(ancestors)

	for i := len(ancestors) - 1; i >= 1; i-- {
		ancestors[i] = strings.TrimPrefix(ancestors[i], ancestors[i-1])
		ancestors[i] = strings.TrimLeft(ancestors[i], "/\\")
	}

	if ancestors[0][1:3] == ":\\" {
		ancestors[0] = ancestors[0][3:]
	} else {
		ancestors[0] = strings.TrimLeft(ancestors[0], "/")
	}

	// fmt.Println(ancestors)
	return
}

func Parent(path string) string {
	ancestors := AncestorList(path)
	if len(ancestors) > 0 {
		return ancestors[len(ancestors)-1]
	}
	return ""
}

func GrandParent(path string) string {
	ancestors := AncestorList(path)
	if len(ancestors) > 1 {
		return ancestors[len(ancestors)-2]
	}
	return ""
}

// WalkFileDir : ignore hidden file or directory
func WalkFileDir(dirname string, recursive bool, exctypes ...string) (filepaths, directories []string, err error) {

	dirname, err = AbsPath(dirname, true)
	if err != nil {
		return nil, nil, err
	}

	if !recursive {

		files, err := os.ReadDir(dirname)
		if err != nil {
			return nil, nil, err
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

			if path := filepath.Join(dirname, filename); FileExists(path) {
				filepaths = append(filepaths, path)
			} else {
				directories = append(directories, path)
			}
		}

	} else {

		skipself := true
		if err = filepath.WalkDir(dirname,
			func(path string, d fs.DirEntry, err error) error {
				if err != nil {
					return err
				}

				filename := d.Name()

				// ignore self folder
				if skipself {
					skipself = false
					return nil
				}

				// ignore hidden file or directory
				if sHasPrefix(filename, ".") {
					return nil
				}

				// ignore any files under hidden directory
				for _, pathSeg := range AncestorList(path) {
					if sHasPrefix(pathSeg, ".") {
						return nil
					}
				}

				// ignore excluded files
				for _, exc := range exctypes {
					if sHasSuffix(filename, "."+exc) {
						return nil
					}
				}

				if FileExists(path) {
					filepaths = append(filepaths, path)
				} else {
					directories = append(directories, path)
				}
				return nil

			}); err != nil {

			return nil, nil, err
		}
	}

	return
}

// MergeDir: if onConflict == nil, later overwrites previous when conflict
func MergeDir(destdir string, onConflict func(existing, incoming []byte) (overwrite bool, overwriteData []byte), srcdirs ...string) error {

	destdir, _ = AbsPath(destdir, false)

	for _, srcdir := range srcdirs {
		srcdir, err := AbsPath(srcdir, true)
		if err != nil {
			return err
		}

		files, dirs, err := WalkFileDir(srcdir, true)
		if err != nil {
			return err
		}

		// create each folder
		for _, dir := range dirs {
			aimdir := filepath.Clean(destdir) + dir[len(srcdir):]
			os.MkdirAll(aimdir, 0700)
		}

		// create non-empty folders, including self folder
		dirs = []string{}
		for _, file := range files {
			dirs = append(dirs, filepath.Dir(file))
		}
		for _, dir := range ts.MkSet(dirs...) {
			aimdir := filepath.Clean(destdir) + dir[len(srcdir):]
			os.Mkdir(aimdir, 0700)
		}

		// fmt.Println("src-dir:", srcdir)
		// fmt.Println("dest-dir:", destdir)

		// copy files
	NEXT_FILE:
		for _, file := range files {
			destdata, err := os.ReadFile(file)
			if err != nil {
				return err
			}

			destfile := strings.Replace(file, srcdir, destdir, 1)

			if FileExists(destfile) && onConflict != nil {
				// fmt.Printf("conflict at: %s\n", destfile)
				existing, err := os.ReadFile(destfile)
				if err != nil {
					return err
				}
				overwrite, data := onConflict(existing, destdata)
				if !overwrite {
					continue NEXT_FILE
				}
				destdata = data
			}

			if err := os.WriteFile(destfile, destdata, os.ModePerm); err != nil {
				return err
			}
		}
	}
	return nil
}
