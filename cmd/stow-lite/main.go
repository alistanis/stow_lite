package main

import (
	"fmt"
	"os"

	"flag"

	stow "github.com/alistanis/stow_lite"
)

var (
	sourceDir      string
	destinationDir string

	suppressErrors        bool
	supressAndPrintErrors bool

	exclusionPattern string
)

func init() {
	flag.BoolVar(&suppressErrors, "s", false, "Supresses errors")
	flag.BoolVar(&supressAndPrintErrors, "sv", false, "Supresses and prints errors")
	flag.StringVar(&exclusionPattern, "z", "", "Excludes this pattern (regex)")
	flag.Parse()
	if len(flag.Args()) < 2 {
		fmt.Println("Must provide a source directory and a destination directory")
		usage()
		os.Exit(-1)
	}
	sourceDir = flag.Args()[0]
	destinationDir = flag.Args()[1]

}

func usage() {
	fmt.Println("stow-lite usage:")
	fmt.Println("\tstow [OPTIONS...] [source dir] [destination dir]")
}

func main() {
	opts := &stow.Options{ErrorBehavior: stow.ExitError, ExclusionPattern: exclusionPattern}

	if suppressErrors {
		opts.ErrorBehavior = stow.SuppressErrors
	}
	if supressAndPrintErrors {
		opts.ErrorBehavior = stow.SuppressAndPrintErrors
	}
	stow.SetOptions(opts)
	err := stow.CreateSymlinks(sourceDir, destinationDir)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
