package raindrops

import "fmt"

func Convert(n int) string {
	conversion := ""
  divisibleBy := func(m int) bool {
    return n % m == 0
  }

	if divisibleBy(3) {
		conversion += "Pling"
	}

	if divisibleBy(5) {
		conversion += "Plang"
	}

	if divisibleBy(7) {
		conversion += "Plong"
	}

	if len(conversion) > 0 {
		return fmt.Sprintf("%s", conversion)
	} else {
		return fmt.Sprintf("%d", n)
	}
}

func ConvertInlined(n int) string {
  conversion := ""

  if n % 3 == 0 {
    conversion += "Pling"
  }

  if n % 5 == 0 {
    conversion += "Plang"
  }

  if n % 7 == 0 {
    conversion += "Plong"
  }

  if len(conversion) > 0 {
    return fmt.Sprintf("%s", conversion)
  } else {
    return fmt.Sprintf("%d", n)
  }
}
