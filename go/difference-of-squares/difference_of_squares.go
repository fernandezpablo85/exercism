package diffsquares

import "math"

func SquareOfSums(n int) int {
  sum := 0
  for i := 1; i <= n; i++ {
    sum += i
  }
  return pow2(sum)
}

func SumOfSquares(n int) int {
  sum := 0
  for i := 1; i <= n; i++ {
    sum += pow2(i)
  }
  return sum
}

func Difference(n int) int {
  return SquareOfSums(n) - SumOfSquares(n)
}

var mem = make(map[int]int)
func pow2(n int) int {
  if mem[n] > 0 {
    return mem[n]
  } else {
    pow := doPow(n)
    mem[n] = pow
    return pow
  }
}

func doPow(n int) int {
  return int(math.Pow(float64(n), float64(2)))
}