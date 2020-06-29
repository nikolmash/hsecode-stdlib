package math

import "math"

func NthPrime(n int) int {
	if n < 1 {
		panic("n is less that 1")
	}
	smallnum := [5]int{2, 3, 5, 7, 11}
	if n < 6 {
		return smallnum[n-1]
	}
	lnn := math.Log(float64(n))
	upper := int(float64(n)*(lnn+math.Log(lnn))) + 10
	arr := make([]bool, upper/2)
	for i := 3; i*i <= upper; i += 2 {
		for j := i * i; j <= upper; j += 2 * i {
			arr[(j-3)/2] = true
		}
	}
	res := 1
	for ind, val := range arr {
		if !val {
			res += 1
		}
		if res == n {
			return 2*ind + 3
		}
	}
	panic("something wrong")
	return 0
}
