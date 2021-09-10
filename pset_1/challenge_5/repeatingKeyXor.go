package challenge_5

import (
	"encoding/hex"
)

func RepeatingKeyXor(input, key string) string {
	output := []byte{}
	bytes := []byte(input)
	keyIndex := 0
	keyMax := len(key)
	for i := 0; i < len(bytes); i++ {
		output = append(output, bytes[i] ^ key[keyIndex])
		keyIndex++
		if keyIndex == keyMax {
			keyIndex = 0
		}
	}
	return hex.EncodeToString(output)
}