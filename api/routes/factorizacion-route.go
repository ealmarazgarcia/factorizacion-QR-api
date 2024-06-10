package routes

import (
	"factorizacion-QR-api/api/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// Definición de los routes

	// Ruta para la factorizacion QR de la matriz
	app.Post("/factorizacionQR", controllers.FactorizacionQR)
}
