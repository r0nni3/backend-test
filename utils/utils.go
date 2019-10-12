package utils

import (
	"errors"
	"io/ioutil"
	//"net/http"
	"os"
	"path/filepath"
	"strings"
)

// ParseCLIArgs parses cli args
// This could be done with a more sofisticated library like cobra
// but this example is simple enough that doesn't make sense to
// install it just to parse one args
func ParseCLIArgs() (target string, err error) {
	if len(os.Args) < 2 {
		err = errors.New("A path or file should be provided to process")
		return
	}

	target = os.Args[1]
	return
}

// CheckTarget Validates that the target file exists and checks its type
// this is help full to know if the target is a file, directory, url, etc...
func CheckTarget(target string) (targets []string, local bool, isDir bool, err error) {
	// TODO: return struct that represents target abstraction, for now
	// this should sufice

	// Local file check
	file, err := os.Stat(target)
	if err == nil {
		local = true
		// Directory not valid
		dir, _ := filepath.Abs(filepath.Dir(target))
		// path base
		base := filepath.Base(target)

		isDir = file.IsDir()

		// full path
		path := dir
		if !isDir {
			path = filepath.Join(path, base)
			targets = append(targets, path)
		}

		if isDir {
			files, err2 := ioutil.ReadDir(path)
			if err2 != nil {
				err = err2
				return
			}

			for _, f := range files {
				target := filepath.Join(path, f.Name())
				targets = append(targets, target)
			}
		}

		return
	}

	// Remote file, tries determine file type for other "types"
	if os.IsNotExist(err) {
		err = nil
		targets = append(targets, target)
	}

	return
}

// GetFileType returs file type
func GetFileType(target string) (string, error) {
	if target == "" {
		return "", errors.New("Empty target has no extension")
	}

	fileType := filepath.Ext(target)
	fileType = strings.Replace(fileType, ".", "", -1)
	fileType = strings.ToLower(fileType)

	return fileType, nil
}
