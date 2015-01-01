package diffsquares

import "math"

func SquareOfSums(n int) int {
  sum := 0
  for i := 0; i <= n; i++ {
    sum += i
  }
  return pow2(sum)
}

func SumOfSquares(n int) int {
  sum := 0
  for i := 0; i <= n; i++ {
    sum += pow2(i)
  }
  return sum
}

func Difference(n int) int {
  return SquareOfSums(n) - SumOfSquares(n)
}

func pow2(n int) int {
  return int(math.Pow(float64(n), float64(2)))
}