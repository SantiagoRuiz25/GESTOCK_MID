package controllers

import (
	"GESTOCK_MID/services"
	"github.com/beego/beego/v2/server/web"
)

type SesionController struct {
	web.Controller
}

// GetAll /v1/sesiones (Ver sesiones activas)
func (c *SesionController) GetAll() {
	resultado, err := services.ObtenerSesiones()
	if err != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]interface{}{"success": false, "status": 500, "message": "Error al obtener las sesiones activas", "error": err.Error()}
		c.ServeJSON()
		return
	}
	c.Data["json"] = resultado
	c.ServeJSON()
}

// Delete /v1/sesiones/:id (Cerrar sesión / Logout)
func (c *SesionController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	resultado, err := services.CerrarSesion(idStr)
	if err != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]interface{}{"success": false, "status": 500, "message": "Error al cerrar la sesión", "error": err.Error()}
		c.ServeJSON()
		return
	}
	c.Data["json"] = resultado
	c.ServeJSON()
}