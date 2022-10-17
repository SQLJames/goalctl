package main

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/sqljames/goalctl/pkg/util/jlogr"
)

// ModulePath returns the module path from the gomod file text.
func modulePath() string {
	var lines []string

	file, err := os.Open("go.mod")
	if err != nil {
		jlogr.Logger.ILog.Error(err, err.Error())

		return filepath.Base(gitRoot())
	}

	defer func(f *os.File) {
		if err := f.Close(); err != nil {
			jlogr.Logger.ILog.Error(err, err.Error(), "fileName", file.Name(), "error", err.Error())

		}
	}(file)

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		// closing file because we are logging a fatal error
		// log.Fatal will exit, and `defer file.Close()` will not run
		log.Println(err)

		return filepath.Base(gitRoot())
	}

	return parseFileLines(lines)
}

func parseFileLines(lines []string) (modulePath string) {
	split := strings.Split(lines[0], " ")

	return split[1]
}
