package util

import (
	"fmt"
	"io/fs"
	"os"
	"path"

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

		return "", fmt.Errorf("OS UserHomeDir: %w", err)
	}

	return dirname, nil
}

func MakeStorageLocation() (storageLocation string) {
	applicationName := info.GetApplicationName()
	baseDir, err := GetHomeDir()
	if err != nil {
		log.Logger.Error(err, "Error getting home directory, setting location to tmp folder.")

		baseDir = os.TempDir()
	}

	storageLocation = path.Join(baseDir, "."+applicationName)

	err = os.MkdirAll(storageLocation, folderPermissons)
	if err != nil {
		log.Logger.Fatal(err, "error creating storagelocation", "location", storageLocation)
	}

	return storageLocation
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
