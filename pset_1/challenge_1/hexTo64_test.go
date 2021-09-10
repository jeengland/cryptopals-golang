package challenge_1

import "testing"

var input = "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
var output = "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"

func TestChallenge1(t *testing.T) {
	ans := HexToBase64(input)
	if ans != output {
		t.Errorf("rcvd %s;\n want %s\n", ans, output)
	}
}
