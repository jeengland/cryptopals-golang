package challenge_1

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func HexToBase64(hexString string) string {
	var bytes, err = hex.DecodeString(hexString)
	if err != nil {
		fmt.Println(err)
	}
	return base64.StdEncoding.EncodeToString(bytes)
}
