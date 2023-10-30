package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

//consultarAPI consulta el api indicada basando en los parametros de fecha inicial y numero de dias
func consultarAPI(fechaInicio string, numDias int) []Compra {
	compras := make([]Compra, 0)

	for i := 0; i < numDias; i++ {
		fecha := calcularFecha(fechaInicio, i)

		url := "https://apirecruit-gjvkhl2c6a-uc.a.run.app/compras/" + fecha

		resp, err := http.Get(url)
		if err != nil {
			fmt.Println("Error al consultar el API:", err)
			continue
		}

		var comprasDia []Compra
		err = json.NewDecoder(resp.Body).Decode(&comprasDia)
		if err != nil {
			fmt.Println("Error al decodificar la respuesta JSON:", err)
			continue
		}

		compras = append(compras, comprasDia...)
	}

	return compras
}

//calcularFecha funcion para calcular los dias subsecuentes solicitados
func calcularFecha(fechaInicio string, dias int) string {

	inicio, err := time.Parse("2006-01-02", fechaInicio)
	if err != nil {
		fmt.Println("Error al analizar la fecha de inicio:", err)
		return ""
	}

	fecha := inicio.AddDate(0, 0, dias).Format("2006-01-02")
	return fecha
}
