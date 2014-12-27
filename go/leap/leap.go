package leap

import (
	"math"
)

func IsLeapYear(year int) bool {
	return IsDivisibleBy(year, 4) &&
		(!IsDivisibleBy(year, 100) || IsDivisibleBy(year, 400))
}

func IsDivisibleBy(year int, div int) bool {
	return math.Mod(float64(year), float64(div)) == 0
}
