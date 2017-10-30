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

// gocovcheck tool parses a coverprofile file format dumped by
//    `go test -coverprofile <file>`
// computes the overall coverage percentage, and compares it with the
// requested minimum coverage. The command fails (return exit != 0) when
// coverage is less than the required minimum.

package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/sergey-a-berezin/gocovcheck/coverage"
)

var usage = "Usage:\n  gocovcheck <coverprofile-file-name> <min-coverage-percent>"

func main() {
	if len(os.Args) != 3 {
		log.Fatal(fmt.Sprintf("ERROR: expected 2 arguments, %d given.\n%s", len(os.Args)-1, usage))
	}

	fileName := os.Args[1]
	minCoverage, err := strconv.ParseFloat(os.Args[2], 64)
	if err != nil || 0.0 > minCoverage || minCoverage > 100.0 {
		log.Fatal(fmt.Sprintf("ERROR: min-coverage-percent must be a float between 0..100. Got %v",
			minCoverage))
	}

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(fmt.Sprintf("ERROR: cannot open file '%s': %s", fileName, err))
	}
	total, covered, err := coverage.Extract(file)
	if err != nil {
		log.Fatal(err)
	}
	percentage := coverage.Percentage(total, covered)
	fmt.Printf("Covered %d of %d statements - %.1f%%, expected %.1f%%",
		covered, total, percentage, minCoverage)
	if percentage < minCoverage {
		fmt.Println(" - LOW COVERAGE")
		os.Exit(1)
	} else {
		fmt.Println("")
	}
}
