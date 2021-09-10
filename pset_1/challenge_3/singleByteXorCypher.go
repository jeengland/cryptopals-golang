package challenge_3

import (
	"math"
	"strings"
	"unicode"
)

var charFreq = map[string]float64{
	"a": 8.167,
	"b": 1.492,
	"c": 2.782,
	"d": 4.253,
	"e": 12.702,
	"f": 2.228,
	"g": 2.015,
	"h": 6.094,
	"i": 6.966,
	"j": 0.153,
	"k": 0.772,
	"l": 4.025,
	"m": 2.406,
	"n": 6.749,
	"o": 7.507,
	"p": 1.929,
	"q": 0.095,
	"r": 5.987,
	"s": 6.327,
	"t": 9.056,
	"u": 2.758,
	"v": 0.978,
	"w": 2.360,
	"x": 0.150,
	"y": 1.974,
	"z": 0.074,
}

func SingleByteXorCypher(bytes []byte) (string, float64) {
	var lowest = 100.000
	var result = ""

	for i := 0; i < 256; i++ {
		key := byte(i)
		decoded := decode(bytes, key)
		variance := getVariance(decoded)
		if variance < lowest {
			result = decoded
			lowest = variance
		}
	}
	return result, lowest
}

func decode(encoded []byte, key byte) string {
	decoded := []byte{}
	for i := range encoded {
		decoded = append(decoded, encoded[i]^key)
	}
	return string(decoded)
}

func getVariance(testString string) float64 {
	testString = strings.ToLower(testString)
	variance := 0.000
	for c := range charFreq {
		actualFreq := float64(strings.Count(testString, c)) / float64(len(testString)) * float64(100)
		expectFreq := charFreq[c]
		variance += math.Abs(expectFreq - actualFreq)
	}

	for char := range testString {
		if !unicode.IsLetter(rune(testString[char])) && testString[char] != 32 {
			variance += 1.0
		}
	}

	return variance
}


