package utils

import "gonum.org/v1/gonum/mat"

func MatrixToSlice(m *mat.Dense) [][]float64 {
	rows, cols := m.Dims()
	slice := make([][]float64, rows)
	for i := range slice {
		slice[i] = make([]float64, cols)
		for j := range slice[i] {
			slice[i][j] = m.At(i, j)
		}
	}
	return slice
}

func ValidateMatrixDimensions(matrix [][]float64) (int, int, error) {
	rows := len(matrix)
	if rows == 0 {
		return 0, 0, ErrEmptyMatrix
	}
	cols := len(matrix[0])

	for i := 1; i < rows; i++ {
		if len(matrix[i]) != cols {
			return 0, 0, ErrInconsistentColumns
		}
	}

	if cols == 0 {
		return 0, 0, ErrNoColumns
	}

	if rows < cols {
		return 0, 0, ErrInvalidDimensions
	}

	return rows, cols, nil
}
