package challenge_4

import (
	"cryptopals/pset_1/challenge_3"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"strings"
)

func singleByteXorDetector(candidates []string) (string, int) {
	solution := ""
	lowest := 100.0
	var index int
	for i := range candidates {
		bytes, err := hex.DecodeString(candidates[i])
		if err != nil {
			fmt.Println(err)
		}
		res, vrc := challenge_3.SingleByteXorCypher(bytes)
		if vrc < lowest {
			solution = res
			index = i
			lowest = vrc
		}
	}
	return solution, index
}

func FileToStringSlice(fileRef string) []string {
	bytes, err := ioutil.ReadFile(fileRef)

	if err != nil {
		fmt.Println(err)
	}

	slice := strings.Split(string(bytes), "\n")

	return slice
}