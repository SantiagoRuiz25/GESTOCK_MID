package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"GESTOCK_MID/models"
	web "github.com/beego/beego/v2/server/web"
)

const crudCategoriaURL = "http://localhost:8081/v1/categorias"

type CategoriasController struct {
	web.Controller
}

// GetAll consulta todas las categorías desde el servicio CRUD.
func (c *CategoriasController) GetAll() {
	respuesta, errCrud := http.Get(crudCategoriaURL)
	if errCrud != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  http.StatusInternalServerError,
			"Message": "Error de conexión con la base de datos o servicio CRUD",
		}
		_ = c.ServeJSON()
		return
	}
	defer respuesta.Body.Close()

	var lista []models.Categoria
	_ = json.NewDecoder(respuesta.Body).Decode(&lista)

	c.Ctx.Output.SetStatus(http.StatusOK)
	c.Data["json"] = map[string]interface{}{
		"Success": true,
		"Status":  http.StatusOK,
		"Message": "Lista de categorías obtenida con éxito",
		"Data":    lista,
	}
	_ = c.ServeJSON()
}

// GetById consulta una categoría por ID desde el servicio CRUD.
func (c *CategoriasController) GetById() {
	idCategoria := c.Ctx.Input.Param(":id")
	if strings.TrimSpace(idCategoria) == "" {
		c.Ctx.Output.SetStatus(http.StatusBadRequest)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  http.StatusBadRequest,
			"Message": "El ID es obligatorio",
		}
		_ = c.ServeJSON()
		return
	}

	id, err := strconv.Atoi(idCategoria)
	if err != nil || id <= 0 {
		c.Ctx.Output.SetStatus(http.StatusBadRequest)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  http.StatusBadRequest,
			"Message": "ID ingresado es inválido",
		}
		_ = c.ServeJSON()
		return
	}

	respuesta, errCrud := http.Get(crudCategoriaURL + "/" + strconv.Itoa(id))
	if errCrud != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  http.StatusInternalServerError,
			"Message": "Error de conexión con la base de datos o servicio CRUD",
		}
		_ = c.ServeJSON()
		return
	}
	defer respuesta.Body.Close()

	if respuesta.StatusCode == http.StatusNotFound {
		c.Ctx.Output.SetStatus(http.StatusNotFound)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  http.StatusNotFound,
			"Message": "No se encontró la categoría solicitada en el sistema",
		}
		_ = c.ServeJSON()
		return
	}

	var categoria models.Categoria
	_ = json.NewDecoder(respuesta.Body).Decode(&categoria)

	c.Ctx.Output.SetStatus(http.StatusOK)
	c.Data["json"] = map[string]interface{}{
		"Success": true,
		"Status":  http.StatusOK,
		"Message": "Categoría obtenida con éxito",
		"Data":    categoria,
	}
	_ = c.ServeJSON()
}

// Post crea una nueva categoría enviándola al servicio CRUD.
func (c *CategoriasController) Post() {
	var nuevaCategoria models.Categoria

	bodyBytes := c.Ctx.Input.RequestBody
	var err error
	if len(bodyBytes) == 0 && c.Ctx.Request.Body != nil {
		err = json.NewDecoder(c.Ctx.Request.Body).Decode(&nuevaCategoria)
	} else {
		err = json.Unmarshal(bodyBytes, &nuevaCategoria)
	}

	if err != nil {
		c.Ctx.Output.SetStatus(http.StatusBadRequest)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  http.StatusBadRequest,
			"Message": "El formato del JSON recibido es inválido",
			"Error":   err.Error(),
		}
		_ = c.ServeJSON()
		return
	}

	nuevaCategoria.Nombre = strings.TrimSpace(nuevaCategoria.Nombre)
	if nuevaCategoria.Nombre == "" {
		c.Ctx.Output.SetStatus(http.StatusBadRequest)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  http.StatusBadRequest,
			"Message": "Campos obligatorios incompletos",
		}
		_ = c.ServeJSON()
		return
	}

	jsonBody, _ := json.Marshal(nuevaCategoria)
	resp, err := http.Post(crudCategoriaURL, "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  http.StatusInternalServerError,
			"Message": "Error de conexión con el servicio CRUD backend (puerto 8081)",
		}
		_ = c.ServeJSON()
		return
	}
	defer resp.Body.Close()

	var resultado interface{}
	_ = json.NewDecoder(resp.Body).Decode(&resultado)

	if resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusOK {
		c.Ctx.Output.SetStatus(resp.StatusCode)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  resp.StatusCode,
			"Message": "No se pudo crear la categoría en el CRUD",
			"Error":   resultado,
		}
		_ = c.ServeJSON()
		return
	}

	c.Ctx.Output.SetStatus(http.StatusCreated)
	c.Data["json"] = map[string]interface{}{
		"Success": true,
		"Status":  http.StatusCreated,
		"Message": "Categoría creada exitosamente",
		"Data":    resultado,
	}
	_ = c.ServeJSON()
}

// Put actualiza una categoría existente por ID.
func (c *CategoriasController) Put() {
	idParam := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idParam)
	if err != nil || id <= 0 {
		c.Ctx.Output.SetStatus(http.StatusBadRequest)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  http.StatusBadRequest,
			"Message": "ID ingresado es inválido",
		}
		_ = c.ServeJSON()
		return
	}

	var categoria models.Categoria
	bodyBytes := c.Ctx.Input.RequestBody
	if len(bodyBytes) == 0 && c.Ctx.Request.Body != nil {
		_ = json.NewDecoder(c.Ctx.Request.Body).Decode(&categoria)
	} else {
		_ = json.Unmarshal(bodyBytes, &categoria)
	}

	categoria.ID = id
	categoria.Nombre = strings.TrimSpace(categoria.Nombre)

	jsonBody, err := json.Marshal(categoria)
	if err != nil {
		c.Ctx.Output.SetStatus(http.StatusBadRequest)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  http.StatusBadRequest,
			"Message": "Error al serializar la categoría para actualizar",
		}
		_ = c.ServeJSON()
		return
	}

	req, err := http.NewRequest(http.MethodPut, crudCategoriaURL+"/"+idParam, bytes.NewBuffer(jsonBody))
	if err != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  http.StatusInternalServerError,
			"Message": "Error al preparar la petición PUT hacia el CRUD",
		}
		_ = c.ServeJSON()
		return
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  http.StatusInternalServerError,
			"Message": "Error de conexión con el servicio CRUD backend (puerto 8081)",
		}
		_ = c.ServeJSON()
		return
	}
	defer resp.Body.Close()

	var resultado interface{}
	_ = json.NewDecoder(resp.Body).Decode(&resultado)

	if resp.StatusCode != http.StatusOK {
		c.Ctx.Output.SetStatus(resp.StatusCode)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  resp.StatusCode,
			"Message": "No se pudo actualizar la categoría",
			"Error":   resultado,
		}
		_ = c.ServeJSON()
		return
	}

	c.Ctx.Output.SetStatus(http.StatusOK)
	c.Data["json"] = map[string]interface{}{
		"Success": true,
		"Status":  http.StatusOK,
		"Message": "Categoría actualizada correctamente",
		"Data":    resultado,
	}
	_ = c.ServeJSON()
}

// Delete elimina una categoría por ID.
func (c *CategoriasController) Delete() {
	idParam := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idParam)
	if err != nil || id <= 0 {
		c.Ctx.Output.SetStatus(http.StatusBadRequest)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  http.StatusBadRequest,
			"Message": "ID ingresado es inválido",
		}
		_ = c.ServeJSON()
		return
	}

	req, err := http.NewRequest(http.MethodDelete, crudCategoriaURL+"/"+idParam, nil)
	if err != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  http.StatusInternalServerError,
			"Message": "Error al preparar la petición DELETE",
		}
		_ = c.ServeJSON()
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  http.StatusInternalServerError,
			"Message": "Error de conexión con el servicio CRUD backend (puerto 8081)",
		}
		_ = c.ServeJSON()
		return
	}
	defer resp.Body.Close()

	var resultado interface{}
	_ = json.NewDecoder(resp.Body).Decode(&resultado)

	if resp.StatusCode != http.StatusOK {
		c.Ctx.Output.SetStatus(resp.StatusCode)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  resp.StatusCode,
			"Message": "No se pudo eliminar la categoría",
			"Error":   resultado,
		}
		_ = c.ServeJSON()
		return
	}

	c.Ctx.Output.SetStatus(http.StatusOK)
	c.Data["json"] = map[string]interface{}{
		"Success": true,
		"Status":  http.StatusOK,
		"Message": "Categoría eliminada correctamente",
	}
	_ = c.ServeJSON()
}