package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"GESTOCK_MID/models"
)

const urlBackendProductos = "http://localhost:8081/v1/producto"

// Cliente HTTP repetido en los servicios
var httpClient = &http.Client{Timeout: 15 * time.Second}

func GetAllProductosService() ([]models.Producto, error) {
	resp, err := httpClient.Get(urlBackendProductos)
	if err != nil {
		return nil, errors.New("error de conexión con el servicio CRUD backend")
	}
	defer resp.Body.Close()

	if resp.StatusCode == 404 {
		return nil, errors.New("recurso no encontrado en el sistema")
	}

	var lista []models.Producto
	json.NewDecoder(resp.Body).Decode(&lista)
	return lista, nil
}

func CreateProductoService(producto *models.Producto) error {
	if producto.Nombre == "" || producto.Codigo == "" {
		return errors.New("el nombre y el código son obligatorios")
	}

	jsonBytes, _ := json.Marshal(producto)
	resp, err := httpClient.Post(urlBackendProductos, "application/json", bytes.NewBuffer(jsonBytes))
	if err != nil {
		return errors.New("error de conexión con el servicio CRUD backend")
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return errors.New("el backend respondió con un error")
	}
	return nil
}