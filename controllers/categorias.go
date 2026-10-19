package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"

	"GESTOCK_MID/models"
	web "github.com/beego/beego/v2/server/web"
)

type CategoriasController struct {
	web.Controller
}

const crudCategoriaURL = "http://localhost:8081/v1/categorias"

// 1. GET ALL: Consultar todas las categorías
func (c *CategoriasController) GetAll() {
	respuesta, errCrud := http.Get(crudCategoriaURL)
	if errCrud != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  500,
			"Message": "Error de conexión con la base de datos o servicio CRUD",
		}
		c.ServeJSON()
		return
	}
	defer respuesta.Body.Close()

	var lista []models.Categoria
	json.NewDecoder(respuesta.Body).Decode(&lista)

	c.Ctx.Output.SetStatus(200)
	c.Data["json"] = map[string]interface{}{
		"Success": true,
		"Status":  200,
		"Message": "Lista de categorías obtenida con éxito",
		"Data":    lista,
	}
	c.ServeJSON()
}

// 2. GET BY ID: Consultar categoría por ID
func (c *CategoriasController) GetById() {
	idCategoria := c.Ctx.Input.Param(":id")

	if idCategoria == "" {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  400,
			"Message": "El ID es requerido",
		}
		c.ServeJSON()
		return
	}

	id, err := strconv.Atoi(idCategoria)
	if err != nil || id <= 0 {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  400,
			"Message": "Id ingresado es invalido",
		}
		c.ServeJSON()
		return
	}

	respuesta, errCrud := http.Get(crudCategoriaURL + "/" + strconv.Itoa(id))
	if errCrud != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  500,
			"Message": "Error de conexión con la base de datos o servicio CRUD",
		}
		c.ServeJSON()
		return
	}
	defer respuesta.Body.Close()

	if respuesta.StatusCode == 404 {
		c.Ctx.Output.SetStatus(404)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  404,
			"Message": "No se encontró la categoría solicitada en el sistema",
		}
		c.ServeJSON()
		return
	}

	var categoria models.Categoria
	json.NewDecoder(respuesta.Body).Decode(&categoria)

	c.Ctx.Output.SetStatus(200)
	c.Data["json"] = map[string]interface{}{
		"Success": true,
		"Status":  200,
		"Message": "Categoría obtenida con éxito",
		"Data":    categoria,
	}
	c.ServeJSON()
}

// 3. POST: Crear nueva categoría
func (c *CategoriasController) Post() {
	var nuevaCategoria models.Categoria

	// Lectura robusta dual del cuerpo de la petición
	bodyBytes := c.Ctx.Input.RequestBody
	var err error
	if len(bodyBytes) == 0 && c.Ctx.Request.Body != nil {
		err = json.NewDecoder(c.Ctx.Request.Body).Decode(&nuevaCategoria)
	} else {
		err = json.Unmarshal(bodyBytes, &nuevaCategoria)
	}

	if err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  400,
			"Message": "El formato del JSON recibido es invalido",
			"Error":   err.Error(),
		}
		c.ServeJSON()
		return
	}

	if nuevaCategoria.Nombre == "" {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  400,
			"Message": "Campos obligatorios incompletos",
		}
		c.ServeJSON()
		return
	}

	// Serializar y enviar al CRUD (puerto 8081)
	jsonBody, _ := json.Marshal(nuevaCategoria)
	resp, err := http.Post(crudCategoriaURL, "application/json", bytes.NewBuffer(jsonBody))
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

	var resultado interface{}
	json.NewDecoder(resp.Body).Decode(&resultado)

	if resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusOK {
		c.Ctx.Output.SetStatus(resp.StatusCode)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  resp.StatusCode,
			"Message": "No se pudo crear la categoría en el CRUD",
			"Error":   resultado,
		}
		c.ServeJSON()
		return
	}

	c.Ctx.Output.SetStatus(201)
	c.Data["json"] = map[string]interface{}{
		"Success": true,
		"Status":  201,
		"Message": "Categoría creada exitosamente",
		"Data":    resultado,
	}
	c.ServeJSON()
}

// 4. PUT: Actualizar categoría
func (c *CategoriasController) Put() {
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

	var categoria models.Categoria
	bodyBytes := c.Ctx.Input.RequestBody
	if len(bodyBytes) == 0 && c.Ctx.Request.Body != nil {
		json.NewDecoder(c.Ctx.Request.Body).Decode(&categoria)
	} else {
		json.Unmarshal(bodyBytes, &categoria)
	}

	categoria.ID = id

	jsonBody, err := json.Marshal(categoria)
	if err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  400,
			"Message": "Error al serializar la categoría para actualizar",
		}
		c.ServeJSON()
		return
	}

	req, err := http.NewRequest(http.MethodPut, crudCategoriaURL+"/"+idParam, bytes.NewBuffer(jsonBody))
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
			"Message": "No se pudo actualizar la categoría",
			"Error":   resultado,
		}
		c.ServeJSON()
		return
	}

	c.Ctx.Output.SetStatus(200)
	c.Data["json"] = map[string]interface{}{
		"Success": true,
		"Status":  200,
		"Message": "Categoría actualizada correctamente",
		"Data":    resultado,
	}
	c.ServeJSON()
}

// 5. DELETE: Eliminar categoría
func (c *CategoriasController) Delete() {
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

	req, err := http.NewRequest(http.MethodDelete, crudCategoriaURL+"/"+idParam, nil)
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
			"Message": "No se pudo eliminar la categoría",
			"Error":   resultado,
		}
		c.ServeJSON()
		return
	}

	c.Ctx.Output.SetStatus(200)
	c.Data["json"] = map[string]interface{}{
		"Success": true,
		"Status":  200,
		"Message": "Categoría eliminada correctamente",
	}
	c.ServeJSON()
}