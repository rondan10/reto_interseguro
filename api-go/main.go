package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"gonum.org/v1/gonum/mat"
)

type MatrixRequest struct {
	Matrix [][]float64 `json:"matrix"`
}

type QRResponse struct {
	Q [][]float64 `json:"q"`
	R [][]float64 `json:"r"`
}

func matrixToSlice(m *mat.Dense) [][]float64 {
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

func calculateQR(matrix [][]float64) (QRResponse, error) {
	rows := len(matrix)
	if rows == 0 {
		return QRResponse{}, fiber.NewError(fiber.StatusBadRequest, "Matrix cannot be empty")
	}
	cols := len(matrix[0])

	for i := 1; i < rows; i++ {
		if len(matrix[i]) != cols {
			return QRResponse{}, fiber.NewError(fiber.StatusBadRequest, "All rows must have the same number of columns")
		}
	}

	if cols == 0 {
		return QRResponse{}, fiber.NewError(fiber.StatusBadRequest, "Matrix must have at least one column")
	}

	if rows < cols {
		return QRResponse{}, fiber.NewError(fiber.StatusBadRequest,
			"Para la factorización QR, el número de filas debe ser mayor o igual al número de columnas")
	}

	flat := make([]float64, rows*cols)
	for i := range matrix {
		copy(flat[i*cols:(i+1)*cols], matrix[i])
	}

	m := mat.NewDense(rows, cols, flat)

	var qr mat.QR
	qr.Factorize(m)

	var q mat.Dense
	var r mat.Dense

	q.Reset()
	qr.QTo(&q)

	r.Reset()
	qr.RTo(&r)

	qRows, qCols := q.Dims()
	rRows, rCols := r.Dims()

	if qRows == 0 || qCols == 0 || rRows == 0 || rCols == 0 {
		return QRResponse{}, fiber.NewError(fiber.StatusInternalServerError,
			fmt.Sprintf("Invalid dimensions after factorization: Q(%d,%d), R(%d,%d)", qRows, qCols, rRows, rCols))
	}

	return QRResponse{
		Q: matrixToSlice(&q),
		R: matrixToSlice(&r),
	}, nil
}

func sendToNodeAPI(qr QRResponse) error {
	nodeAPIURL := "http://localhost:3001/api/process-qr"

	jsonData, err := json.Marshal(qr)
	if err != nil {
		return err
	}

	resp, err := http.Post(nodeAPIURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("node.js API returned status: %d", resp.StatusCode)
	}

	return nil
}

func main() {
	app := fiber.New()

	// Middlewareee
	app.Use(logger.New())
	app.Use(cors.New())

	app.Post("/api/qr", func(c *fiber.Ctx) error {
		var req MatrixRequest
		if err := c.BodyParser(&req); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
		}

		qr, err := calculateQR(req.Matrix)
		if err != nil {
			return err
		}

		if err := sendToNodeAPI(qr); err != nil {
			log.Printf("Error sending to Node.js API: %v", err)
		}

		return c.JSON(qr)
	})

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	log.Fatal(app.Listen(":3000"))
}
