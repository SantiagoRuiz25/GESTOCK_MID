package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"GESTOCK_MID/models"

	beego "github.com/beego/beego/v2/server/web"
)

type MantenimientoController struct {
	beego.Controller
}

// GET: Listar mantenimientos
func (c *MantenimientoController) GetAll() {
	mantenimientos, err := models.GetAllMantenimientos()
	if err != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.Data["json"] = map[string]string{"error": "Error al obtener los mantenimientos"}
		c.ServeJSON()
		return
	}

	c.Ctx.Output.SetStatus(http.StatusOK)
	c.Data["json"] = mantenimientos
	c.ServeJSON()
}

// GET: Obtener mantenimiento por ID
func (c *MantenimientoController) GetById() {
	idParam := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idParam)
	if err != nil || id <= 0 {
		c.Ctx.Output.SetStatus(http.StatusBadRequest)
		c.Data["json"] = map[string]string{"error": "ID de mantenimiento inválido"}
		c.ServeJSON()
		return
	}

	mantenimientos, err := models.GetAllMantenimientos()
	if err != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.Data["json"] = map[string]string{"error": "Error al obtener los mantenimientos"}
		c.ServeJSON()
		return
	}

	for _, m := range mantenimientos {
		if m.ID == id {
			c.Ctx.Output.SetStatus(http.StatusOK)
			c.Data["json"] = m
			c.ServeJSON()
			return
		}
	}

	c.Ctx.Output.SetStatus(http.StatusNotFound)
	c.Data["json"] = map[string]string{"error": "Mantenimiento no encontrado"}
	c.ServeJSON()
}

// POST: Crear mantenimiento
func (c *MantenimientoController) Post() {
	var nuevoMantenimiento models.Mantenimiento
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &nuevoMantenimiento); err != nil {
		c.Ctx.Output.SetStatus(http.StatusBadRequest)
		c.Data["json"] = map[string]string{"error": "Datos inválidos"}
		c.ServeJSON()
		return
	}

	if err := nuevoMantenimiento.Create(); err != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.Data["json"] = map[string]string{"error": "Error al guardar el mantenimiento"}
		c.ServeJSON()
		return
	}

	c.Ctx.Output.SetStatus(http.StatusCreated)
	c.Data["json"] = nuevoMantenimiento
	c.ServeJSON()
}

// PUT: Actualizar mantenimiento
func (c *MantenimientoController) Put() {
	var mantenimiento models.Mantenimiento
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &mantenimiento); err != nil {
		c.Ctx.Output.SetStatus(http.StatusBadRequest)
		c.Data["json"] = map[string]string{"error": "Datos inválidos"}
		c.ServeJSON()
		return
	}

	if err := models.UpdateMantenimiento(&mantenimiento); err != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.Data["json"] = map[string]string{"error": "Error al actualizar el mantenimiento"}
		c.ServeJSON()
		return
	}

	c.Ctx.Output.SetStatus(http.StatusOK)
	c.Data["json"] = mantenimiento
	c.ServeJSON()
}