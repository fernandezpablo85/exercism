package raindrops

import "fmt"

func Convert(n int) string {
	conversion := ""

	if n%3 == 0 {
		conversion += "Pling"
	}

	if n%5 == 0 {
		conversion += "Plang"
	}

	if n%7 == 0 {
		conversion += "Plong"
	}

	if len(conversion) > 0 {
		return fmt.Sprintf("%s", conversion)
	} else {
		return fmt.Sprintf("%d", n)
	}
}
