package entities

// Representa la solicitud que contiene la matriz
type MatrizRequest struct {
	Matriz [][]float64 `json:"matriz"`
}