package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

// ModulePath returns the module path from the gomod file text.
func modulePath() string {
	file, err := os.Open("go.mod")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	var lines []string
	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return parseFileLines(lines)
}

func parseFileLines(lines []string) (modulePath string) {
	split := strings.Split(lines[0], " ")
	return split[1]
}
