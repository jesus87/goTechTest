package main

import (
	"strings"
	"testing"
)

func TestProcesarCSV(t *testing.T) {
	// Define un ejemplo de datos CSV en formato string
	csvData := `organization,usuario,rol
org1,jperez,admin
org1,jperez,superadmin
org1,asosa,writer
org2,jperez,admin
org2,rrodriguez,writer
org2,rrodriguez,editor
`

	// Convierte el ejemplo de datos CSV en un lector
	reader := strings.NewReader(csvData)

	// Llama a la función de procesamiento con el lector
	organizations, err := procesarCSV(reader)
	if err != nil {
		t.Errorf("Error al procesar CSV: %v", err)
	}

	// Realiza aserciones en los datos resultantes (organizations)
	// ...

	// Por ejemplo, puedes verificar si se han creado las organizaciones y usuarios correctamente.
	if len(organizations) != 2 {
		t.Errorf("Se esperaban 2 organizaciones, pero se encontraron %d", len(organizations))
	}

}
