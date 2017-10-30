// Copyright 2016 Sergey Berezin, sergey@sergeyberezin.com

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
