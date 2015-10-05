package slice

func Frist(length int, value string) string {
	first, _ := First(length, value)
	return first
}

func First(length int, value string) (string, bool) {
	if len(value) < length {
		return "", false
	} else {
		return value[:length], true
	}
}

func All(length int, value string) []string {
	return allImpl(length, value, []string{})
}

func allImpl(length int, value string, all []string) []string {
	if length > len(value) {
		return all
	} else {
		first := Frist(length, value)
		rest := value[1:]
		updated := append(all, first)
		return allImpl(length, rest, updated)
	}
}
