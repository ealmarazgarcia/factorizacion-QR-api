package entities

// Representa la respuesta que contienen las matrices
type MatrizResponse struct {
	Matrices [2][]float64 `json:"matrices"`
}