// Read in an input text file into slice of KeyVals, encode to JSON, write to intermediate file
// Read intermediate file, decode JSON, and write to output file
// Output file should be identical to input file
package jsonEncDec

import (
	"bufio"
	"encoding/json"
	"io"
	"log"
	"os"
)

type KeyVal struct {
	Key string
	Val string
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func encode(inFileName, interFileName, outFileName string) {
	// Open file
	inF, err := os.Open(inFileName)
	checkError(err)
	defer inF.Close()

	// Read file with scanner into slice of *KeyVal
	var kvs []*KeyVal
	scanner := bufio.NewScanner(inF)
	for scanner.Scan() {
		kvs = append(kvs, &KeyVal{Key: scanner.Text(), Val: ""})
	}

	// Encode to JSON and write to intermediate file
	intF, err := os.OpenFile(interFileName, os.O_CREATE|os.O_WRONLY, 0644)
	checkError(err)
	defer intF.Close()

	enc := json.NewEncoder(intF)
	for _, kv := range kvs {
		err = enc.Encode(&kv)
	}

	// Read intermediate json file and decode JSON to slice of *KeyVal
	jsonF, err := os.Open(interFileName)
	checkError(err)
	defer jsonF.Close()

	dec := json.NewDecoder(jsonF)
	var kv KeyVal
	kvs = nil
	for {
		kv := kv

		if err := dec.Decode(&kv); err == io.EOF {
			break
		} else {
			checkError(err)
		}
		kvs = append(kvs, &kv)

	}

	// Write slice to file
	outF, err := os.OpenFile(outFileName, os.O_CREATE|os.O_WRONLY, 0644)
	checkError(err)
	defer outF.Close()

	w := bufio.NewWriter(outF)
	for _, kv := range kvs {
		_, err := w.WriteString(kv.Key + "\n")
		checkError(err)

	}
	w.Flush()

}
