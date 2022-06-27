package utils

import (
	"math"
)

func CalculateEntropy(dataBytes []byte, bytesRead int) (entropy float64, err error) {

	byteCounts := make([]int, 256)
	for i := 0; i < bytesRead; i++ {
		byteCounts[int(dataBytes[i])]++
	}

	for i := 0; i < 256; i++ {
		px := float64(byteCounts[i]) / float64(bytesRead)
		if px > 0 {
			entropy += -px * math.Log2(px)
		}
	}

	// Returns rounded to nearest two decimals.
	return math.Round(entropy*100) / 100, nil
}
