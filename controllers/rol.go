package controllers

import (
	"encoding/json"
	"GESTOCK_MID/services"
	"github.com/beego/beego/v2/server/web"
)

type RolMidController struct {
	web.Controller
}

func (c *RolMidController) Post() {
	var rol map[string]interface{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &rol); err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]interface{}{"success": false, "status": 400, "message": "JSON mal formado", "error": err.Error()}
		c.ServeJSON()
		return
	}

	resultado, err := services.CrearRol(rol)
	if err != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]interface{}{"success": false, "status": 500, "message": "Error al conectar con el CRUD", "error": err.Error()}
		c.ServeJSON()
		return
	}

	c.Data["json"] = resultado
	c.ServeJSON()
}

func (c *RolMidController) GetAll() {
	resultado, err := services.ObtenerRoles()
	if err != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]interface{}{"success": false, "status": 500, "message": "Error al obtener los roles", "error": err.Error()}
		c.ServeJSON()
		return
	}
	c.Data["json"] = resultado
	c.ServeJSON()
}

func (c *RolMidController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	resultado, err := services.ObtenerRolPorId(idStr)
	if err != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]interface{}{"success": false, "status": 500, "message": "Error al obtener el rol", "error": err.Error()}
		c.ServeJSON()
		return
	}
	c.Data["json"] = resultado
	c.ServeJSON()
}

func (c *RolMidController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	resultado, err := services.EliminarRol(idStr)
	if err != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]interface{}{"success": false, "status": 500, "message": "Error al eliminar el rol", "error": err.Error()}
		c.ServeJSON()
		return
	}
	c.Data["json"] = resultado
	c.ServeJSON()
}