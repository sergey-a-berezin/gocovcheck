// Copyright 2019 Sergey Berezin

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// jsonread tool takes a file and a key, reads the file as a flat JSON format,
// and prints out the value of the key. If key is not present, prints the default value.

package main

// This app computes an app version based on the git number (the number of
// commits since the beginning) and the short git hash.

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os/exec"
)

// Version creates a version string based on `git log` output.
func Version(r io.Reader) string {
	scanner := bufio.NewScanner(r)
	latestHash := ""
	count := 0
	if scanner.Scan() {
		latestHash = scanner.Text()
		count++
	}
	for scanner.Scan() {
		count++
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("Error running `git log`: %s", err)
	}
	return fmt.Sprintf("%05d-%s", count, latestHash)
}

func main() {
	cmd := exec.Command("git", "log", `--format=%h`)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(Version(stdout))
}
