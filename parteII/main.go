package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type UserRoles struct {
	Username string   `json:"username"`
	Roles    []string `json:"roles"`
}

type OrganizationData struct {
	Organization string      `json:"organization"`
	Users        []UserRoles `json:"users"`
}

func procesarCSV(input io.Reader) ([]OrganizationData, error) {
	// Lee el archivo CSV
	reader := csv.NewReader(input)
	lines, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	organizations := make([]OrganizationData, 0)
	currentOrganization := ""
	currentUser := ""
	roles := make([]string, 0)

	// realimos las iteraciones para iterar sobre las lineas recuperadas
	for i, line := range lines {

		if i == 0 {
			continue
		}
		organization := line[0]
		user := line[1]
		role := line[2]

		// validamos si la organización ha cambiado
		if organization != currentOrganization {
			// Agregamos la organización anterior a la estructura de datos
			if currentOrganization != "" {
				userRoles := UserRoles{Username: currentUser, Roles: roles}
				lastOrganization := &organizations[len(organizations)-1]
				lastOrganization.Users = append(lastOrganization.Users, userRoles)
			}

			// inicializamos la nueva organización
			currentOrganization = organization
			currentUser = user
			roles = []string{role}
			organizations = append(organizations, OrganizationData{Organization: currentOrganization, Users: []UserRoles{}})
		} else if user != currentUser {
			// verificamos si el usuario actual ha cambiado
			currentUser = user
			roles = []string{role}
		} else {
			// agregamos roles
			roles = append(roles, role)
		}
	}

	// por ultimo agregamos la última organización a la estructura de datos
	if currentOrganization != "" {
		userRoles := UserRoles{Username: currentUser, Roles: roles}
		lastOrganization := &organizations[len(organizations)-1]
		lastOrganization.Users = append(lastOrganization.Users, userRoles)
	}

	// finalmente convertimos a JSON
	return organizations, nil
}
func main() {
	// abrimos el archivo CSV
	file, err := os.Open("datos.csv")
	if err != nil {
		fmt.Println("Error al abrir el archivo:", err)
		return
	}
	defer file.Close()

	// llamamos a la funcion que se encarga de realizar el procesamiento
	organizations, err := procesarCSV(file)

	// convertimos el resultado a formato json
	resultJson, err := json.Marshal(organizations)
	if err != nil {
		fmt.Println("Error al convertir a JSON:", err)
		return
	}

	// Imprimimos el resultado
	fmt.Println(string(resultJson))
}
