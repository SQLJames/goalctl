package main

import (
	"crypto/md5"  //#nosec G501 -- used to generate file checksums on build
	"crypto/sha1" //#nosec G505 -- used to generate file checksums on build
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"hash"
	"io"
	"io/fs"
	"os"
	"path"
	"path/filepath"

	"github.com/sqljames/goalctl/pkg/util/jlogr"
)

type Hash struct {
	name       string
	newHash    hash.Hash
	byteLength int
}

var (
	filePermissions fs.FileMode = 0o666
	/* #nosec G401 */
	hashes = []Hash{
		{
			name:       "md5",
			newHash:    md5.New(), //#nosec G401 -- used to generate file checksums on build
			byteLength: 16,
		},
		{
			name:       "sha1",
			newHash:    sha1.New(), //#nosec G401 -- used to generate file checksums on build
			byteLength: 20,
		},
		{
			name:       "sha256",
			newHash:    sha256.New(),
			byteLength: 32,
		},
	}
)

func (h *Hash) getFileChecksums(filePath string) (returnHashString string, err error) {
	sanitizedInput := filepath.Clean(filePath)

	file, err := os.Open(sanitizedInput)
	if err != nil {
		jlogr.Logger.ILog.Error(err, err.Error())

		return returnHashString, err
	}

	defer func(f *os.File) {
		if err := f.Close(); err != nil {
			jlogr.Logger.ILog.Error(err, err.Error())

		}
	}(file)

	if _, err := io.Copy(h.newHash, file); err != nil {
		jlogr.Logger.ILog.Error(err, err.Error())

		return returnHashString, err
	}

	hashInBytes := h.newHash.Sum(nil)[:h.byteLength]

	returnHashString = hex.EncodeToString(hashInBytes)

	return returnHashString, nil
}

func generateChecksum(filePath string) {
	sanitizedInput := filepath.Clean(filePath)

	files, err := os.ReadDir(sanitizedInput)
	if err != nil {
		jlogr.Logger.ILog.Fatal(err, err.Error())
	}

	for _, file := range files {
		workingFile := path.Join(sanitizedInput, file.Name())
		for _, hash := range hashes {
			checksum, err := hash.getFileChecksums(workingFile)
			if err != nil {
				jlogr.Logger.ILog.Fatal(err, err.Error())
			}

			writeChecksum(sanitizedInput, file.Name(), hash.name, checksum)
		}
	}
}

func writeChecksum(filePath, filename, checksumName, checksum string) {
	sanitizedInput := filepath.Clean(filePath)

	err := os.WriteFile(path.Join(sanitizedInput, fmt.Sprintf("%s.%s.txt", filename, checksumName)), []byte(checksum), filePermissions)
	if err != nil {
		jlogr.Logger.ILog.Fatal(err, err.Error())
	}
}
