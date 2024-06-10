package controllers

import (
	"factorizacion-QR-api/pkg/entities"
	"factorizacion-QR-api/pkg/services"
	"github.com/gofiber/fiber/v2"
)

// Funciones del controller

// Función de factorizacion QR
func FactorizacionQR(context *fiber.Ctx) error {

	var Request entities.MatrizRequest

	// Validación
	if Error := context.BodyParser(&Request); Error != nil {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "No se pudo parsear la solicitud",
		})
	}

	filas := len(Request.Matriz)
	columnas := len(Request.Matriz[0])

	// Validación
	if filas == columnas{
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "La matriz no debe tener la misma cantidad de filas y columnas.",
		})
	}

	// Validación
	for i := 1; i < filas; i++ {
		if len(Request.Matriz[0]) != len(Request.Matriz[i]){
			return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "La matriz no debe tener una cantidad diferente de columnas.",
			})
		}
	}

	// Consume la función del servicio
    MatrizQ, MatrizR := services.CalculoFactorizacionQR(Request.Matriz)

    var MatrizRetorno [2][]float64

    // Cargar las matrices Q y R
	MatrizRetorno[0] = MatrizQ
	MatrizRetorno[1] = MatrizR

    return context.JSON(entities.MatrizResponse{Matrices: MatrizRetorno})
}