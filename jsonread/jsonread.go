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

// jsonread tool takes a file and a key, reads the file as a flat JSON format,
// and prints out the value of the key. If key is not present, prints the default value.

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

var usage = `Usage:
  jsonread <file.json> <key> <default>

  Prints the value of the <key> in a JSON file, assumed to be a top-level map.
  If the key is not present, prints <default>.  
`

func main() {
	if len(os.Args) != 4 {
		log.Fatal(fmt.Sprintf("ERROR: expected 3 arguments, %d given.\n%s", len(os.Args)-1, usage))
	}

	fileName := os.Args[1]
	key := os.Args[2]
	defaultValue := os.Args[3]

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(fmt.Sprintf("ERROR: cannot open file '%s': %s", fileName, err))
	}
	dec := json.NewDecoder(file)
	var topMap map[string]interface{}
	if err = dec.Decode(&topMap); err != nil {
		log.Fatal(fmt.Sprintf("ERROR: failed to decode JSON object in %s: %s", fileName, err))
	}
	if value, ok := topMap[key]; ok {
		fmt.Println(value)
	} else {
		fmt.Println(defaultValue)
	}
}
