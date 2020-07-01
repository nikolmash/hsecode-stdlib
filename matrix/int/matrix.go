package matrix

type Matrix struct {
	Rows   int
	Cols   int
	matrix []int
}

func New(n, m int) *Matrix {
	M := Matrix{Rows: n, Cols: m, matrix: make([]int, n*m)}
	return &M
}

func (M *Matrix) idx(i, j int) int {
	if i >= M.Rows || j >= M.Cols {
		panic("the indices are out of range")
	} else if i < 0 || j < 0 {
		panic("the indices mustn't be zero")
	}
	return i*M.Cols + j
}

func (M *Matrix) Get(i, j int) int {
	ind := M.idx(i, j)
	return M.matrix[ind]
}

func (M *Matrix) Set(i, j int, v int) {
	ind := M.idx(i, j)
	M.matrix[ind] = v
}
