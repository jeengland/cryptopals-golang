package challenge_6

import (
	"cryptopals/pset_1/challenge_3"
	"fmt"
)

func BreakRepeatingKeyXor(input []byte) string {
	bestKeysizes := FindBestKeySizes(input)
	testCases := SplitByKeySize(input, bestKeysizes)
	bestVariance := 100.00
	bestCaseResults := []string{}
	for i := range testCases {
		testCase := testCases[i]
		keysize := float64(len(testCase))
		testCaseResults := []string{}
		testCaseTotalVariance := 0.00
		for j := range testCase {
			key, variance := challenge_3.SingleByteXorCypher(testCase[j])
			testCaseResults = append(testCaseResults, key)
			testCaseTotalVariance += variance
		}
		testCaseAverageVariance := testCaseTotalVariance / keysize
		if testCaseAverageVariance < bestVariance {
			bestVariance = testCaseAverageVariance
			bestCaseResults = testCaseResults
		}
	}
	return Interpolate(bestCaseResults)
}

func Interpolate(xs []string) string {
	maxLength := 0
	result := ""
	for i := range xs {
		if len(xs[i]) > maxLength {
			maxLength = len(xs[i])
		}
	}
	for j := 0; j < maxLength; j++ {
		for k := range xs {
			if len(xs[k]) <= j{
				continue
			}
			result += string(xs[k][j])
		}
	}
	return result
}

func SplitByKeySize(bytes []byte, keysizes []int) [][][]byte {
	results := [][][]byte{}
	for i := range keysizes {
		max := keysizes[i]
		options := make([][]byte, max)
		sliceIndex := 0
		for j := 0; j < len(bytes); j++ {
			options[sliceIndex] = append(options[sliceIndex], bytes[j])
			sliceIndex++
			if sliceIndex == max {
				sliceIndex = 0
			}
		}
		results = append(results, options)
	}
	return results
}

func FindBestKeySizes(bytes []byte) []int {
	keysizeMin := 2
	keysizeMax := 40
	first, second, third := 0, 0, 0
	firstDistance, secondDistance, thirdDistance := 100.000, 100.000, 100.000
	for i := keysizeMin; i <keysizeMax; i++ {
		normalizedDistance := 0.000
		for j := 0; j < 7; j ++ {
			slice1 := bytes[i * j : i * (j + 1)]
			slice2 := bytes[i * (j + 1) : i * (j + 1) + i * (j + 2)]
			distance := GetEditDistance(slice1, slice2)
			normalizedDistance += float64(distance) / float64(i)
		}

		averageDistance := normalizedDistance / 7

		if averageDistance < firstDistance {
			thirdDistance, secondDistance, firstDistance = secondDistance, firstDistance, averageDistance
			third, second, first = second, first, i
		} else if averageDistance< secondDistance {
			thirdDistance, secondDistance = secondDistance, averageDistance
			third, second = second, i
		} else if averageDistance < thirdDistance {
			thirdDistance = averageDistance
			third = i
		}
	}
	return []int{first, second, third}
}

func GetEditDistance(a, b []byte) int {
	editDistance := 0
	aBits := BytesToBinary(a)
	bBits := BytesToBinary(b)
	for i := range aBits {
		if aBits[i] != bBits[i] {
			editDistance += 1
		}
	}
	return editDistance
}

func BytesToBinary(b []byte) string {
	var binary string
	for _, c := range b {
		binary = fmt.Sprintf("%s%08b", binary, c)
	}
	return binary
}