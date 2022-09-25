package main

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/magefile/mage/sh"
)

type scanner struct {
	command string
	source  string
	version string
	runArgs []string
}

var scanners = []scanner{
	{command: "golangci-lint", source: "github.com/golangci/golangci-lint/cmd/golangci-lint", version: "latest", runArgs: []string{"run"}},
	{command: "govulncheck", source: "golang.org/x/vuln/cmd/govulncheck", version: "latest", runArgs: []string{"./..."}},
}

func (s *scanner) getInstallURL() (installURL string) {
	return fmt.Sprintf("%s@%s", s.source, s.version)
}
func runStaticScanners() (err error) {
	for _, scanner := range scanners {
		log.Printf("--> Running Scanner: %s\n", scanner.command)
		if err := sh.RunV(scanner.command, scanner.runArgs...); err != nil {
			return err
		}
	}
	return nil
}
func confirmScanners() (err error) {
	for _, scanner := range scanners {
		err = installIfMissing(scanner.command, scanner.getInstallURL())
		if err != nil {
			return err
		}
	}
	return nil
}

// installIfMissing checks for existence then installs a file if it's not there
func installIfMissing(executableName, installURL string) (err error) {
	log.Printf("--> Checking if Scanner Exists: %s\n", executableName)
	_, missing := exec.LookPath(executableName)
	if missing != nil {
		log.Printf("--> Scanner Missing Installing Scanner: %s\n", executableName)
		err := sh.Run("go", "install", installURL)
		if err != nil {
			return err
		}
	}
	return nil
}
