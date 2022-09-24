package main

import (
	"fmt"
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

var (
	goexe      = "go"
	binaryPath = path.Join(gitRoot(), "bin")
	tempPath   = path.Join(gitRoot(), "tmp")
	dirs       = []string{binaryPath, tempPath}

	targets = []target{
		{GOOS: "linux", GOARCH: "amd64"},
		{GOOS: "darwin", GOARCH: "amd64"},
		{GOOS: "windows", GOARCH: "amd64"},
	}
)

func init() {
	if exe := os.Getenv("GOEXE"); exe != "" {
		goexe = exe
	}
}

func flags() string {
	f := ldflags{}

	f["pkg/info.applicationName"] = filepath.Base(modulePath())
	f["pkg/version.buildDate"] = time.Now().Format(time.RFC3339)
	f["pkg/version.release"] = getRelease()
	f["pkg/version.buildTag"] = gitTag()
	f["pkg/version.commitHash"] = gitCommitHash()
	f["pkg/version.buildVersion"] = getVersion()
	f["pkg/version.buildBranch"] = gitBranch()

	return f.String()
}

func ensureDirs() error {
	fmt.Println("--> Ensuring output directories")

	for _, dir := range dirs {
		if !fileExists(dir) {
			fmt.Printf("    creating '%s'\n", dir)
			if err := os.MkdirAll(dir, 0750); err != nil {
				return err
			}
		}
	}
	return nil
}

// Clean up after yourself
func Clean() {
	fmt.Println("--> Cleaning output directories")

	for _, dir := range dirs {
		fmt.Printf("    removing '%s'\n", dir)
		err := os.RemoveAll(dir)
		if err != nil {
			log.Fatal("error running clean command", err.Error())
		}
	}
}

// Vendor dependencies with go modules
func Vendor() {
	fmt.Println("--> Updating dependencies")
	err := sh.Run(goexe, "mod", "tidy")
	if err != nil {
		log.Fatal("error running Vendor command", err.Error())
	}
}

// Build the application for local running
func Build() error {
	mg.SerialDeps(Vendor, ensureDirs)

	sourcePath := gitRoot()
	if err := sh.Run(goexe, "build", "-o", binaryPath, "-ldflags="+flags(), sourcePath); err != nil {
		return err
	}

	return nil

}

// Release the application for all defined targets
func Release() error {
	mg.SerialDeps(Vendor, ensureDirs)

	cgoEnabled := os.Getenv("CGO_ENABLED") == "1"

	var wg sync.WaitGroup
	wg.Add(len(targets))

	fmt.Printf("--> Building '%s' for release\n", gitRoot())
	for _, t := range targets {
		t.SourceDir = gitRoot()
		go func(t target) {
			defer wg.Done()

			env := map[string]string{
				"GOOS":   t.GOOS,
				"GOARCH": t.GOARCH,
			}

			if cgoEnabled && runtime.GOOS != env["GOOS"] {
				fmt.Printf("      CGO is enabled, skipping compilation of %s for %s\n", t.name(), env["GOOS"])
				return
			}
			fmt.Printf("      Building %s\n", t.name())

			err := sh.RunWith(env, goexe, "build", "-o", path.Join(binaryPath, t.name()), "-ldflags="+flags(), t.SourceDir)
			//err := sh.RunWith(env, goexe, "build", "-o", path.Join(binaryPath, t.Name()), t.SourceDir)
			if err != nil {
				fmt.Printf("compilation failed: %s\n", err.Error())
				return
			}
		}(t)
	}

	wg.Wait()

	return nil
}

// StaticSecurity runs various static checkers to ensure you minimize security holes
func Scan() (err error) {
	fmt.Println("--> Scanning code")
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

// Test the codebase
func Test() error {
	mg.SerialDeps(Vendor, ensureDirs)

	fmt.Println("--> Testing codebase")
	results, err := sh.Output(goexe, "test", "-cover", "-e", "internal", "-e", "cache", "./...")
	fmt.Print("    ")
	fmt.Println(strings.Replace(results, "\n", "\n    ", -1))

	return err
}
