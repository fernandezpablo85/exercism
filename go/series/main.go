package slice

func Frist(length int, value string) string {
	return value[:length]
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
