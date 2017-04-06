package stow_lite

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

const (
	SuppressErrors = iota
	SuppressAndPrintErrors
	ExitError
)

var (
	options *Options
)

type Options struct {
	ErrorBehavior    int
	ExclusionPattern string
}

func SetOptions(o *Options) {
	options = o
}

func init() {
	options = &Options{ErrorBehavior: ExitError}
}

func CreateSymlinks(oldDir, newDir string) error {

	visit := func(path string, f os.FileInfo, err error) error {
		if f.IsDir() {
			return nil
		}

		if options.ExclusionPattern != "" {
			match, err := regexp.MatchString(options.ExclusionPattern, f.Name())
			if err != nil {
				return err
			}
			if match {
				fmt.Println("skipping " + f.Name() + " due to exlusion pattern: " + options.ExclusionPattern)
				return nil
			}
		}

		oldDir = fixPath(oldDir)
		newDir = fixPath(newDir)
		e := os.Symlink(oldDir+f.Name(), newDir+f.Name())
		if e == nil {
			fmt.Println(oldDir + f.Name() + " -> " + newDir + f.Name())
			return nil
		}
		switch options.ErrorBehavior {
		case SuppressErrors:
			return nil
		case SuppressAndPrintErrors:
			fmt.Println(e)
			return nil
		case ExitError:
			fallthrough
		default:
			return e
		}
	}
	return filepath.Walk(oldDir, visit)
}

func fixPath(s string) string {
	if !strings.HasSuffix(s, "/") {
		s += "/"
	}
	return s
}
