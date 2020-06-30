package radix

func Sort(data []uint64) {
	var position int
	var r uint8
	sorted := make([]uint64, len(data))
	copy_data := make([]uint64, len(data))
	copy(copy_data, data)
	for i := 0; i < 8; i++ {
		C := [256]int{}
		for j := range copy_data {
			r = uint8(copy_data[j])
			C[r]++
		}
		position = C[0]
		for j := 1; j < 255; j++ {
			position = C[j]
			C[j] = C[j-1] + position
		}
		for j := range data {
			r = uint8(copy_data[j])
			sorted[C[r]-1] = data[j]
			copy_data[j] >>= 8
		}
	}
	copy(data, sorted)
}
