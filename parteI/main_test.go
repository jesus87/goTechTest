package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestConsultarAPI(t *testing.T) {
	// Crear un servidor de prueba para simular la respuesta HTTP
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Simular una respuesta JSON del servidor externo
		respJSON := `[{
			"clientId": 1000223,
			"phone": "992003040",
			"nombre": "Juan Mata",
			"compro": true,
			"tdc": "gold",
			"monto": 123.20,
			"date": "2020-02-20"
		}]`

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(respJSON))
	}))

	// Asegurarnos de cerrar el servidor de prueba al final de la función
	defer server.Close()

	// Llamar a la función consultarAPI con el URL del servidor de prueba
	compras := consultarAPI("2020-02-20", 1)

	// Verificar que se haya devuelto al menos una compra
	//si no, lanzamos un error
	if len(compras) == 0 {
		t.Errorf("Se esperaba al menos una compra, pero se encontraron %d", len(compras))
	}

}

func TestCalcularEstadisticas(t *testing.T) {
	req, err := http.NewRequest("GET", "/resumen/2019-12-01?dias=5", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Configurar un enrutador Mux
	r := mux.NewRouter()
	r.HandleFunc("/resumen/{fechaInicio}", calcularEstadisticas).Methods("GET")

	rr := httptest.NewRecorder()

	// Llamar al manejador calcularEstadisticas
	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("El manejador devolvió un código de estado incorrecto: %v, esperado %v", status, http.StatusOK)
	}

}
