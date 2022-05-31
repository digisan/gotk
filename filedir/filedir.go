package filedir

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
	"hash"
	"io"
	"io/fs"
	"log"
	"math"
	"os"
	"os/user"
	"path/filepath"
	"strings"

	. "github.com/digisan/go-generics/v2"
	"github.com/digisan/gotk/strs"
)

// e.g. fpath is "/a/b/c/d.txt",
// if [fromlast] is 1 and newtail is [D], return is "/a/b/c/D.txt"
// if [fromlast] is 2 and newtail is [D], return is "/a/b/D.txt"
// if [fromlast] is 1 and newtail is [D/E.txt] return is "/a/b/c/D/E.txt"
func ChangeFilePath(strict bool, fpath, newtail string, fromlast int, keepext, cp, mv bool) string {

	if strict && !FileExists(fpath) {
		log.Fatalf("[%s] is not existing", fpath)
	}

	sep := string(os.PathSeparator)

	ext := ""
	if keepext {
		iLastExt := strings.LastIndex(fpath, ".")
		iLastPart := strings.LastIndex(fpath, sep)
		if iLastExt > iLastPart {
			ext = fpath[iLastExt:]
		}
	}

	newpath := ""
	head, tail := "", strs.SplitPartFromLast(fpath, sep, fromlast)
	if fromlast == 1 {
		head = strings.TrimRight(fpath, tail)
		newpath = head + newtail + ext
	} else {
		idx := strings.LastIndex(fpath, sep+tail+sep)
		head = fpath[:idx]
		newpath = head + sep + newtail + ext
	}

	if !cp && !mv {
		return newpath
	}

	if FileExists(fpath) {
		if err := os.MkdirAll(filepath.Dir(newpath), os.ModePerm); err != nil {
			log.Fatalln(err)
		}
		switch {
		case cp:
			buf, err := os.ReadFile(fpath)
			if err != nil {
				log.Fatalln(err)
			}
			if err = os.WriteFile(newpath, buf, os.ModePerm); err != nil {
				log.Fatalln(err)
			}
		case mv:
			if err := os.Rename(fpath, newpath); err != nil {
				log.Fatalln(err)
			}
		}
	}

	return newpath
}

// ".txt" => ".txt", "txt" => ".txt", " " => "", "" => ""
func DotExt(ext string) string {
	s0 := strings.Trim(ext, " \t")
	// fmt.Println(s0)
	s1 := strings.TrimLeft(s0, ".")
	// fmt.Println(s1)
	s2 := strings.TrimRight("."+s1, ".")
	// fmt.Println(s2)
	s3 := strings.Trim(s2, " \t")
	// fmt.Println(s3)
	return s3
}

// DirExists :
func DirExists(path string) bool {
	path, _ = AbsPath(path, false)
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}

// FileExists :
func FileExists(path string) bool {
	path, _ = AbsPath(path, false)
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// AllFilesExist:
func AllFilesExist(paths ...string) bool {
	for _, filename := range paths {
		if !FileExists(filename) {
			return false
		}
	}
	return len(paths) > 0
}

// AllDirsExist:
func AllDirsExist(paths ...string) bool {
	for _, path := range paths {
		if !DirExists(path) {
			return false
		}
	}
	return len(paths) > 0
}

// AllExistAsWhole: If not all of paths exist, then remove all of them.
// in other words, all exist OR all gone.
// return string is the first error path
func AllExistAsWhole(paths ...string) (bool, string) {
	allExist, errPath := true, ""
	for _, path := range paths {
		if FileExists(path) || DirExists(path) {
			continue
		}
		allExist = false
		errPath = path
		break
	}
	if !allExist {
		for _, path := range paths {
			if err := Remove(path, true); err != nil {
				log.Fatalln(err)
			}
		}
	}
	return allExist, errPath
}

// IsFileEmpty :
func IsFileEmpty(path string) (bool, error) {
	path, err := AbsPath(path, true)
	if err != nil {
		return true, err
	}

	info, err := os.Stat(path)
	if err != nil {
		log.Fatalf("Could NOT Get file Status: %v", err)
	}
	return info.Size() == 0, nil
}

// IsDirEmpty :
func IsDirEmpty(path string) (bool, error) {
	path, err := AbsPath(path, true)
	if err != nil {
		return true, err
	}

	fs, err := os.ReadDir(path)
	if err != nil {
		log.Fatalf("Could NOT ReadDir: %v", err)
	}
	return len(fs) == 0, nil
}

// AbsPath : if check(false), error always nil
func AbsPath(path string, check bool) (string, error) {
	if filepath.IsAbs(path) {
		return path, nil
	}
	if strings.HasPrefix(path, "~/") {
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
	ancestors = Reverse(ancestors)
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

func DirSize(path, unit string) (float64, error) {
	mUnitScale := map[string]float64{
		"k": 1024,
		"K": 1024,
		"m": 1024 * 1024,
		"M": 1024 * 1024,
		"g": 1024 * 1024 * 1024,
		"G": 1024 * 1024 * 1024,
		"t": 1024 * 1024 * 1024 * 1024,
		"T": 1024 * 1024 * 1024 * 1024,
	}
	var size float64
	err := filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			size += float64(info.Size())
		}
		return err
	})
	if scale, ok := mUnitScale[unit]; ok {
		return math.Ceil(size / scale), err
	}
	return math.Ceil(size), err
}

// WalkFileDir : ignore hidden file or directory
func WalkFileDir(path string, recursive bool, exctypes ...string) (filepaths, directories []string, err error) {

	path, err = AbsPath(path, true)
	if err != nil {
		return nil, nil, err
	}

	if !recursive {

		files, err := os.ReadDir(path)
		if err != nil {
			return nil, nil, err
		}

	NEXT_FILE:
		for _, file := range files {

			filename := file.Name()

			// ignore hidden file or directory
			if strings.HasPrefix(filename, ".") {
				continue
			}

			// ignore excluded files
			for _, exc := range exctypes {
				if strings.HasSuffix(filename, "."+exc) {
					continue NEXT_FILE
				}
			}

			if fp := filepath.Join(path, filename); FileExists(fp) {
				filepaths = append(filepaths, fp)
			} else {
				directories = append(directories, fp)
			}
		}

	} else {

		skipself := true
		if err = filepath.WalkDir(path,
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
				if strings.HasPrefix(filename, ".") {
					return nil
				}

				// ignore any files under hidden directory
				for _, pathSeg := range AncestorList(path) {
					if strings.HasPrefix(pathSeg, ".") {
						return nil
					}
				}

				// ignore excluded files
				for _, exc := range exctypes {
					if strings.HasSuffix(filename, "."+exc) {
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

// MergeDir: if onConflict == nil, overwrites previous when conflict
func MergeDir(destdir string, move bool, onConflict func(existing, incoming []byte) (overwrite bool, overwriteData []byte), srcdirs ...string) error {

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
		for _, dir := range Settify(dirs...) {
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

		if move {
			os.RemoveAll(srcdir)
		}
	}
	return nil
}

// h: [md5.New() / sha1.New() / sha256.New()]
func FileHash(path string, h hash.Hash) string {
	if !FileExists(path) {
		return ""
	}
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	_, err = io.Copy(h, f)
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("%x", h.Sum(nil))
}

func SelfMD5() string {
	return FileHash(os.Args[0], md5.New())
}

func SelfSHA1() string {
	return FileHash(os.Args[0], sha1.New())
}

func SelfSHA256() string {
	return FileHash(os.Args[0], sha256.New())
}

func Remove(path string, rmEmptyDir bool) error {
	path, err := AbsPath(path, false)
	if err != nil {
		return err
	}
	if err := os.RemoveAll(path); err != nil {
		return err
	}
	if rmEmptyDir {
		ls := AncestorList(path)
		for i := len(ls); i > 0; i-- {
			p := string(os.PathSeparator) + filepath.Join(ls[:i]...)
			if DirExists(p) {
				empty, err := IsDirEmpty(p)
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
		}
	}
	return nil
}

func RmFilesIn(path string, recursive, rmEmptyDir bool, exts ...string) error {
	if DirExists(path) {
		files, _, err := WalkFileDir(path, recursive)
		if err != nil {
			return err
		}
		for _, file := range files {
			for _, ext := range exts {
				if len(ext) > 0 && strings.HasSuffix(file, DotExt(ext)) {
					if err = Remove(file, rmEmptyDir); err != nil {
						return err
					}
					break
				}
			}
		}
		return nil
	}
	return fmt.Errorf("directory [%v] is not existing", path)
}
