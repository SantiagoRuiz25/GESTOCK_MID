package controllers

import (
	"encoding/json"
	"strconv"

	"GESTOCK_MID/models"

	beego "github.com/beego/beego/v2/server/web"
)

type AuditoriaController struct {
	beego.Controller
}

// 1. GET ALL: Consultar todas las auditorías
func (c *AuditoriaController) GetAll() {
	auditorias, err := models.GetAllAuditorias()
	if err != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  500,
			"Message": "Error al obtener los registros de auditoría",
		}
		c.ServeJSON()
		return
	}

	c.Ctx.Output.SetStatus(200)
	c.Data["json"] = map[string]interface{}{
		"Success": true,
		"Status":  200,
		"Message": "Lista de auditorías obtenida con éxito",
		"Data":    auditorias,
	}
	c.ServeJSON()
}

// 2. GET BY ID: Consultar una auditoría por ID
func (c *AuditoriaController) GetById() {
	idParam := c.Ctx.Input.Param(":id")

	if idParam == "" {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  400,
			"Message": "El ID es requerido",
		}
		c.ServeJSON()
		return
	}

	id, err := strconv.Atoi(idParam)
	if err != nil || id <= 0 {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  400,
			"Message": "Id ingresado es invalido",
		}
		c.ServeJSON()
		return
	}

	// Consultar la auditoría por ID a través del modelo
	auditorias, err := models.GetAllAuditorias()
	if err != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  500,
			"Message": "Error al consultar la auditoría",
		}
		c.ServeJSON()
		return
	}

	var auditoriaEncontrada *models.Auditoria
	for _, item := range auditorias {
		if item.IdAuditoria == id {
			auditoriaEncontrada = &item
			break
		}
	}

	if auditoriaEncontrada == nil {
		c.Ctx.Output.SetStatus(404)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  404,
			"Message": "No se encontró el registro de auditoría solicitado",
		}
		c.ServeJSON()
		return
	}

	c.Ctx.Output.SetStatus(200)
	c.Data["json"] = map[string]interface{}{
		"Success": true,
		"Status":  200,
		"Message": "Auditoría obtenida con éxito",
		"Data":    auditoriaEncontrada,
	}
	c.ServeJSON()
}

// 3. POST: Crear registro de auditoría
func (c *AuditoriaController) Post() {
	var nuevaAuditoria models.Auditoria

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &nuevaAuditoria); err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  400,
			"Message": "El formato del JSON recibido es invalido",
		}
		c.ServeJSON()
		return
	}

	if err := nuevaAuditoria.Create(); err != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  500,
			"Message": "Error al guardar el registro de auditoría",
		}
		c.ServeJSON()
		return
	}

	c.Ctx.Output.SetStatus(201)
	c.Data["json"] = map[string]interface{}{
		"Success": true,
		"Status":  201,
		"Message": "Auditoría creada exitosamente",
		"Data":    nuevaAuditoria,
	}
	c.ServeJSON()
}

// 4. PUT: Actualizar auditoría por ID
func (c *AuditoriaController) Put() {
	idParam := c.Ctx.Input.Param(":id")

	if idParam == "" {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  400,
			"Message": "El ID es requerido",
		}
		c.ServeJSON()
		return
	}

	id, err := strconv.Atoi(idParam)
	if err != nil || id <= 0 {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  400,
			"Message": "Id ingresado es invalido",
		}
		c.ServeJSON()
		return
	}

	var auditoriaActualizada models.Auditoria
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &auditoriaActualizada); err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  400,
			"Message": "El formato del JSON recibido es invalido",
		}
		c.ServeJSON()
		return
	}

	if err := auditoriaActualizada.Update(id); err != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  500,
			"Message": "Error al actualizar la auditoría",
		}
		c.ServeJSON()
		return
	}

	c.Ctx.Output.SetStatus(200)
	c.Data["json"] = map[string]interface{}{
		"Success": true,
		"Status":  200,
		"Message": "Auditoría actualizada correctamente",
		"Data":    auditoriaActualizada,
	}
	c.ServeJSON()
}

// 5. DELETE: Eliminar auditoría por ID
func (c *AuditoriaController) Delete() {
	idParam := c.Ctx.Input.Param(":id")

	if idParam == "" {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  400,
			"Message": "El ID es requerido",
		}
		c.ServeJSON()
		return
	}

	id, err := strconv.Atoi(idParam)
	if err != nil || id <= 0 {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  400,
			"Message": "Id ingresado es invalido",
		}
		c.ServeJSON()
		return
	}

	if err := models.DeleteAuditoria(id); err != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  500,
			"Message": "Error al eliminar el registro de auditoría",
		}
		c.ServeJSON()
		return
	}

	c.Ctx.Output.SetStatus(200)
	c.Data["json"] = map[string]interface{}{
		"Success": true,
		"Status":  200,
		"Message": "Auditoría eliminada correctamente",
	}
	c.ServeJSON()
}