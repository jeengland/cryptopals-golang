package challenge_4

import (
	"testing"
)

var inputFile = "./challenge_4_input.txt"

func TestSingleByteXorDetector(t *testing.T) {
	input := FileToStringSlice(inputFile)
	output := "Now that the party is jumping\n"
	ans, _ := singleByteXorDetector(input)
	if (ans != output) {
		t.Errorf("rcvd %s;\n want %s\n", ans, output)
	}
}

func TestFileToStringSlice(t *testing.T) {
	fileRef := "./test_input.txt"
	slice := FileToStringSlice(fileRef)

	if slice[0] != "test1" || slice[1] != "test2" || slice[2] != "test3" {
		t.Errorf("Slice not correctly parsed")
	}
}
