package main

import (
	"fmt"
	"log"

	"github.com/magefile/mage/sh"
	"github.com/sqljames/goalctl/pkg/util/jlogr"
)

type scanner struct {
	command string
	source  string
	version string
	runArgs []string
}

var scanners = []scanner{
	{command: "govulncheck", source: "golang.org/x/vuln/cmd/govulncheck", version: "latest", runArgs: []string{"./..."}},
	{command: "gosec", source: "github.com/securego/gosec/v2/cmd/gosec", version: "latest", runArgs: []string{"./..."}},
	{command: "golangci-lint", source: "github.com/golangci/golangci-lint/cmd/golangci-lint", version: "latest", runArgs: []string{"run"}},
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
		err = install(scanner.command, scanner.getInstallURL())
		if err != nil {
			jlogr.Logger.ILog.Error(err, err.Error())

			return err
		}
	}

	return nil
}

// installIfMissing checks for existence then installs a file if it's not there.
func install(executableName, installURL string) (err error) {
	log.Printf("--> Checking if Scanner Exists: %s\n", executableName)

	log.Printf("--> Scanner Missing Installing Scanner: %s\n", executableName)

	err = sh.Run("go", "install", installURL)
	if err != nil {
		jlogr.Logger.ILog.Error(err, err.Error())

		return err
	}

	return nil
}
