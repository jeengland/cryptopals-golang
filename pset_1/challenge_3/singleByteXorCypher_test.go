package challenge_3

import (
	"encoding/hex"
	"fmt"
	"testing"
)

var input = "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"

func TestChallenge2(t *testing.T) {
	bytes, err := hex.DecodeString(input)
	if err != nil {
		fmt.Println(err)
	}
	ans, _ := SingleByteXorCypher(bytes)
	output := "Cooking MC's like a pound of bacon"
	if ans !=  output {
		t.Errorf("rcvd %s;\n want %s\n", ans, output)
	}
}

func TestDecode(t *testing.T) {
	key := byte(80)
	output := "Kggcafo(EK/{(dacm(i(xg}fl(gn(jikgf"
	bytes, err := hex.DecodeString(input)
	if err != nil {
		fmt.Println(err)
	}
	ans := decode(bytes, key)
	if ans != output {
		t.Errorf("rcvd %s;\n want %s\n", ans, output)
	}
}

func TestGetVariance(t *testing.T) {
	i := "The quick fast fox jumps over the lazy dog"
	output := "54.6994"
	variance := getVariance(i)
	ans := fmt.Sprintf("%.4f", variance)
	if ans != output {
		t.Errorf("rcvd %s;\n want %s\n", ans, output)
	}
}
