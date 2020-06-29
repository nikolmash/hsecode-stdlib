package math

import "math"

func NthPrime(n int) int {
	if n < 1 {
		panic("n is less that 1")
	}
	smallnum := []int{2, 3, 5, 7, 11, 13}
	if n < 6 {
		return smallnum[int(n)-1]
	}
	upper := int(float64(n)*math.Log(float64(n)) + float64(n)*math.Log(math.Log(float64(n))))
	arr := make([]bool, upper+10)

	for i := 2; i*i < upper; i++ {
		for j := i * i; j < upper; j += i {
			arr[j] = true
		}
	}
	res := 0
	for ind, val := range arr {
		if !val && ind > 1 {
			res += 1
		}
		if res == n {
			return ind
		}
	}
	panic("something wrong")
	return 0
}
