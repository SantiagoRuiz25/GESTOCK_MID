package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"

	"GESTOCK_MID/models"

	"github.com/beego/beego/v2/server/web"
)

type ProductosController struct {
	web.Controller
}

// 1. GET ALL: Consultar todos los productos
func (c *ProductosController) GetAll() {
	respuesta, errCrud := http.Get("http://localhost:8081/v1/producto")
	if errCrud != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  500,
			"Message": "Error de conexión con el servicio CRUD",
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
			"Message": "No se encontraron productos registrados",
		}
		c.ServeJSON()
		return
	}

	var listaProductos []models.Producto
	json.NewDecoder(respuesta.Body).Decode(&listaProductos)

	c.Ctx.Output.SetStatus(200)
	c.Data["json"] = map[string]interface{}{
		"Success": true,
		"Status":  200,
		"Message": "Lista de productos obtenida con éxito",
		"Data":    listaProductos,
	}
	c.ServeJSON()
}

// 2. GET BY ID: Consultar un producto por ID
func (c *ProductosController) GetById() {
	idProducto := c.Ctx.Input.Param(":id")

	if idProducto == "" {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  400,
			"Message": "El ID es requerido",
		}
		c.ServeJSON()
		return
	}

	id, err := strconv.Atoi(idProducto)
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

	respuesta, errCrud := http.Get("http://localhost:8081/v1/producto/" + strconv.Itoa(id))
	if errCrud != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  500,
			"Message": "Error de conexión con el servicio CRUD",
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
			"Message": "El producto solicitado no existe",
		}
		c.ServeJSON()
		return
	}

	var producto models.Producto
	json.NewDecoder(respuesta.Body).Decode(&producto)

	c.Ctx.Output.SetStatus(200)
	c.Data["json"] = map[string]interface{}{
		"Success": true,
		"Status":  200,
		"Message": "Producto obtenido con éxito",
		"Data":    producto,
	}
	c.ServeJSON()
}

// 3. POST: Crear nuevo producto
func (c *ProductosController) Post() {
	var nuevoProducto models.Producto

	err := json.Unmarshal(c.Ctx.Input.RequestBody, &nuevoProducto)
	if err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  400,
			"Message": "El formato del JSON recibido es invalido",
		}
		c.ServeJSON()
		return
	}

	if nuevoProducto.Nombre == "" || nuevoProducto.Codigo == "" {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  400,
			"Message": "Campos obligatorios incompletos",
		}
		c.ServeJSON()
		return
	}

	respuesta, errCrud := http.Post(
		"http://localhost:8081/v1/producto",
		"application/json",
		bytes.NewBuffer(c.Ctx.Input.RequestBody),
	)
	if errCrud != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  500,
			"Message": "Error de conexión con el servicio CRUD",
		}
		c.ServeJSON()
		return
	}
	defer respuesta.Body.Close()

	var crearProducto models.Producto
	json.NewDecoder(respuesta.Body).Decode(&crearProducto)

	c.Ctx.Output.SetStatus(201)
	c.Data["json"] = map[string]interface{}{
		"Success": true,
		"Status":  201,
		"Message": "Producto creado con éxito",
		"Data":    crearProducto,
	}
	c.ServeJSON()
}

// 4. PUT: Actualizar producto por ID
func (c *ProductosController) Put() {
	idProducto := c.Ctx.Input.Param(":id")

	if idProducto == "" {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  400,
			"Message": "El ID es requerido",
		}
		c.ServeJSON()
		return
	}

	id, err := strconv.Atoi(idProducto)
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

	client := &http.Client{}
	req, errReq := http.NewRequest("PUT", "http://localhost:8081/v1/producto/"+strconv.Itoa(id), bytes.NewBuffer(c.Ctx.Input.RequestBody))
	if errReq != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  500,
			"Message": "Error de conexión con el servicio CRUD",
		}
		c.ServeJSON()
		return
	}
	req.Header.Set("Content-Type", "application/json")

	respuesta, errCrud := client.Do(req)
	if errCrud != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  500,
			"Message": "Error de conexión con el servicio CRUD",
		}
		c.ServeJSON()
		return
	}
	defer respuesta.Body.Close()

	c.Ctx.Output.SetStatus(200)
	c.Data["json"] = map[string]interface{}{
		"Success": true,
		"Status":  200,
		"Message": "Producto actualizado con éxito",
	}
	c.ServeJSON()
}

// 5. DELETE: Eliminar producto por ID
func (c *ProductosController) Delete() {
	idProducto := c.Ctx.Input.Param(":id")

	if idProducto == "" {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  400,
			"Message": "El ID es requerido",
		}
		c.ServeJSON()
		return
	}

	id, err := strconv.Atoi(idProducto)
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

	client := &http.Client{}
	req, errReq := http.NewRequest("DELETE", "http://localhost:8081/v1/producto/"+strconv.Itoa(id), nil)
	if errReq != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  500,
			"Message": "Error de conexión con el servicio CRUD",
		}
		c.ServeJSON()
		return
	}

	respuesta, errCrud := client.Do(req)
	if errCrud != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  500,
			"Message": "Error de conexión con el servicio CRUD",
		}
		c.ServeJSON()
		return
	}
	defer respuesta.Body.Close()

	c.Ctx.Output.SetStatus(200)
	c.Data["json"] = map[string]interface{}{
		"Success": true,
		"Status":  200,
		"Message": "Producto eliminado correctamente",
	}
	c.ServeJSON()
}