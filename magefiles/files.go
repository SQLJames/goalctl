package main

import "os"

// fileExists returns whether the given file or directory exists or not.
func fileExists(path string) bool {
	_, err := os.Stat(path)

	return err == nil
}
