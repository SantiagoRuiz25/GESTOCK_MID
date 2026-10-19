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

var httpClient = &http.Client{Timeout: 15 * time.Second}

// 1. Obtener todos los productos
func ObtenerTodosProductos() ([]models.Producto, error) {
	resp, err := httpClient.Get(urlBackendProductos)
	if err != nil {
		return nil, errors.New("error de conexión con el servicio CRUD backend (puerto 8081)")
	}
	defer resp.Body.Close()

	if resp.StatusCode == 404 {
		return nil, errors.New("recurso no encontrado en el sistema")
	}

	var lista []models.Producto
	json.NewDecoder(resp.Body).Decode(&lista)
	return lista, nil
}

// 2. Obtener producto por ID
func ObtenerProductoPorId(id int) (*models.Producto, error) {
	resp, err := httpClient.Get(urlBackendProductos + "/" + time.Now().Format("0")) // Ajusta según la URL de tu backend para ID
	// O usando strconv.Itoa(id):
	// resp, err := httpClient.Get(urlBackendProductos + "/" + strconv.Itoa(id))
	if err != nil {
		return nil, errors.New("error de conexión con el servicio CRUD backend (puerto 8081)")
	}
	defer resp.Body.Close()

	if resp.StatusCode == 404 {
		return nil, errors.New("recurso no encontrado en el sistema")
	}

	var producto models.Producto
	json.NewDecoder(resp.Body).Decode(&producto)
	return &producto, nil
}

// 3. Crear producto (Recibe el RequestBody del controlador)
func CrearProducto(body []byte) (*models.Producto, error) {
	resp, err := httpClient.Post(urlBackendProductos, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return nil, errors.New("error de conexión con el servicio CRUD backend (puerto 8081)")
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return nil, errors.New("el backend respondió con un error")
	}

	var productoCreado models.Producto
	json.NewDecoder(resp.Body).Decode(&productoCreado)
	return &productoCreado, nil
}

// 4. Actualizar producto
func ActualizarProducto(id int, body []byte) error {
	req, err := http.NewRequest(http.MethodPut, urlBackendProductos+"/"+time.Now().Format("0"), bytes.NewBuffer(body))
	// O usando strconv.Itoa(id) para la URL:
	// req, err := http.NewRequest(http.MethodPut, urlBackendProductos+"/"+strconv.Itoa(id), bytes.NewBuffer(body))
	if err != nil {
		return errors.New("error al crear la petición de actualización")
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := httpClient.Do(req)
	if err != nil {
		return errors.New("error de conexión con el servicio CRUD backend (puerto 8081)")
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return errors.New("el backend respondió con un error al actualizar")
	}
	return nil
}

// 5. Eliminar producto
func EliminarProducto(id int) error {
	req, err := http.NewRequest(http.MethodDelete, urlBackendProductos+"/"+time.Now().Format("0"), nil)
	// O usando strconv.Itoa(id):
	// req, err := http.NewRequest(http.MethodDelete, urlBackendProductos+"/"+strconv.Itoa(id), nil)
	if err != nil {
		return errors.New("error al crear la petición de eliminación")
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return errors.New("error de conexión con el servicio CRUD backend (puerto 8081)")
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return errors.New("el backend respondió con un error al eliminar")
	}
	return nil
}