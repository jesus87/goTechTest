package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	//definimos la ruta y la funcion handler para este endpoint
	r.HandleFunc("/resumen/{fechaInicio}", calcularEstadisticas).Methods("GET")

	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}
