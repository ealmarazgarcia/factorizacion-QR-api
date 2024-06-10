package main

import (
	"factorizacion-QR-api/api/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Configurar las rutas
	routes.SetupRoutes(app)

	// Iniciar el servidor
	app.Listen(":8080")
}
