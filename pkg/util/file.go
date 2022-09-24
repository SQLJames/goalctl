package util

import (
	"errors"
	"fmt"
	"os"
	"path"
	"path/filepath"

	"github.com/sqljames/goalctl/pkg/info"
	"github.com/sqljames/goalctl/pkg/log"
)

func GetHomeDir() (directory string, err error) {
	dirname, err := os.UserHomeDir()
	if err != nil {
		log.Logger.Error(err, "getting home directory")
		return "", err
	}
	return dirname, nil
}

func MakeStorageLocation() (StorageLocation string, err error) {
	homedir, err := GetHomeDir()
	if err != nil {
		return "", err
	}
	applicationName := info.GetApplicationName()
	StorageLocation = path.Join(homedir, "." + applicationName)
	err = os.MkdirAll(StorageLocation, os.FileMode(0755))
	if err != nil {
		log.Logger.Error(err, "error creating storagelocation")
		return "", err
	}
	return StorageLocation, nil
}

func CreateifNotexist(file string) (err error) {
	cleanedFile := filepath.Clean(file)
	if !FileExists(cleanedFile) {
		file, err := os.Create(cleanedFile)
		defer func() {
			if err := file.Close(); err != nil {
				log.Logger.Warn("issue closing file", "fileName", file.Name(), "error", err.Error())
			}
		}()

		if err != nil {
			log.Logger.Error(err, "unable to create file")
			return err
		}
	}
	return nil
}

func FileExists(fileName string) bool {
	if _, err := os.Stat(fileName); err == nil {
		log.Logger.Trace("file exists")
		return true

	} else if errors.Is(err, os.ErrNotExist) {
		log.Logger.Trace("file not found")
		return false

	} else {
		return false

	}
}

func JoinPath(basePath, leaf string) (fullPath string) {
	path := path.Join(basePath, leaf)
	log.Logger.Trace(fmt.Sprintf("path is %s", path))
	return path
}
