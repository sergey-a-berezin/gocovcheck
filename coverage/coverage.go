// Copyright 2016 sergey@sergeyberezin.com

// Package coverage parses the coverprofile file format and computes the total
// statement coverage.
package coverage

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// parseLine parses a single entry (line) of the coverprofile format, and
// extracts the total and covered number of statements represented by this
// entry.
func parseLine(line string) (total, covered uint64, err error) {
	parts := strings.Split(line, " ")
	switch len(parts) {
	case 2:
		if line != "mode: set" {
			err = fmt.Errorf("Failed to parse line: unsupported mode: %s", line)
		}
	case 3:
		total, err = strconv.ParseUint(parts[1], 10, 64)
		if err != nil {
			return
		}
		covered, err = strconv.ParseUint(parts[2], 10, 64)
		if covered != 0 { // mode: set semantics
			covered = total
		}
	default:
		err = fmt.Errorf("Failed to parse line: unrecognized format: %s", line)
	}
	return
}

// Extract parses the coverprofile input and extracts the overall coverage
// information.
func Extract(r io.Reader) (total, covered uint64, err error) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		if e := scanner.Err(); e != nil {
			err = fmt.Errorf("reading file '%s': %s", os.Args[1], err)
			return
		}
		totalBlock, coveredBlock, e := parseLine(line)
		if e != nil {
			err = e
			return
		}
		total += totalBlock
		covered += coveredBlock
	}
	return
}

// Percentage takes the extracted coverage (see Extract()) and computes the
// percentage coverage.  When no code is present (total == 0), the default
// coverage is 0.0.
func Percentage(total, covered uint64) (percentage float64) {
	if total > 0 {
		percentage = 100.0 * float64(covered) / float64(total)
	}
	return
}
