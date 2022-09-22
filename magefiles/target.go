package main

import (
	"fmt"
	"path"
)

type target struct {
	GOOS      string
	GOARCH    string
	SourceDir string
}

func (t target) name() string {
	_, n := path.Split(gitRoot())
	name := fmt.Sprintf(
		"%s-%s-%s",
		n,
		t.GOOS,
		t.GOARCH,
	)

	if t.GOOS == "windows" {
		name += ".exe"
	}

	return name
}
