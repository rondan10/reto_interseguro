package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"matrix-api/internal/domain"
	"matrix-api/internal/service"

	"github.com/gofiber/fiber/v2"
)

type MatrixHandler struct {
	matrixService *service.MatrixService
}

func NewMatrixHandler(matrixService *service.MatrixService) *MatrixHandler {
	return &MatrixHandler{
		matrixService: matrixService,
	}
}

func (h *MatrixHandler) HandleQRFactorization(c *fiber.Ctx) error {
	var req domain.MatrixRequest
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}

	qr, err := h.matrixService.CalculateQR(req.Matrix)
	if err != nil {
		return err
	}

	if err := h.sendToNodeAPI(qr); err != nil {
		log.Printf("Error sending to Node.js API: %v", err)
	}

	return c.JSON(qr)
}

func (h *MatrixHandler) HealthCheck(c *fiber.Ctx) error {
	return c.SendString("OK")
}

func (h *MatrixHandler) sendToNodeAPI(qr domain.QRResponse) error {
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
