package service

import (
	"fmt"
	"matrix-api/internal/domain"
	"matrix-api/internal/utils"

	"github.com/gofiber/fiber/v2"
	"gonum.org/v1/gonum/mat"
)

type MatrixService struct{}

func NewMatrixService() *MatrixService {
	return &MatrixService{}
}

func (s *MatrixService) CalculateQR(matrix [][]float64) (domain.QRResponse, error) {
	rows, cols, err := utils.ValidateMatrixDimensions(matrix)
	if err != nil {
		return domain.QRResponse{}, err
	}

	flat := make([]float64, rows*cols)
	for i := range matrix {
		copy(flat[i*cols:(i+1)*cols], matrix[i])
	}

	m := mat.NewDense(rows, cols, flat)
	var qr mat.QR
	qr.Factorize(m)

	var q, r mat.Dense
	q.Reset()
	r.Reset()
	qr.QTo(&q)
	qr.RTo(&r)

	// Verifico las dimensiones resultantes
	qRows, qCols := q.Dims()
	rRows, rCols := r.Dims()

	if qRows == 0 || qCols == 0 || rRows == 0 || rCols == 0 {
		return domain.QRResponse{}, fiber.NewError(fiber.StatusInternalServerError,
			fmt.Sprintf("Invalid dimensions after factorization: Q(%d,%d), R(%d,%d)", qRows, qCols, rRows, rCols))
	}

	return domain.QRResponse{
		Q: utils.MatrixToSlice(&q),
		R: utils.MatrixToSlice(&r),
	}, nil
}
