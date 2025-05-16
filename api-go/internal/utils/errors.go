package utils

import (
	"github.com/gofiber/fiber/v2"
)

var (
	ErrEmptyMatrix         = fiber.NewError(fiber.StatusBadRequest, "Matrix cannot be empty")
	ErrInconsistentColumns = fiber.NewError(fiber.StatusBadRequest, "All rows must have the same number of columns")
	ErrNoColumns           = fiber.NewError(fiber.StatusBadRequest, "Matrix must have at least one column")
	ErrInvalidDimensions   = fiber.NewError(fiber.StatusBadRequest,
		"Para la factorización QR, el número de filas debe ser mayor o igual al número de columnas")
)
