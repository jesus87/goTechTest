package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//calcularEstadisticas funcion handler para cualcular las estadisticas
func calcularEstadisticas(w http.ResponseWriter, r *http.Request) {
	// Obtener la fecha de inicio y el número de días desde el parametro de la  URL
	fechaInicio := mux.Vars(r)["fechaInicio"]

	numDias, _ := strconv.Atoi(r.URL.Query().Get("dias"))

	// Consultar el API y obtener los datos
	compras := consultarAPI(fechaInicio, numDias)

	// Calcular las estadísticas
	total := 0.0
	comprasPorTDC := make(map[string]float64)
	noCompraron := 0
	compraMasAlta := 0.0

	for _, compra := range compras {
		total += compra.Monto
		comprasPorTDC[compra.TDC] += compra.Monto
		if !compra.Compro {
			noCompraron++
		}
		if compra.Monto > compraMasAlta {
			compraMasAlta = compra.Monto
		}
	}

	// Crear la estructura de resumen
	resumen := Resumen{
		Total:         total,
		ComprasPorTDC: comprasPorTDC,
		NoCompraron:   noCompraron,
		CompraMasAlta: compraMasAlta,
	}

	// Responder con las estadísticas en formato JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resumen)
}
