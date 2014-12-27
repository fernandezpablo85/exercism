package leap

func IsLeapYear(year int) bool {
	switch {
	case IsDivisibleBy(year, 400):
		return true
	case IsDivisibleBy(year, 100):
		return false
	case IsDivisibleBy(year, 4):
		return true
	default:
		return false
	}
}

func IsDivisibleBy(year int, div int) bool {
	return (year % div) == 0
}
