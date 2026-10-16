package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"GESTOCK_MID/models"
)

const urlBackendBodegas = "http://localhost:8081/v1/bodega"

// Cliente HTTP configurado con un tiempo límite de espera
var httpClientBodegas = &http.Client{
	Timeout: 15 * time.Second,
}

// 1. Obtener todas las bodegas
func GetAllBodegasService() ([]models.Bodegas, error) {
	resp, err := httpClientBodegas.Get(urlBackendBodegas)
	if err != nil {
		return nil, errors.New("error de conexión con el servicio CRUD backend (puerto 8081)")
	}
	defer resp.Body.Close()

	if resp.StatusCode == 404 {
		return nil, errors.New("recurso no encontrado en el sistema")
	}
	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("el backend respondió con un estado de error: %d", resp.StatusCode)
	}

	var lista []models.Bodegas
	if err := json.NewDecoder(resp.Body).Decode(&lista); err != nil {
		return nil, errors.New("error al deserializar la respuesta JSON del backend")
	}

	return lista, nil
}

// 2. Obtener una bodega por ID
func GetBodegaByIdService(id int) (*models.Bodegas, error) {
	url := fmt.Sprintf("%s/%d", urlBackendBodegas, id)
	resp, err := httpClientBodegas.Get(url)
	if err != nil {
		return nil, errors.New("error de conexión con el servicio CRUD backend (puerto 8081)")
	}
	defer resp.Body.Close()

	if resp.StatusCode == 404 {
		return nil, errors.New("recurso no encontrado en el sistema")
	}
	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("el backend respondió con un estado de error: %d", resp.StatusCode)
	}

	var bodega models.Bodegas
	if err := json.NewDecoder(resp.Body).Decode(&bodega); err != nil {
		return nil, errors.New("error al deserializar la respuesta JSON del backend")
	}

	return &bodega, nil
}

// 3. Crear una nueva bodega
func CreateBodegaService(bodega *models.Bodegas) error {
	// Regla de negocio: Validar campos obligatorios
	if bodega.Nombre == "" || bodega.Codigo == 0 {
		return errors.New("el nombre y el código de la bodega son obligatorios")
	}

	jsonBytes, err := json.Marshal(bodega)
	if err != nil {
		return errors.New("error al serializar el cuerpo de la petición")
	}

	resp, err := httpClientBodegas.Post(urlBackendBodegas, "application/json", bytes.NewBuffer(jsonBytes))
	if err != nil {
		return errors.New("error de conexión con el servicio CRUD backend (puerto 8081)")
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return fmt.Errorf("el backend respondió con un estado de error: %d", resp.StatusCode)
	}

	return nil
}

// 4. Actualizar una bodega
func UpdateBodegaService(id int, bodega *models.Bodegas) error {
	url := fmt.Sprintf("%s/%d", urlBackendBodegas, id)

	jsonBytes, err := json.Marshal(bodega)
	if err != nil {
		return errors.New("error al serializar el cuerpo de la petición")
	}

	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(jsonBytes))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := httpClientBodegas.Do(req)
	if err != nil {
		return errors.New("error de conexión con el servicio CRUD backend (puerto 8081)")
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return fmt.Errorf("el backend respondió con un estado de error: %d", resp.StatusCode)
	}

	return nil
}

// 5. Eliminar una bodega
func DeleteBodegaService(id int) error {
	url := fmt.Sprintf("%s/%d", urlBackendBodegas, id)

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}

	resp, err := httpClientBodegas.Do(req)
	if err != nil {
		return errors.New("error de conexión con el servicio CRUD backend (puerto 8081)")
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return fmt.Errorf("el backend respondió con un estado de error: %d", resp.StatusCode)
	}

	return nil
}