package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"GESTOCK_MID/models"

	beego "github.com/beego/beego/v2/server/web"
)

type OperacionController struct {
	beego.Controller
}

func (c *OperacionController) GetAll() {
	operaciones, err := models.GetAllOperaciones()
	if err != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.Data["json"] = map[string]string{"error": err.Error()}
		c.ServeJSON()
		return
	}

	c.Ctx.Output.SetStatus(http.StatusOK)
	c.Data["json"] = operaciones
	c.ServeJSON()
}

func (c *OperacionController) GetById() {
	id := c.Ctx.Input.Param(":id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		c.Ctx.Output.SetStatus(http.StatusBadRequest)
		c.Data["json"] = map[string]string{"error": "Invalid operation ID"}
		c.ServeJSON()
		return
	}

	// Si tienes models.GetOperacionById(intId), reemplaza esto por:
	// operacion, err := models.GetOperacionById(intId)
	operaciones, err := models.GetAllOperaciones()
	if err != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.Data["json"] = map[string]string{"error": err.Error()}
		c.ServeJSON()
		return
	}

	for _, op := range operaciones {
		if op.IdOperacion == intId {
			c.Ctx.Output.SetStatus(http.StatusOK)
			c.Data["json"] = op
			c.ServeJSON()
			return
		}
	}

	c.Ctx.Output.SetStatus(http.StatusNotFound)
	c.Data["json"] = map[string]string{"error": "Operation not found"}
	c.ServeJSON()
}

func (c *OperacionController) Post() {
	var nuevaOperacion models.Operacion
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &nuevaOperacion); err != nil {
		c.Ctx.Output.SetStatus(http.StatusBadRequest)
		c.Data["json"] = map[string]string{"error": err.Error()}
		c.ServeJSON()
		return
	}

	if err := models.CreateOperacion(&nuevaOperacion); err != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.Data["json"] = map[string]string{"error": err.Error()}
		c.ServeJSON()
		return
	}

	c.Ctx.Output.SetStatus(http.StatusCreated)
	c.Data["json"] = nuevaOperacion
	c.ServeJSON()
}

func (c *OperacionController) Delete() {
	id := c.Ctx.Input.Param(":id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		c.Ctx.Output.SetStatus(http.StatusBadRequest)
		c.Data["json"] = map[string]string{"error": "Invalid operation ID"}
		c.ServeJSON()
		return
	}

	if err := models.DeleteOperacion(intId); err != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.Data["json"] = map[string]string{"error": err.Error()}
		c.ServeJSON()
		return
	}

	c.Ctx.Output.SetStatus(http.StatusOK)
	c.Data["json"] = map[string]string{"message": "Operation deleted successfully"}
	c.ServeJSON()
}