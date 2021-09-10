package challenge_7

import (
	"crypto/aes"
	"fmt"
)

func DecryptAESinECB(bytes, key []byte) string {
	aes, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println(err)
	}
	decrypted := make([]byte, len(bytes))
	bs := aes.BlockSize()

	d := decrypted
	for len(bytes) > 0 {
		aes.Decrypt(d, bytes[:bs])
		d = d[bs:]
		bytes = bytes[bs:]
	}

	return string(decrypted)
}

