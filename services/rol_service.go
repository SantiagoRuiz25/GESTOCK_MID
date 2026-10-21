package services

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// CrearRol envía una petición POST al CRUD
func CrearRol(rolData map[string]interface{}) (map[string]interface{}, error) {
	jsonBytes, err := json.Marshal(rolData)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post("http://localhost:8081/v1/roles", "application/json", bytes.NewBuffer(jsonBytes))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var resultado map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&resultado)
	return resultado, nil
}

// ObtenerRoles envía una petición GET al CRUD
func ObtenerRoles() (interface{}, error) {
	resp, err := http.Get("http://localhost:8081/v1/roles")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var resultado interface{}
	json.NewDecoder(resp.Body).Decode(&resultado)
	return resultado, nil
}

// ObtenerRolPorId envía una petición GET por ID al CRUD
func ObtenerRolPorId(id string) (interface{}, error) {
	resp, err := http.Get("http://localhost:8081/v1/roles/" + id)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var resultado interface{}
	json.NewDecoder(resp.Body).Decode(&resultado)
	return resultado, nil
}

// EliminarRol envía una petición DELETE al CRUD
func EliminarRol(id string) (map[string]interface{}, error) {
	req, err := http.NewRequest("DELETE", "http://localhost:8081/v1/roles/"+id, nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var resultado map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&resultado)
	return resultado, nil
}