package controllers

import (
	"github.com/beego/beego/v2/server/web"
	_ "github.com/beego/beego/v2/server/web"
	"strconv"
)

type BodegasController struct {
	web.Controller
}

func (c *BodegasController) GetById() {
	idBodega := c.Ctx.Input.Param(":id")

	if idBodega == "" {
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  400,
			"Message": "El ID es requerido",
		}
		c.ServeJSON()
		return
	} 

	id, err := strconv.Atoi(idBodega)	
	if err !=nil || id <= 0 {
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status": 404,
			"Mensaje": "Id ingresado es invalido",
		}
		c.ServeJSON()
		return

	}

	func (c *BodegasController) Post() {
	var nuevaBodega models.Bodegas

	err := json.Unmarshal(c.Ctx.Input.RequestBody, &nuevaBodega)
	if err != nil {
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  400,
			"Message": "El formato del JSON recibido es invalido",
		}
		c.ServeJSON()
		return
	}

	if nuevaBodega.Nombre == "" || nuevaBodega.Codigo == "" {
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  400,
			"Message": "Campos obligatorios incompletos",
		}
		c.ServeJSON()
		return
	}

	c.Data["json"] = map[string]interface{}{
		"Success": true,
		"Status":  201,
		"Message": "Bodega creada exitosamente",
		"Data":    nuevaBodega,
	}
	c.ServeJSON()
}
}





