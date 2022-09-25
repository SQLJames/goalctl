package main

import (
	"os"
	"os/exec"
	"strings"

	"github.com/pkg/errors"
)

func gitTag() string {
	var err error
	tagOrBranch, ok := os.LookupEnv("CI_COMMIT_REF_NAME")
	if !ok {
		tagOrBranch, err = shellout("git", "describe", "--tags", "--always")
		if err != nil {
			ee, ok := errors.Cause(err).(*exec.ExitError)
			if ok && ee.Exited() {
				// probably no git tag
				return "dev"
			}
			return "dev"
		}
	}

	return strings.TrimSuffix(tagOrBranch, "\n")
}

func gitCommitHash() string {
	var err error

	hash, ok := os.LookupEnv("CI_COMMIT_SHA")
	if !ok {
		hash, err = shellout("git", "rev-parse", "HEAD")
		if err != nil {
			return ""
		}
	}

	return strings.TrimSpace(hash)
}

func gitBranch() string {
	branch, err := shellout("git", "rev-parse", "--abbrev-ref", "HEAD")
	if err != nil {
		return err.Error()
	}
	branch = strings.TrimSpace(branch)
	return branch
}

// gitRoot returns the root directory from the git repo
func gitRoot() string {
	directory, err := shellout("git", "rev-parse", "--show-toplevel")
	if err != nil {
		return ""
	}
	directory = strings.TrimSpace(directory)
	return directory
}
