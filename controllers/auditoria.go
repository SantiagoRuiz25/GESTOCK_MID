package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"GESTOCK_MID/models"

	beego "github.com/beego/beego/v2/server/web"
)

type AuditoriaController struct {
	beego.Controller
}

// GetAll busca todas las auditorías
func (c *AuditoriaController) GetAll() {
	auditorias, err := models.GetAllAuditorias()
	if err != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  http.StatusInternalServerError,
			"Message": err.Error(),
		}
		_ = c.ServeJSON()
		return
	}

	c.Ctx.Output.SetStatus(http.StatusOK)
	c.Data["json"] = map[string]interface{}{
		"Success": true,
		"Status":  http.StatusOK,
		"Message": "Auditorías obtenidas con éxito",
		"Data":    auditorias,
	}
	_ = c.ServeJSON()
}

// GetById busca una auditoría específica por su ID
func (c *AuditoriaController) GetById() {
	idParam := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idParam)
	if err != nil || id <= 0 {
		c.Ctx.Output.SetStatus(http.StatusBadRequest)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  http.StatusBadRequest,
			"Message": "ID inválido",
		}
		_ = c.ServeJSON()
		return
	}

	auditorias, err := models.GetAllAuditorias()
	if err != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  http.StatusInternalServerError,
			"Message": err.Error(),
		}
		_ = c.ServeJSON()
		return
	}

	for _, a := range auditorias {
		if a.IdAuditoria == id {
			c.Ctx.Output.SetStatus(http.StatusOK)
			c.Data["json"] = map[string]interface{}{
				"Success": true,
				"Status":  http.StatusOK,
				"Message": "Auditoría obtenida con éxito",
				"Data":    a,
			}
			_ = c.ServeJSON()
			return
		}
	}

	c.Ctx.Output.SetStatus(http.StatusNotFound)
	c.Data["json"] = map[string]interface{}{
		"Success": false,
		"Status":  http.StatusNotFound,
		"Message": "Auditoría no encontrada",
	}
	_ = c.ServeJSON()
}

// Post crea una nueva auditoría
func (c *AuditoriaController) Post() {
	var auditoria models.Auditoria
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &auditoria); err != nil {
		c.Ctx.Output.SetStatus(http.StatusBadRequest)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  http.StatusBadRequest,
			"Message": err.Error(),
		}
		_ = c.ServeJSON()
		return
	}

	if err := auditoria.Create(); err != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  http.StatusInternalServerError,
			"Message": err.Error(),
		}
		_ = c.ServeJSON()
		return
	}

	c.Ctx.Output.SetStatus(http.StatusCreated)
	c.Data["json"] = map[string]interface{}{
		"Success": true,
		"Status":  http.StatusCreated,
		"Message": "Auditoría creada correctamente",
		"Data":    auditoria,
	}
	_ = c.ServeJSON()
}

// Put actualiza una auditoría existente por su ID
func (c *AuditoriaController) Put() {
	idParam := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idParam)
	if err != nil || id <= 0 {
		c.Ctx.Output.SetStatus(http.StatusBadRequest)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  http.StatusBadRequest,
			"Message": "ID inválido",
		}
		_ = c.ServeJSON()
		return
	}

	var auditoriaActualizada models.Auditoria
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &auditoriaActualizada); err != nil {
		c.Ctx.Output.SetStatus(http.StatusBadRequest)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  http.StatusBadRequest,
			"Message": err.Error(),
		}
		_ = c.ServeJSON()
		return
	}

	if err := auditoriaActualizada.Update(id); err != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  http.StatusInternalServerError,
			"Message": err.Error(),
		}
		_ = c.ServeJSON()
		return
	}

	c.Ctx.Output.SetStatus(http.StatusOK)
	c.Data["json"] = map[string]interface{}{
		"Success": true,
		"Status":  http.StatusOK,
		"Message": "Auditoría actualizada correctamente",
	}
	_ = c.ServeJSON()
}

// Delete elimina una auditoría por su ID
func (c *AuditoriaController) Delete() {
	idParam := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idParam)
	if err != nil || id <= 0 {
		c.Ctx.Output.SetStatus(http.StatusBadRequest)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  http.StatusBadRequest,
			"Message": "ID inválido",
		}
		_ = c.ServeJSON()
		return
	}

	if err := models.DeleteAuditoria(id); err != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  http.StatusInternalServerError,
			"Message": err.Error(),
		}
		_ = c.ServeJSON()
		return
	}

	c.Ctx.Output.SetStatus(http.StatusOK)
	c.Data["json"] = map[string]interface{}{
		"Success": true,
		"Status":  http.StatusOK,
		"Message": "Auditoría eliminada correctamente",
	}
	_ = c.ServeJSON()
}