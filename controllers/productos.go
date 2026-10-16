package controllers

import (
	"strconv"

	"GESTOCK_MID/services"

	beego "github.com/beego/beego/v2/server/web"
)

type ProductosController struct {
	beego.Controller
}

// 1. GET ALL: Consultar todos los productos
func (c *ProductosController) GetAll() {
	lista, err := services.ObtenerTodosProductos()
	if err != nil {
		status := 500
		if err.Error() == "recurso no encontrado en el sistema" {
			status = 404
		}
		c.Ctx.Output.SetStatus(status)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  status,
			"Message": err.Error(),
		}
		c.ServeJSON()
		return
	}

	c.Ctx.Output.SetStatus(200)
	c.Data["json"] = map[string]interface{}{
		"Success": true,
		"Status":  200,
		"Message": "Lista de productos obtenida con éxito",
		"Data":    lista,
	}
	c.ServeJSON()
}

// 2. GET BY ID: Consultar producto por ID
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

	producto, err := services.ObtenerProductoPorId(id)
	if err != nil {
		status := 500
		if err.Error() == "recurso no encontrado en el sistema" {
			status = 404
		}
		c.Ctx.Output.SetStatus(status)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  status,
			"Message": err.Error(),
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

// 3. POST: Crear nuevo producto
func (c *ProductosController) Post() {
	productoCreado, err := services.CrearProducto(c.Ctx.Input.RequestBody)
	if err != nil {
		status := 400
		if err.Error() == "error de conexión con el servicio CRUD backend (puerto 8081)" {
			status = 500
		}
		c.Ctx.Output.SetStatus(status)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  status,
			"Message": err.Error(),
		}
		c.ServeJSON()
		return
	}

	c.Ctx.Output.SetStatus(201)
	c.Data["json"] = map[string]interface{}{
		"Success": true,
		"Status":  201,
		"Message": "Producto creado exitosamente",
		"Data":    productoCreado,
	}
	c.ServeJSON()
}

// 4. PUT: Actualizar producto por ID
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

	err = services.ActualizarProducto(id, c.Ctx.Input.RequestBody)
	if err != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  500,
			"Message": err.Error(),
		}
		c.ServeJSON()
		return
	}

	c.Ctx.Output.SetStatus(200)
	c.Data["json"] = map[string]interface{}{
		"Success": true,
		"Status":  200,
		"Message": "Producto actualizado correctamente",
	}
	c.ServeJSON()
}

// 5. DELETE: Eliminar producto por ID
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

	err = services.EliminarProducto(id)
	if err != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  500,
			"Message": err.Error(),
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