package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"GESTOCK_MID/models"

	beego "github.com/beego/beego/v2/server/web"
)

type IncidenciaController struct {
	beego.Controller
}

func (c *IncidenciaController) GetAll() {
	incidencias, err := models.GetAllIncidencias()
	if err != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.Data["json"] = map[string]string{"error": err.Error()}
		c.ServeJSON()
		return
	}

	c.Ctx.Output.SetStatus(http.StatusOK)
	c.Data["json"] = incidencias
	c.ServeJSON()
}

func (c *IncidenciaController) GetById() {
	idParam := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idParam)
	if err != nil || id <= 0 {
		c.Ctx.Output.SetStatus(http.StatusBadRequest)
		c.Data["json"] = map[string]string{"error": "ID de incidencia inválido"}
		c.ServeJSON()
		return
	}

	incidencias, err := models.GetAllIncidencias()
	if err != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.Data["json"] = map[string]string{"error": err.Error()}
		c.ServeJSON()
		return
	}

	for _, inc := range incidencias {
		if inc.IdIncidencia == id {
			c.Ctx.Output.SetStatus(http.StatusOK)
			c.Data["json"] = inc
			c.ServeJSON()
			return
		}
	}

	c.Ctx.Output.SetStatus(http.StatusNotFound)
	c.Data["json"] = map[string]string{"error": "Incidencia no encontrada"}
	c.ServeJSON()
}

func (c *IncidenciaController) Post() {
	var nuevaIncidencia models.Incidencia
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &nuevaIncidencia); err != nil {
		c.Ctx.Output.SetStatus(http.StatusBadRequest)
		c.Data["json"] = map[string]string{"error": err.Error()}
		c.ServeJSON()
		return
	}

	if err := nuevaIncidencia.Create(); err != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.Data["json"] = map[string]string{"error": err.Error()}
		c.ServeJSON()
		return
	}

	c.Ctx.Output.SetStatus(http.StatusCreated)
	c.Data["json"] = nuevaIncidencia
	c.ServeJSON()
}