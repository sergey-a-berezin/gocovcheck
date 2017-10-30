// Copyright 2017 Sergey Berezin

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
