package controllers

import "database/sql"

import "log"

import "strings"
import "github.com/lib/pq"

import (
	"GESTOCK_MID/models"
	"encoding/json"
	"strconv"

	"github.com/beego/beego/v2/server/web"
)

type ProductosController struct {
	web.Controller
}

func (c *ProductosController) GetAll() {
	productos, err := models.ListarProductos(c.Ctx.Request.Context())
	if err != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  500,
			"Message": "No se pudieron consultar los productos",
		}
		c.ServeJSON()
		return
	}

	c.Data["json"] = map[string]interface{}{
		"Success": true,
		"Status":  200,
		"Message": "Productos obtenidos con éxito",
		"Data":    productos,
	}
	c.ServeJSON()
}

func (c *ProductosController) GetById() {
	id, err := strconv.ParseInt(c.Ctx.Input.Param(":id"), 10, 32)
	if err != nil || id <= 0 {
		c.responderProducto(400, "El ID debe ser un entero positivo", nil)
		return
	}

	producto, err := models.BuscarProductoPorID(c.Ctx.Request.Context(), id)
	switch {
	case err == sql.ErrNoRows:
		c.responderProducto(404, "Producto no encontrado", nil)
	case err != nil:
		log.Printf("Error al consultar producto %d: %v", id, err)
		c.responderProducto(500, "No se pudo consultar el producto", nil)
	default:
		c.responderProducto(200, "Producto obtenido con éxito", &producto)
	}
}

func (c *ProductosController) responderProducto(estado int, mensaje string, producto *models.Producto) {
	respuesta := map[string]interface{}{
		"Success": estado >= 200 && estado < 300,
		"Status":  estado,
		"Message": mensaje,
	}
	if producto != nil {
		respuesta["Data"] = producto
	}

	c.Ctx.Output.SetStatus(estado)
	c.Data["json"] = respuesta
	c.ServeJSON()
}

func (c *ProductosController) Post() {
	var producto models.Producto
	var estado int
	var mensaje string

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &producto); err != nil {
		estado, mensaje = 400, "El JSON recibido es inválido"
	} else {
		producto.Codigo = strings.TrimSpace(producto.Codigo)
		producto.Nombre = strings.TrimSpace(producto.Nombre)

		switch {
		case producto.Codigo == "" || producto.Nombre == "":
			estado, mensaje = 400, "El código y el nombre son obligatorios"
		case producto.Id_categoria <= 0 || producto.Id_bodega <= 0:
			estado, mensaje = 400, "La categoría y la bodega deben tener un ID positivo"
		case producto.Precio < 0 || producto.Precio > 9999999999.99 || producto.Stock < 0:
			estado, mensaje = 400, "El precio o el stock están fuera del rango admitido"
		default:
			err := models.CrearProducto(c.Ctx.Request.Context(), &producto)
			if err == nil {
				estado, mensaje = 201, "Producto creado con éxito"
			} else if fallo, ok := err.(*pq.Error); ok && fallo.Code == "23505" {
				estado, mensaje = 409, "Ya existe un producto con ese código"
			} else {
				log.Printf("Error al insertar producto: %T: %v", err, err)
				estado, mensaje = 500, "No se pudo guardar el producto"
			}
		}
	}

	respuesta := map[string]interface{}{
		"Success": estado == 201,
		"Status":  estado,
		"Message": mensaje,
	}
	if estado == 201 {
		respuesta["Data"] = producto
	}

	c.Ctx.Output.SetStatus(estado)
	c.Data["json"] = respuesta
	c.ServeJSON()
}

func (c *ProductosController) Put() {
	id, err := strconv.ParseInt(c.Ctx.Input.Param(":id"), 10, 32)
	if err != nil || id <= 0 {
		c.responderProducto(400, "El ID debe ser un entero positivo", nil)
		return
	}

	var producto models.Producto
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &producto); err != nil {
		c.responderProducto(400, "El JSON recibido es inválido", nil)
		return
	}

	producto.ID = int(id)
	producto.Codigo = strings.TrimSpace(producto.Codigo)
	producto.Nombre = strings.TrimSpace(producto.Nombre)

	switch {
	case producto.Codigo == "" || producto.Nombre == "":
		c.responderProducto(400, "El código y el nombre son obligatorios", nil)
		return
	case producto.Id_categoria <= 0 || producto.Id_bodega <= 0:
		c.responderProducto(400, "La categoría y la bodega deben tener un ID positivo", nil)
		return
	case producto.Precio < 0 || producto.Precio > 9999999999.99 || producto.Stock < 0:
		c.responderProducto(400, "El precio o el stock están fuera del rango admitido", nil)
		return
	}

	err = models.ActualizarProducto(c.Ctx.Request.Context(), &producto)
	switch fallo, esErrorPG := err.(*pq.Error); {
	case err == nil:
		c.responderProducto(200, "Producto actualizado con éxito", &producto)
	case err == sql.ErrNoRows:
		c.responderProducto(404, "Producto no encontrado", nil)
	case esErrorPG && fallo.Code == "23505":
		c.responderProducto(409, "Ya existe un producto con ese código", nil)
	default:
		log.Printf("Error al actualizar producto %d: %v", id, err)
		c.responderProducto(500, "No se pudo actualizar el producto", nil)
	}
}

func (c *ProductosController) Delete() {
	id, err := strconv.ParseInt(c.Ctx.Input.Param(":id"), 10, 32)
	if err != nil || id <= 0 {
		c.responderProducto(400, "El ID debe ser un entero positivo", nil)
		return
	}

	err = models.EliminarProducto(c.Ctx.Request.Context(), id)
	switch {
	case err == sql.ErrNoRows:
		c.responderProducto(404, "Producto no encontrado", nil)
	case err != nil:
		log.Printf("Error al eliminar producto %d: %v", id, err)
		c.responderProducto(500, "No se pudo eliminar el producto", nil)
	default:
		c.responderProducto(200, "Producto eliminado con éxito", nil)
	}
}
