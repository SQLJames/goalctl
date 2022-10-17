package main

import (
	"os"
	"os/exec"
	"strings"

	"github.com/pkg/errors"
	"github.com/sqljames/goalctl/pkg/util/jlogr"
)

func gitTag() string {
	var err error

	tagOrBranch, ok := os.LookupEnv("CI_COMMIT_REF_NAME")
	if !ok {
		tagOrBranch, err = gitShellOut("describe", "--tags", "--always")
		if err != nil {
			jlogr.Logger.ILog.Error(err, err.Error())
			var exitError exec.ExitError

			ok := errors.As(err, &exitError)
			if ok && exitError.Exited() {
				// probably no git tag.
				return devRelease
			}

			return devRelease
		}
	}

	return strings.TrimSuffix(tagOrBranch, "\n")
}

func gitCommitHash() string {
	var err error

	hash, ok := os.LookupEnv("CI_COMMIT_SHA")
	if !ok {
		hash, err = gitShellOut("rev-parse", "HEAD")
		if err != nil {
			jlogr.Logger.ILog.Error(err, err.Error())

			return ""
		}
	}

	return strings.TrimSpace(hash)
}

func gitBranch() string {
	branch, err := gitShellOut("rev-parse", "--abbrev-ref", "HEAD")
	if err != nil {
		jlogr.Logger.ILog.Error(err, err.Error())

		return err.Error()
	}

	return strings.TrimSpace(branch)
}

// gitRoot returns the root directory from the git repo.
func gitRoot() string {
	directory, err := gitShellOut("rev-parse", "--show-toplevel")
	if err != nil {
		jlogr.Logger.ILog.Error(err, err.Error())

		return ""
	}

	return strings.TrimSpace(directory)
}

func gitShellOut(args ...string) (string, error) {
	c := exec.Command("git", args...)
	c.Env = os.Environ()
	c.Stderr = os.Stderr

	bytes, err := c.Output()
	if err != nil {
		jlogr.Logger.ILog.Error(err, err.Error())

		return "", errors.Wrapf(err, `failed to run %v %q`, "git", args)
	}

	return string(bytes), nil
}
