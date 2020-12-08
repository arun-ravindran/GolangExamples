package jsonEncDec

import (
	"bufio"
	"os"
	"testing"
)

func TestEncode(t *testing.T) {
	encode("input.txt", "intermediate.json", "output.txt")
	inF, err := os.Open("input.txt")
	checkError(err)
	defer inF.Close()
	outF, err := os.Open("output.txt")
	checkError(err)
	defer outF.Close()
	scanner := bufio.NewScanner(inF)
	inStr := ""
	for scanner.Scan() {
		inStr += scanner.Text()

	}
	scanner = bufio.NewScanner(outF)
	outStr := ""
	for scanner.Scan() {
		outStr += scanner.Text()

	}
	if inStr != outStr {
		t.Error("Input and output files do not match")
	}

}
