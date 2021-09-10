package challenge_2

import "testing"

var input = "1c0111001f010100061a024b53535009181c"
var key = "686974207468652062756c6c277320657965"
var output = "746865206b696420646f6e277420706c6179"

func TestChallenge2(t *testing.T) {
	ans := FixedXor(input, key)
	if ans != output {
		t.Errorf("rcvd %s;\n want %s\n", ans, output)
	}
}
