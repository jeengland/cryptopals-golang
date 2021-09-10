package challenge_2

import (
	"encoding/hex"
	"fmt"
)

func FixedXor(input, key string) string {
	iBytes, err := hex.DecodeString(input)
	if err != nil {
		fmt.Println(err)
	}
	kBytes, err := hex.DecodeString(key)
	if err != nil {
		fmt.Println(err)
	}

	output := []byte{}
	for i := range iBytes {
		output = append(output, iBytes[i]^kBytes[i])
	}
	return hex.EncodeToString(output)
}
