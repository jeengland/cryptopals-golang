package challenge_7

import (
	"cryptopals/pset_1/challenge_4"
	"encoding/base64"
	"fmt"
	"strings"
	"testing"
)

var inputFile = "./challenge_7_input.txt"

func TestDecryptAESinECB(t *testing.T) {
	input, err := base64.StdEncoding.DecodeString(strings.Join(challenge_4.FileToStringSlice(inputFile), ""))
	if err != nil {
		fmt.Println(err)
	}
	key := []byte("YELLOW SUBMARINE")
	ans := DecryptAESinECB(input, key)

	if !strings.Contains(ans, "Play that funky music") {
		t.Errorf("Incorrect return")
	}
}
