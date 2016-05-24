package hamming

import (
	"errors"
)

const testVersion = 4

var DifferentLengths = errors.New("strings must have same length")

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return -1, DifferentLengths
	}
	distance := 0
	for i := range a {
		if a[i] != b[i] {
			distance++
		}
	}
	return distance, nil
}
