package radix

func Sort(data []uint64) {
	var position, n_position, r int
	sorted := make([]uint64, len(data))
	for i := 0; i < 8; i++ {
		C := [256]int{}
		for j := range data {
			r = int((data[j] >> (i * 8)) & 0b11111111)
			C[r]++
		}
		position = C[0]
		C[0] = 0
		for j := 1; j < 256; j++ {
			n_position = C[j]
			C[j] = C[j-1] + position
			position = n_position
		}
		for j := range data {
			r = int((data[j] >> (i * 8)) & 0b11111111)
			sorted[C[r]] = data[j]
			C[r]++
		}
		copy(data, sorted)
	}
}
