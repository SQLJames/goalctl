package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"time"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

const (
	devRelease        string      = "dev"
	folderPermissions fs.FileMode = 0o750
)

var (
	goexe      = getGoExe()
	binaryPath = path.Join(gitRoot(), "bin")
	tempPath   = path.Join(gitRoot(), "tmp")
	dirs       = []string{binaryPath, tempPath}

	targets = []target{
		{GOOS: "linux", GOARCH: "amd64"},
		{GOOS: "darwin", GOARCH: "amd64"},
		{GOOS: "windows", GOARCH: "amd64"},
	}
)

func getGoExe() (goexe string) {
	goexe = os.Getenv("GOEXE")
	if goexe == "" {
		goexe = "go"
	}

	return goexe
}

func flags() string {
	pkgFlags := ldflags{}

	pkgFlags["pkg/info.applicationName"] = filepath.Base(modulePath())
	pkgFlags["pkg/version.buildDate"] = time.Now().Format(time.RFC3339)
	pkgFlags["pkg/version.release"] = getRelease()
	pkgFlags["pkg/version.buildTag"] = gitTag()
	pkgFlags["pkg/version.commitHash"] = gitCommitHash()
	pkgFlags["pkg/version.buildVersion"] = getVersion()
	pkgFlags["pkg/version.buildBranch"] = gitBranch()

	return pkgFlags.String()
}

func ensureDirs() error {
	log.Println("--> Ensuring output directories")

	for _, dir := range dirs {
		if !fileExists(dir) {
			log.Printf("    creating '%s'\n", dir)

			if err := os.MkdirAll(dir, folderPermissions); err != nil {
				return fmt.Errorf("OS MkdirAll: %w", err)
			}
		}
	}

	return nil
}

// Clean up after yourself.
func Clean() {
	log.Println("--> Cleaning output directories")

	for _, dir := range dirs {
		log.Printf("    removing '%s'\n", dir)

		err := os.RemoveAll(dir)
		if err != nil {
			log.Fatal("error running clean command", err.Error())
		}
	}
}

// Vendor dependencies with go modules.
func Vendor() {
	log.Println("--> Updating dependencies")

	err := sh.Run(goexe, "mod", "tidy")
	if err != nil {
		log.Fatal("error running Vendor command", err.Error())
	}
}

// Build the application for local running.
func Build() error {
	mg.SerialDeps(Vendor, ensureDirs)

	sourcePath := gitRoot()
	if err := sh.Run(goexe, "build", "-o", binaryPath, "-ldflags="+flags(), sourcePath); err != nil {
		return fmt.Errorf("sh Run: %w", err)
	}

	return nil
}

// Release the application for all defined targets.
func Release() {
	var waitGroup sync.WaitGroup

	mg.SerialDeps(Vendor, ensureDirs, Clean)
	waitGroup.Add(len(targets))

	cgoEnabled := os.Getenv("CGO_ENABLED") == "1"

	log.Printf("--> Building '%s' for release\n", gitRoot())

	for _, buildTarget := range targets {
		buildTarget.SourceDir = gitRoot()
		go func(buildTarget target) {
			defer waitGroup.Done()

			env := map[string]string{
				"GOOS":   buildTarget.GOOS,
				"GOARCH": buildTarget.GOARCH,
			}

			if cgoEnabled && runtime.GOOS != env["GOOS"] {
				log.Printf("      CGO is enabled, skipping compilation of %s for %s\n", buildTarget.name(), env["GOOS"])

				return
			}

			log.Printf("      Building %s\n", buildTarget.name())

			err := sh.RunWith(env, goexe, "build", "-o", path.Join(binaryPath, buildTarget.name()), "-ldflags="+flags(), buildTarget.SourceDir)
			if err != nil {
				log.Printf("compilation failed: %s\n", err.Error())

				return
			}
		}(buildTarget)
	}

	waitGroup.Wait()
	generateChecksum(binaryPath)
}

// Scan runs various static checkers to ensure you minimize security holes and have good formatting.
func Scan() (err error) {
	log.Println("--> Scanning code")

	err = confirmScanners()
	if err != nil {
		return err
	}

	err = runStaticScanners()
	if err != nil {
		return err
	}

	return nil
}

// Test the codebase.
func Test() error {
	mg.SerialDeps(Vendor, ensureDirs)

	log.Println("--> Testing codebase")

	results, err := sh.Output(goexe, "test", "-cover", "-e", "internal", "-e", "cache", "./...")

	log.Println(strings.ReplaceAll(results, "\n", "\n    "))

	return fmt.Errorf("sh Output: %w", err)
}
