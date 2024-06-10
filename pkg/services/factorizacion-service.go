package services

import (
	"gonum.org/v1/gonum/mat"
)

// Funciones del service

func CalculoFactorizacionQR(InformacionMatriz [][]float64) ([]float64, []float64) {
	filas := len(InformacionMatriz)
	columnas := len(InformacionMatriz[0])
	primerNumero := filas
	segundoNumero := 0

	vueltas := (filas * columnas)
	matriz := make([]float64, vueltas)

	for i := 0; i < filas; i++ {
		for j := 0; j < columnas; j++ {
			matriz[i * columnas + j] = InformacionMatriz[i][j]
		}
	}

	if filas > columnas {
		primerNumero = filas
		segundoNumero = columnas
	}

	if columnas > filas {
		primerNumero = columnas
		segundoNumero = filas
	}

	NuevaMatriz := mat.NewDense(primerNumero, segundoNumero, matriz)

	// Creamos una matriz para almacenar el resultado
	var MatrizQR mat.QR
	MatrizQR.Factorize(NuevaMatriz)

	// Extraemos las matrices Q y R
	var Q mat.Dense
	var R mat.Dense
	MatrizQR.QTo(&Q)
	MatrizQR.RTo(&R)

	ParteMatrizQ := Q.RawMatrix()
	ParteMatrizR := R.RawMatrix()

	// Crea un segmento para contener los datos
	QData := make([]float64, len(ParteMatrizQ.Data))
	RData := make([]float64, len(ParteMatrizR.Data))

	// Copia datos
	copy(QData, ParteMatrizQ.Data)
	copy(RData, ParteMatrizR.Data)

	return QData, RData
}
