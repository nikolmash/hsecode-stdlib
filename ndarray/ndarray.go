package ndarray

type NDArray struct {
	shape []int
}

func New(shape ...int) *NDArray {
	for _, dim := range shape {
		if dim < 0 {
			panic("one of the dimensions are negative!")
		}
	}
	N := NDArray{shape: shape}
	return &N
}

func (nda *NDArray) Idx(indicies ...int) int {
	if len(indicies) != len(nda.shape) {
		panic("wrong dimension")
	}
	for i := range indicies {
		if indicies[i] >= nda.shape[i] {
			panic("one of the indexes are out of range")
		}
	}
	res := indicies[0]
	for i := 1; i < len(indicies); i++ {
		res = nda.shape[i]*res + indicies[i]
	}
	return res
}
