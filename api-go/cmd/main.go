package main

import (
	"log"

	"matrix-api/internal/infrastructure/http/handlers"
	"matrix-api/internal/service"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()

	app.Use(logger.New())
	app.Use(cors.New())

	matrixService := service.NewMatrixService()
	matrixHandler := handlers.NewMatrixHandler(matrixService)

	// Rutas
	app.Post("/api/qr", matrixHandler.HandleQRFactorization)
	app.Get("/health", matrixHandler.HealthCheck)

	log.Fatal(app.Listen(":3000"))
}
