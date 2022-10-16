package util

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path"
	"path/filepath"

	"github.com/sqljames/goalctl/pkg/info"
	"github.com/sqljames/goalctl/pkg/util/jlogr"
)

var (
	folderPermissons       fs.FileMode = 0o755
	filePermissions        fs.FileMode = 0o600
	defaultStorageLocation             = getDefaultApplicationFolder()
)

func getHomeDir() (directory string, err error) {
	dirname, err := os.UserHomeDir()
	if err != nil {
		jlogr.Logger.ILog.Error(err, "getting home directory")

		return "", fmt.Errorf("OS UserHomeDir: %w", err)
	}

	return dirname, nil
}

func getDefaultApplicationFolder() (storageLocation string) {
	applicationName := info.GetApplicationName()

	baseDir, err := getHomeDir()
	if err != nil {
		jlogr.Logger.ILog.Fatal(err, "Error getting home directory", "error", err)
	}

	storageLocation = path.Join(baseDir, "."+applicationName)

	return storageLocation
}

func MakeApplicationFolder(folder string) (location string, err error) {
	storageLocation := JoinPath(defaultStorageLocation, filepath.Clean(folder))

	err = os.MkdirAll(storageLocation, folderPermissons)
	if err != nil {
		jlogr.Logger.ILog.Error(err, "error creating storagelocation", "location", storageLocation)

		return "", err
	}

	return storageLocation, nil
}

func GetApplicationFile(leafDirectory string, fileName string, openFlagflag int) (*os.File, error) {
	location, err := MakeApplicationFolder(filepath.Clean(leafDirectory))
	if err != nil {
		return nil, err
	}

	if !fileExists(JoinPath(location, fileName)) {
		return os.Create(JoinPath(location, fileName))
	}
	return os.OpenFile(JoinPath(location, fileName), openFlagflag, filePermissions)
}

func fileExists(fileName string) (exists bool) {
	if _, err := os.Stat(fileName); err != nil {
		return false
	}

	return true
}

func JoinPath(basePath, leaf string) (fullPath string) {
	return path.Join(basePath, leaf)
}

func MakeSchemaDirectory(embedFiles embed.FS) (string, error) {
	schemaLocation := "schema"

	location, err := MakeApplicationFolder(schemaLocation)
	if err != nil {
		return "", err
	}

	err = fs.WalkDir(embedFiles, ".", func(filePath string, fileEntry fs.DirEntry, err error) error {
		if fileEntry.IsDir() {
			_, err := MakeApplicationFolder(JoinPath(schemaLocation, filePath))

			return err
		}

		bytes, err := fs.ReadFile(embedFiles, filePath)
		if err != nil {
			return err
		}

		schemaFile, err := GetApplicationFile(filepath.Dir(filePath), fileEntry.Name(), os.O_CREATE|os.O_RDWR|os.O_APPEND)
		if err != nil {
			return err
		}

		defer schemaFile.Close()
		log.Println(schemaFile.Name())

		_, err = schemaFile.Write(bytes)
		if err != nil {
			return err
		}

		return nil
	})

	return location, err
}
