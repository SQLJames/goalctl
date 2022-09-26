package util

import (
	"fmt"
	"io/fs"
	"os"
	"path"
	"path/filepath"

	"github.com/sqljames/goalctl/pkg/info"
	"github.com/sqljames/goalctl/pkg/log"
)

var (
	folderPermissons fs.FileMode = 0o755
)

func GetHomeDir() (directory string, err error) {
	dirname, err := os.UserHomeDir()
	if err != nil {
		log.Logger.Error(err, "getting home directory")

		return "", err
	}

	return dirname, nil
}

func MakeStorageLocation() (storageLocation string, err error) {
	homedir, err := GetHomeDir()
	if err != nil {
		return "", err
	}
	applicationName := info.GetApplicationName()
	storageLocation = path.Join(homedir, "."+applicationName)
	err = os.MkdirAll(storageLocation, folderPermissons)
	if err != nil {
		log.Logger.Error(err, "error creating storagelocation")

		return "", err
	}
	return storageLocation, nil
}

func CreateifNotexist(file string) (err error) {
	cleanedFile := filepath.Clean(file)
	if !FileExists(cleanedFile) {
		file, err := os.Create(cleanedFile)
		defer closeFile(file)

		if err != nil {
			log.Logger.Error(err, "unable to create file")

			return err
		}
	}

	return nil
}

func FileExists(fileName string) (exists bool) {
	if _, err := os.Stat(fileName); err != nil {
		return false
	}

	return true
}

func JoinPath(basePath, leaf string) (fullPath string) {
	joinedPath := path.Join(basePath, leaf)
	log.Logger.Trace(fmt.Sprintf("path is %s", joinedPath))
	return joinedPath
}

func closeFile(f *os.File) {
	if err := f.Close(); err != nil {
		log.Logger.Warn("issue closing file", "fileName", f.Name(), "error", err.Error())
	}
}
