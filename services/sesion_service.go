package services

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// ObtenerSesiones envía una petición GET para obtener las sesiones activas
func ObtenerSesiones() (interface{}, error) {
	resp, err := http.Get("http://localhost:8081/v1/sesiones")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var resultado interface{}
	json.NewDecoder(resp.Body).Decode(&resultado)
	return resultado, nil
}

// RegistrarSesion envía una petición POST para crear una sesión al hacer login
func RegistrarSesion(sesionData map[string]interface{}) (map[string]interface{}, error) {
	jsonBytes, err := json.Marshal(sesionData)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post("http://localhost:8081/v1/sesiones", "application/json", bytes.NewBuffer(jsonBytes))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var resultado map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&resultado)
	return resultado, nil
}

// CerrarSesion envía una petición DELETE para cerrar la sesión (logout)
func CerrarSesion(id string) (map[string]interface{}, error) {
	req, err := http.NewRequest("DELETE", "http://localhost:8081/v1/sesiones/"+id, nil)
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