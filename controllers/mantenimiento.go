package controllers

import (
	"encoding/json"
	"strconv"

	"GESTOCK_MID/models"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
)

type MantenimientoController struct {
	beego.Controller
}

// GetAll: Obtener todos los mantenimientos
func (c *MantenimientoController) GetAll() {
	o := orm.NewOrm()
	var mantenimientos []models.Mantenimiento
	_, err := o.QueryTable(new(models.Mantenimiento)).All(&mantenimientos)
	if err != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]interface{}{"Success": false, "Message": "Error al consultar los mantenimientos"}
		c.ServeJSON()
		return
	}

	c.Ctx.Output.SetStatus(200)
	c.Data["json"] = map[string]interface{}{"Success": true, "Data": mantenimientos}
	c.ServeJSON()
}

// GetById: Obtener un mantenimiento por ID
func (c *MantenimientoController) GetById() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]interface{}{"Success": false, "Message": "ID inválido"}
		c.ServeJSON()
		return
	}

	o := orm.NewOrm()
	mantenimiento := &models.Mantenimiento{IdMantenimiento: id}
	err = o.Read(mantenimiento)
	if err != nil {
		c.Ctx.Output.SetStatus(404)
		c.Data["json"] = map[string]interface{}{"Success": false, "Message": "Mantenimiento no encontrado"}
		c.ServeJSON()
		return
	}

	c.Ctx.Output.SetStatus(200)
	c.Data["json"] = map[string]interface{}{"Success": true, "Data": mantenimiento}
	c.ServeJSON()
}

// Post: Crear un nuevo mantenimiento
func (c *MantenimientoController) Post() {
	var mantenimiento models.Mantenimiento
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &mantenimiento); err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]interface{}{"Success": false, "Message": "El formato del JSON recibido es inválido"}
		c.ServeJSON()
		return
	}

	o := orm.NewOrm()
	_, err := o.Insert(&mantenimiento)
	if err != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]interface{}{"Success": false, "Message": "Error al registrar el mantenimiento"}
		c.ServeJSON()
		return
	}

	c.Ctx.Output.SetStatus(201)
	c.Data["json"] = map[string]interface{}{"Success": true, "Data": mantenimiento}
	c.ServeJSON()
}