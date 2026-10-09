package controllers

import (
	"GESTOCK_MID/models"
	"encoding/json"
	"strconv"

	"github.com/beego/beego/v2/server/web"
)

type ProductosController struct {
	web.Controller
}

func (c *ProductosController) GetById() {
	idProducto := c.Ctx.Input.Param(":id")

	if idProducto == "" {
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
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  400,
			"Message": "Id ingresado es invalido",
		}
		c.ServeJSON()
		return
	}

}

func (c *ProductosController) Post() {
	var nuevoProducto models.Producto

	err := json.Unmarshal(c.Ctx.Input.RequestBody, &nuevoProducto)
	if err != nil {
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  400,
			"Message": "El formato del JSON recibido es invalido",
		}
		c.ServeJSON()
		return
	}


if nuevoProducto.Nombre == "" || nuevoProducto.Codigo == "" {
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  400,
			"Message": "Campos obligatorios incompletos",
		}
		c.ServeJSON()
		return
	}

}