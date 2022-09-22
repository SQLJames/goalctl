package main

import (
	"os"
)

func getVersion() string {
	defaultValue := "0.0.0"
	v := "VERSION"
	value, set := os.LookupEnv(v)
	if set {
		return value
	}
	return defaultValue
}

func getRelease() string {
	defaultValue := "dev"
	v := "RELEASE"
	value, set := os.LookupEnv(v)
	if set {
		return value
	}
	return defaultValue
}
