package controllers

import (
	_ "github.com/beego/beego/v2/server/web"
)

type BodegasController struct {
	web.Controller
}

func (c *BodegasController) GetById() {
	id := c.Ctx.Input.Param(":id")

	if id == "" {
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  400,
			"Message": "El ID es requerido",
		}
		c.ServeJSON()
		return
	}
}


