package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"

	"GESTOCK_MID/models"

	beego "github.com/beego/beego/v2/server/web"
)

type BodegaController struct {
	beego.Controller
}

// URL base del microservicio CRUD (Puerto 8081)
const crudURL = "http://localhost:8081/v1/bodega"

// ==========================================
// 1. GET ALL: Obtener todas las bodegas
// ==========================================
func (c *BodegaController) GetAll() {
	resp, err := http.Get(crudURL)
	if err != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  500,
			"Message": "Error de conexión con el servicio CRUD (8081)",
		}
		c.ServeJSON()
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		c.Ctx.Output.SetStatus(resp.StatusCode)
		var errorCrud map[string]interface{}
		json.NewDecoder(resp.Body).Decode(&errorCrud)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  resp.StatusCode,
			"Message": "Error devuelto por el CRUD",
			"Error":   errorCrud,
		}
		c.ServeJSON()
		return
	}

	var bodegas []models.Bodegas
	if err := json.NewDecoder(resp.Body).Decode(&bodegas); err != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  500,
			"Message": "Error al procesar la respuesta del CRUD",
		}
		c.ServeJSON()
		return
	}

	c.Ctx.Output.SetStatus(200)
	c.Data["json"] = map[string]interface{}{
		"Success": true,
		"Status":  200,
		"Message": "Bodegas obtenidas con éxito",
		"Data":    bodegas,
	}
	c.ServeJSON()
}

// ==========================================
// 2. GET BY ID: Obtener una bodega por ID
// ==========================================
func (c *BodegaController) GetById() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  400,
			"Message": "El ID ingresado es inválido",
		}
		c.ServeJSON()
		return
	}

	resp, err := http.Get(crudURL + "/" + idStr)
	if err != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  500,
			"Message": "Error de conexión con el servicio CRUD (8081)",
		}
		c.ServeJSON()
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		c.Ctx.Output.SetStatus(resp.StatusCode)
		var errorCrud map[string]interface{}
		json.NewDecoder(resp.Body).Decode(&errorCrud)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  resp.StatusCode,
			"Message": "Bodega no encontrada o error en el CRUD",
			"Error":   errorCrud,
		}
		c.ServeJSON()
		return
	}

	var bodega models.Bodegas
	if err := json.NewDecoder(resp.Body).Decode(&bodega); err != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  500,
			"Message": "Error al procesar los datos de la bodega",
		}
		c.ServeJSON()
		return
	}

	c.Ctx.Output.SetStatus(200)
	c.Data["json"] = map[string]interface{}{
		"Success": true,
		"Status":  200,
		"Message": "Bodega obtenida con éxito",
		"Data":    bodega,
	}
	c.ServeJSON()
}

// ==========================================
// 3. POST: Crear una bodega
// ==========================================
func (c *BodegaController) Post() {
	var bodega models.Bodegas
	
	// Usamos NewDecoder con DiscardUnknown o lectura directa para asegurar el mapeo de minúsculas
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &bodega)
	if err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  400,
			"Message": "El formato del JSON recibido es inválido",
			"Error":   err.Error(), // Muestra el detalle exacto si falla el parseo
		}
		c.ServeJSON()
		return
	}

	// Validación básica de negocio (permitiendo código > 0)
	if bodega.Nombre == "" || bodega.Codigo == 0 {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  400,
			"Message": "Los campos 'nombre' y 'codigo' son obligatorios",
		}
		c.ServeJSON()
		return
	}

	jsonBody, _ := json.Marshal(bodega)
	resp, err := http.Post(crudURL, "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  500,
			"Message": "Error de conexión con el servicio CRUD (8081)",
		}
		c.ServeJSON()
		return
	}
	defer resp.Body.Close()

	var resultado interface{}
	json.NewDecoder(resp.Body).Decode(&resultado)

	if resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusOK {
		c.Ctx.Output.SetStatus(resp.StatusCode)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  resp.StatusCode,
			"Message": "No se pudo crear la bodega en el CRUD",
			"Error":   resultado,
		}
		c.ServeJSON()
		return
	}

	c.Ctx.Output.SetStatus(201)
	c.Data["json"] = map[string]interface{}{
		"Success": true,
		"Status":  201,
		"Message": "Bodega creada exitosamente",
		"Data":    resultado,
	}
	c.ServeJSON()
}

// ==========================================
// 4. PUT: Actualizar una bodega
// ==========================================
func (c *BodegaController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  400,
			"Message": "El ID ingresado es inválido",
		}
		c.ServeJSON()
		return
	}

	var bodega models.Bodegas
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &bodega); err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  400,
			"Message": "El formato del JSON recibido es inválido",
			"Error":   err.Error(),
		}
		c.ServeJSON()
		return
	}

	jsonBody, _ := json.Marshal(bodega)
	req, err := http.NewRequest(http.MethodPut, crudURL+"/"+idStr, bytes.NewBuffer(jsonBody))
	if err != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  500,
			"Message": "Error al preparar la petición PUT",
		}
		c.ServeJSON()
		return
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  500,
			"Message": "Error de conexión con el servicio CRUD (8081)",
		}
		c.ServeJSON()
		return
	}
	defer resp.Body.Close()

	var resultado interface{}
	json.NewDecoder(resp.Body).Decode(&resultado)

	if resp.StatusCode != http.StatusOK {
		c.Ctx.Output.SetStatus(resp.StatusCode)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  resp.StatusCode,
			"Message": "No se pudo actualizar la bodega",
			"Error":   resultado,
		}
		c.ServeJSON()
		return
	}

	c.Ctx.Output.SetStatus(200)
	c.Data["json"] = map[string]interface{}{
		"Success": true,
		"Status":  200,
		"Message": "Bodega actualizada correctamente",
		"Data":    resultado,
	}
	c.ServeJSON()
}

// ==========================================
// 5. DELETE: Eliminar una bodega
// ==========================================
func (c *BodegaController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  400,
			"Message": "El ID ingresado es inválido",
		}
		c.ServeJSON()
		return
	}

	req, err := http.NewRequest(http.MethodDelete, crudURL+"/"+idStr, nil)
	if err != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  500,
			"Message": "Error al preparar la petición DELETE",
		}
		c.ServeJSON()
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  500,
			"Message": "Error de conexión con el servicio CRUD (8081)",
		}
		c.ServeJSON()
		return
	}
	defer resp.Body.Close()

	var resultado interface{}
	json.NewDecoder(resp.Body).Decode(&resultado)

	if resp.StatusCode != http.StatusOK {
		c.Ctx.Output.SetStatus(resp.StatusCode)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  resp.StatusCode,
			"Message": "No se pudo eliminar la bodega",
			"Error":   resultado,
		}
		c.ServeJSON()
		return
	}

	c.Ctx.Output.SetStatus(200)
	c.Data["json"] = map[string]interface{}{
		"Success": true,
		"Status":  200,
		"Message": "Bodega eliminada correctamente",
	}
	c.ServeJSON()
}