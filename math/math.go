package math

import "math"

func NthPrime(n int) int {
	if n < 1 {
		panic("n is less that 1")
	}
	smallnum := [...]int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103, 107, 109, 113, 127, 131, 137, 139, 149, 151, 157, 163, 167, 173, 179, 181, 191, 193, 197, 199}
	if n < len(smallnum) {
		return smallnum[n-1]
	}
	lnn := math.Log(float64(n))
	upper := int(float64(n)*(lnn+math.Log(lnn))) + 20
	arr := make([]bool, upper/2)
	for i := 3; i*i <= upper; i += 2 {
		if !arr[(i-3)/2] {
			for j := i * i; j <= upper; j += 2 * i {
				arr[(j-3)/2] = true
			}
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
