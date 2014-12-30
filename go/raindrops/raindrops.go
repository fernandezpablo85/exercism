package raindrops

import "fmt"

func Convert(n int) string {
	pling := ""
	plang := ""
	plong := ""
	factors := primeFactors(n)

	if contains(3, factors) {
		pling = "Pling"
	}

	if contains(5, factors) {
		plang = "Plang"
	}

	if contains(7, factors) {
		plong = "Plong"
	}

	anyFactor := len(pling) > 0 || len(plang) > 0 || len(plong) > 0

	if anyFactor {
		return fmt.Sprintf("%s%s%s", pling, plang, plong)
	} else {
		return fmt.Sprintf("%d", n)
	}
}

func contains(n int, ns []int) bool {
	for _, i := range ns {
		if i == n {
			return true
		}
	}
	return false
}

func primeFactors(n int) []int {
	var factors []int
	for i := 2; i <= n; i++ {
		if n%i == 0 && IsPrime(i) {
			factors = append(factors, i)
		}
	}
	return factors
}

func IsPrime(n int) bool {
	prime := true
	for i := (n - 1); i > 1 && prime; i-- {
		prime = prime && (n%i != 0)
	}
	return prime
}
