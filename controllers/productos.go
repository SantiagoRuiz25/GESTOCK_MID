package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"

	"GESTOCK_MID/models"

	beego "github.com/beego/beego/v2/server/web"
)

type ProductosController struct {
	beego.Controller
}

const crudProductoURL = "http://localhost:8081/v1/productos"

// ==========================================
// 1. GET ALL: Consultar todos los productos
// ==========================================
func (c *ProductosController) GetAll() {
	resp, err := http.Get(crudProductoURL)
	if err != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  500,
			"Message": "Error de conexión con el servicio CRUD backend (puerto 8081)",
		}
		c.ServeJSON()
		return
	}
	defer resp.Body.Close()

	var productos []models.Producto
	if err := json.NewDecoder(resp.Body).Decode(&productos); err != nil {
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
		"Message": "Lista de productos obtenida con éxito",
		"Data":    productos,
	}
	c.ServeJSON()
}

// ==========================================
// 2. GET BY ID: Consultar producto por ID
// ==========================================
func (c *ProductosController) GetById() {
	idParam := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idParam)
	if err != nil || id <= 0 {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  400,
			"Message": "Id ingresado es inválido",
		}
		c.ServeJSON()
		return
	}

	resp, err := http.Get(crudProductoURL + "/" + idParam)
	if err != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  500,
			"Message": "Error de conexión con el servicio CRUD backend (puerto 8081)",
		}
		c.ServeJSON()
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		c.Ctx.Output.SetStatus(404)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  404,
			"Message": "recurso no encontrado en el sistema",
		}
		c.ServeJSON()
		return
	}

	var producto models.Producto
	if err := json.NewDecoder(resp.Body).Decode(&producto); err != nil {
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
		"Message": "Producto obtenido con éxito",
		"Data":    producto,
	}
	c.ServeJSON()
}

// ==========================================
// 3. POST: Crear nuevo producto
// ==========================================
func (c *ProductosController) Post() {
	var producto models.Producto

	bodyBytes := c.Ctx.Input.RequestBody
	if len(bodyBytes) == 0 && c.Ctx.Request.Body != nil {
		decoder := json.NewDecoder(c.Ctx.Request.Body)
		if err := decoder.Decode(&producto); err != nil {
			c.Ctx.Output.SetStatus(400)
			c.Data["json"] = map[string]interface{}{
				"Success": false,
				"Status":  400,
				"Message": "El formato del JSON recibido es inválido o está vacío",
				"Error":   err.Error(),
			}
			c.ServeJSON()
			return
		}
	} else {
		if err := json.Unmarshal(bodyBytes, &producto); err != nil {
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
	}

	if producto.Nombre == "" || producto.Codigo == "" {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  400,
			"Message": "Los campos 'nombre' y 'codigo' son obligatorios",
		}
		c.ServeJSON()
		return
	}

	jsonBody, err := json.Marshal(producto)
	if err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  400,
			"Message": "Error al serializar el producto",
		}
		c.ServeJSON()
		return
	}

	req, err := http.NewRequest(http.MethodPost, crudProductoURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  500,
			"Message": "Error al preparar la petición POST hacia el CRUD",
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
			"Message": "error de conexión con el servicio CRUD backend (puerto 8081)",
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
			"Message": "No se pudo crear el producto en el CRUD",
			"Error":   resultado,
		}
		c.ServeJSON()
		return
	}

	c.Ctx.Output.SetStatus(201)
	c.Data["json"] = map[string]interface{}{
		"Success": true,
		"Status":  201,
		"Message": "Producto creado exitosamente",
		"Data":    resultado,
	}
	c.ServeJSON()
}

// ==========================================
// 4. PUT: Actualizar producto por ID
// ==========================================
func (c *ProductosController) Put() {
	idParam := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idParam)
	if err != nil || id <= 0 {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  400,
			"Message": "Id ingresado es inválido",
		}
		c.ServeJSON()
		return
	}

	var producto models.Producto
	bodyBytes := c.Ctx.Input.RequestBody
	if len(bodyBytes) == 0 && c.Ctx.Request.Body != nil {
		decoder := json.NewDecoder(c.Ctx.Request.Body)
		if err := decoder.Decode(&producto); err != nil {
			c.Ctx.Output.SetStatus(400)
			c.Data["json"] = map[string]interface{}{
				"Success": false,
				"Status":  400,
				"Message": "El formato del JSON recibido es inválido o está vacío",
				"Error":   err.Error(),
			}
			c.ServeJSON()
			return
		}
	} else {
		if err := json.Unmarshal(bodyBytes, &producto); err != nil {
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
	}

	producto.ID = id

	jsonBody, err := json.Marshal(producto)
	if err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  400,
			"Message": "Error al serializar el producto para actualizar",
		}
		c.ServeJSON()
		return
	}

	req, err := http.NewRequest(http.MethodPut, crudProductoURL+"/"+idParam, bytes.NewBuffer(jsonBody))
	if err != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  500,
			"Message": "Error al preparar la petición PUT hacia el CRUD",
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
			"Message": "error de conexión con el servicio CRUD backend (puerto 8081)",
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
			"Message": "No se pudo actualizar el producto",
			"Error":   resultado,
		}
		c.ServeJSON()
		return
	}

	c.Ctx.Output.SetStatus(200)
	c.Data["json"] = map[string]interface{}{
		"Success": true,
		"Status":  200,
		"Message": "Producto actualizado correctamente",
		"Data":    resultado,
	}
	c.ServeJSON()
}

// ==========================================
// 5. DELETE: Eliminar producto por ID
// ==========================================
func (c *ProductosController) Delete() {
	idParam := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idParam)
	if err != nil || id <= 0 {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  400,
			"Message": "Id ingresado es inválido",
		}
		c.ServeJSON()
		return
	}

	req, err := http.NewRequest(http.MethodDelete, crudProductoURL+"/"+idParam, nil)
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
			"Message": "error de conexión con el servicio CRUD backend (puerto 8081)",
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
			"Message": "No se pudo eliminar el producto",
			"Error":   resultado,
		}
		c.ServeJSON()
		return
	}

	c.Ctx.Output.SetStatus(200)
	c.Data["json"] = map[string]interface{}{
		"Success": true,
		"Status":  200,
		"Message": "Producto eliminado correctamente",
	}
	c.ServeJSON()
}