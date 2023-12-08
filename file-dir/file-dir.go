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
	"os/exec"
	"os/user"
	"path/filepath"
	"strings"

	. "github.com/digisan/go-generics/v2"
	"github.com/digisan/gotk/strs"
	"github.com/h2non/filetype"
)

const (
	// File Type
	Document    = "document"
	Image       = "image"
	Audio       = "audio"
	Video       = "video"
	Archive     = "archive"
	Application = "application"
	Executable  = "executable"
	Font        = "font"
	Text        = "text"
	Unknown     = "unknown"
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

// "a/b/c/d.txt" => "a/b/c/prefix-d-suffix.txt"
// "a/b/c/d" => "a/b/c/prefix-d-suffix"
func ChangeFileName(fPath, prefix, suffix string) string {
	ext := filepath.Ext(fPath)
	fName := filepath.Base(fPath)
	name := fName[:len(fName)-len(ext)]
	nameNew := prefix + name + suffix + ext
	return filepath.Join(filepath.Dir(fPath), nameNew)
}

// e.g. fPath is "/a/b/c/d.txt",
// if [fromLast] is 1 and newTail is [D], return is "/a/b/c/D.txt"
// if [fromLast] is 2 and newTail is [D], return is "/a/b/D.txt"
// if [fromLast] is 1 and newTail is [D/E.txt] return is "/a/b/c/D/E.txt"
func ChangeFilePath(strict bool, fPath, newTail string, fromLast int, keepExt, cp, mv bool) string {

	if strict && !FileExists(fPath) {
		log.Fatalf("[%s] is not existing", fPath)
	}

	sep := string(os.PathSeparator)

	ext := ""
	if keepExt {
		iLastExt := strings.LastIndex(fPath, ".")
		iLastPart := strings.LastIndex(fPath, sep)
		if iLastExt > iLastPart {
			ext = fPath[iLastExt:]
		}
	}

	newPath := ""
	head, tail := "", strs.SplitPartFromLastTo[string](fPath, sep, fromLast)
	if fromLast == 1 {
		head = strings.TrimRight(fPath, tail)
		newPath = head + newTail + ext
	} else {
		idx := strings.LastIndex(fPath, sep+tail+sep)
		head = fPath[:idx]
		newPath = head + sep + newTail + ext
	}

	if !cp && !mv {
		return newPath
	}

	if FileExists(fPath) {
		if err := os.MkdirAll(filepath.Dir(newPath), os.ModePerm); err != nil {
			log.Fatalln(err)
		}
		switch {
		case cp:
			buf, err := os.ReadFile(fPath)
			if err != nil {
				log.Fatalln(err)
			}
			if err = os.WriteFile(newPath, buf, os.ModePerm); err != nil {
				log.Fatalln(err)
			}
		case mv:
			if err := os.Rename(fPath, newPath); err != nil {
				log.Fatalln(err)
			}
		}
	}

	return newPath
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
	for _, fName := range paths {
		if !FileExists(fName) {
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
	absPath, err := filepath.Abs(path)
	if err != nil {
		log.Fatalf("%v", err)
	}

	if check && (!DirExists(absPath) && !FileExists(absPath)) {
		return absPath, fmt.Errorf("%s doesn't exist", absPath)
	}
	return absPath, nil
}

// RelPath : if check(false), error always nil
func RelPath(path string, check bool) (string, error) {
	basePath, err := AbsPath(filepath.Dir(os.Args[0]), check)
	if err != nil {
		return "", err
	}
	targPath, err := AbsPath(path, check)
	if err != nil {
		return "", err
	}
	// fmt.Println("basePath: ", basePath)
	// fmt.Println("target:   ", targPath)
	return filepath.Rel(basePath, targPath)
}

func AncestorList(path string) (ancestors []string) {
	absPath, _ := AbsPath(path, false)
	for {
		// fmt.Println(absPath)
		absPath = filepath.Dir(absPath)
		if absPath == "/" || absPath[1:] == ":\\" {
			break
		}
		ancestors = append(ancestors, absPath)
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
func WalkFileDir(path string, recursive bool, excTypes ...string) (filepaths, directories []string, err error) {

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

			fName := file.Name()

			// ignore hidden file or directory
			if strings.HasPrefix(fName, ".") {
				continue
			}

			// ignore excluded files
			for _, exc := range excTypes {
				if strings.HasSuffix(fName, "."+exc) {
					continue NEXT_FILE
				}
			}

			if fp := filepath.Join(path, fName); FileExists(fp) {
				filepaths = append(filepaths, fp)
			} else {
				directories = append(directories, fp)
			}
		}

	} else {

		skipSelf := true
		if err = filepath.WalkDir(path,
			func(path string, d fs.DirEntry, err error) error {
				if err != nil {
					return err
				}

				fName := d.Name()

				// ignore self folder
				if skipSelf {
					skipSelf = false
					return nil
				}

				// ignore hidden file or directory
				if strings.HasPrefix(fName, ".") {
					return nil
				}

				// ignore any files under hidden directory
				for _, pathSeg := range AncestorList(path) {
					if strings.HasPrefix(pathSeg, ".") {
						return nil
					}
				}

				// ignore excluded files
				for _, exc := range excTypes {
					if strings.HasSuffix(fName, "."+exc) {
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
func MergeDir(destDir string, move bool, onConflict func(existing, incoming []byte) (overwrite bool, overwriteData []byte), srcDirs ...string) error {

	destDir, _ = AbsPath(destDir, false)

	for _, srcDir := range srcDirs {
		srcDir, err := AbsPath(srcDir, true)
		if err != nil {
			return err
		}

		files, dirs, err := WalkFileDir(srcDir, true)
		if err != nil {
			return err
		}

		// create each folder
		for _, dir := range dirs {
			aimDir := filepath.Clean(destDir) + dir[len(srcDir):]
			os.MkdirAll(aimDir, 0700)
		}

		// create non-empty folders, including self folder
		dirs = []string{}
		for _, file := range files {
			dirs = append(dirs, filepath.Dir(file))
		}
		for _, dir := range Settify(dirs...) {
			aimDir := filepath.Clean(destDir) + dir[len(srcDir):]
			os.Mkdir(aimDir, 0700)
		}

		// fmt.Println("src-dir:", srcDir)
		// fmt.Println("dest-dir:", destDir)

		// copy files
	NEXT_FILE:
		for _, file := range files {
			destData, err := os.ReadFile(file)
			if err != nil {
				return err
			}

			destFile := strings.Replace(file, srcDir, destDir, 1)

			if FileExists(destFile) && onConflict != nil {
				// fmt.Printf("conflict at: %s\n", destFile)
				existing, err := os.ReadFile(destFile)
				if err != nil {
					return err
				}
				overwrite, data := onConflict(existing, destData)
				if !overwrite {
					continue NEXT_FILE
				}
				destData = data
			}

			if err := os.WriteFile(destFile, destData, os.ModePerm); err != nil {
				return err
			}
		}

		if move {
			os.RemoveAll(srcDir)
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

// Wildcard applies in file name
func Remove(path string, rmEmptyDir bool) error {
	path, err := AbsPath(path, false)
	if err != nil {
		return err
	}

	if name := filepath.Base(path); strs.ContainsAny(name, "*", "?") {
		matched, err := filepath.Glob(filepath.Join(filepath.Dir(path), name))
		if err != nil {
			return err
		}
		for _, path := range matched {
			if err := os.RemoveAll(path); err != nil {
				return err
			}
		}
	} else {
		if err := os.RemoveAll(path); err != nil {
			return err
		}
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

func MustCreateDir(path string) {
	path, _ = AbsPath(path, false)
	fPath := path + "/MustCreateDir.temp"
	MustWriteFile(fPath, []byte{})
	if err := os.Remove(fPath); err != nil {
		log.Fatalf("%v", err)
	}
}

func MustCreateDirs(paths ...string) {
	for _, path := range paths {
		MustCreateDir(path)
	}
}

func MustWriteFile(path string, data []byte) {
	dir, _ := AbsPath(filepath.Dir(path), false)
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

func MustAppendFile(path string, data []byte, newline bool) {
	path, _ = AbsPath(path, false)
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

// FileLineScan :
func FileLineScan(path string, f func(line string) (bool, string), outFile string) (string, error) {
	path, err := AbsPath(path, true)
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
	content, err := strs.ScanLine(file, f)
	if err != nil {
		return "", err
	}
	if outFile != "" {
		MustWriteFile(outFile, []byte(content))
	}
	return content, nil
}

// FileLineScanEx :
func FileLineScanEx(path string, nAbove, nBelow int, junkLine string, f func(line string, cache []string) (bool, string), outFile string) (string, error) {
	path, err := AbsPath(path, true)
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
	content, err := strs.ScanLineEx(file, nAbove, nBelow, junkLine, f)
	if err != nil {
		return "", err
	}
	if outFile != "" {
		MustWriteFile(outFile, []byte(content))
	}
	return content, nil
}

func SupportedFileTypes() []string {
	return []string{Document, Image, Audio, Video, Archive, Application, Executable, Font, Text, Unknown}
}

func IsSupportedFileType(fType string) bool {
	fType = strings.ToLower(fType)
	return In(fType, SupportedFileTypes()...)
}

func IsTextFile(fPath string) bool {
	cmd, err := exec.Command("file", fPath).Output()
	if err != nil {
		log.Printf("error @IsTextFile: %s\n", err)
		return false
	}
	output := strings.TrimSpace(string(cmd))
	return strings.HasSuffix(output, " text") || strings.Contains(output, " text ")
}

// if return "unknown", check it is text file
func FileType(f io.ReadSeeker) string {
	defer f.Seek(0, io.SeekStart)

	head := make([]byte, 261)
	f.Read(head)
	switch {
	case filetype.IsImage(head):
		return Image
	case filetype.IsVideo(head):
		return Video
	case filetype.IsAudio(head):
		return Audio
	case filetype.IsDocument(head):
		return Document
	case filetype.IsArchive(head):
		return Archive
	case filetype.IsApplication(head):
		return Application
	case filetype.IsFont(head):
		return Font
	default:
		{
			f.Seek(0, io.SeekStart)
			data, err := io.ReadAll(f)
			if err != nil {
				return Unknown
			}
			tempFile := "/tmp/temp4filetype"
			if err := os.WriteFile(tempFile, data, os.ModePerm); err != nil {
				return Unknown
			}
			defer os.RemoveAll(tempFile)
			if IsTextFile(tempFile) {
				return Text
			}
		}
		return Unknown
	}
}

func CopyFile(src, dst string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	destFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, srcFile)
	if err != nil {
		return err
	}

	return nil
}

func CopyFolder(src, dst string, inclExt ...string) error {
	// Create the destination folder if it doesn't exist
	if err := os.MkdirAll(dst, 0755); err != nil {
		return err
	}

	// Open the source folder
	srcFolder, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFolder.Close()

	// Read the contents of the source folder
	fileInfos, err := srcFolder.Readdir(-1)
	if err != nil {
		return err
	}

	for i, ext := range inclExt {
		inclExt[i] = "." + strings.TrimPrefix(ext, ".")
	}

	// Iterate through the files and copy them
	for _, fileInfo := range fileInfos {

		name := fileInfo.Name()
		if ext := filepath.Ext(name); len(inclExt) > 0 && NotIn(ext, inclExt...) {
			continue
		}

		srcPath := filepath.Join(src, name)
		dstPath := filepath.Join(dst, name)

		if fileInfo.IsDir() {
			// If the file is a directory, recursively copy the folder
			if err := CopyFolder(srcPath, dstPath); err != nil {
				return err
			}
		} else {
			// If the file is a regular file, copy it
			if err := CopyFile(srcPath, dstPath); err != nil {
				return err
			}
		}
	}

	return nil
}
