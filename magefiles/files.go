package main

import "os"

// fileExists returns whether the given file or directory exists or not
func fileExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return true
}

//func fileSize(path string) int64 {
//	f, err := os.Stat(path)
//	if err != nil {
//		return 0
//	}
//
//	return f.Size()
//}
