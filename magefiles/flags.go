package main

import (
	"fmt"
	"strings"
)

// LDFlags wraps the generation of LDFLags meant to be passed to the compiler.
type ldflags map[string]string

// Build the LDFlags for the given package.
func (ldf ldflags) Build(packageName string) string {
	var builder strings.Builder

	for k, v := range ldf {
		fmt.Fprintf(&builder, `-X "%s/%s=%s" `, packageName, k, v)
	}

	return builder.String()
}

// String returns the LDFlags for the detected ModulePath.
func (ldf ldflags) String() string {
	return ldf.Build(modulePath())
}
