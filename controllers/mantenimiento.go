package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"GESTOCK_MID/models"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
)

type MantenimientoController struct {
	beego.Controller
}

// GetAll obtiene todos los mantenimientos.
func (c *MantenimientoController) GetAll() {
	o := orm.NewOrm()
	var mantenimientos []models.Mantenimiento
	_, err := o.QueryTable(new(models.Mantenimiento)).All(&mantenimientos)
	if err != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  http.StatusInternalServerError,
			"Message": "Error al consultar los mantenimientos",
		}
		_ = c.ServeJSON()
		return
	}

	c.Ctx.Output.SetStatus(http.StatusOK)
	c.Data["json"] = map[string]interface{}{
		"Success": true,
		"Status":  http.StatusOK,
		"Message": "Mantenimientos obtenidos con éxito",
		"Data":    mantenimientos,
	}
	_ = c.ServeJSON()
}

// GetById obtiene un mantenimiento por su ID.
func (c *MantenimientoController) GetById() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		c.Ctx.Output.SetStatus(http.StatusBadRequest)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  http.StatusBadRequest,
			"Message": "ID de mantenimiento inválido",
		}
		_ = c.ServeJSON()
		return
	}

	o := orm.NewOrm()
	mantenimiento := &models.Mantenimiento{IdMantenimiento: id}
	if err := o.Read(mantenimiento); err != nil {
		c.Ctx.Output.SetStatus(http.StatusNotFound)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  http.StatusNotFound,
			"Message": "Mantenimiento no encontrado",
		}
		_ = c.ServeJSON()
		return
	}

	c.Ctx.Output.SetStatus(http.StatusOK)
	c.Data["json"] = map[string]interface{}{
		"Success": true,
		"Status":  http.StatusOK,
		"Message": "Mantenimiento obtenido con éxito",
		"Data":    mantenimiento,
	}
	_ = c.ServeJSON()
}

// Post crea un nuevo mantenimiento.
func (c *MantenimientoController) Post() {
	var mantenimiento models.Mantenimiento
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &mantenimiento); err != nil {
		c.Ctx.Output.SetStatus(http.StatusBadRequest)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  http.StatusBadRequest,
			"Message": "El formato del JSON recibido es inválido",
		}
		_ = c.ServeJSON()
		return
	}

	o := orm.NewOrm()
	if _, err := o.Insert(&mantenimiento); err != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  http.StatusInternalServerError,
			"Message": "Error al registrar el mantenimiento",
		}
		_ = c.ServeJSON()
		return
	}

	c.Ctx.Output.SetStatus(http.StatusCreated)
	c.Data["json"] = map[string]interface{}{
		"Success": true,
		"Status":  http.StatusCreated,
		"Message": "Mantenimiento registrado con éxito",
		"Data":    mantenimiento,
	}
	_ = c.ServeJSON()
}

// Put actualiza un mantenimiento existente por su ID.
func (c *MantenimientoController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		c.Ctx.Output.SetStatus(http.StatusBadRequest)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  http.StatusBadRequest,
			"Message": "ID de mantenimiento inválido",
		}
		_ = c.ServeJSON()
		return
	}

	var mantenimiento models.Mantenimiento
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &mantenimiento); err != nil {
		c.Ctx.Output.SetStatus(http.StatusBadRequest)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  http.StatusBadRequest,
			"Message": "El formato del JSON recibido es inválido",
		}
		_ = c.ServeJSON()
		return
	}

	mantenimiento.IdMantenimiento = id
	o := orm.NewOrm()
	if _, err := o.Update(&mantenimiento); err != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  http.StatusInternalServerError,
			"Message": "Error al actualizar el mantenimiento",
		}
		_ = c.ServeJSON()
		return
	}

	c.Ctx.Output.SetStatus(http.StatusOK)
	c.Data["json"] = map[string]interface{}{
		"Success": true,
		"Status":  http.StatusOK,
		"Message": "Mantenimiento actualizado con éxito",
		"Data":    mantenimiento,
	}
	_ = c.ServeJSON()
}

// Delete elimina un mantenimiento por su ID.
func (c *MantenimientoController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		c.Ctx.Output.SetStatus(http.StatusBadRequest)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  http.StatusBadRequest,
			"Message": "ID de mantenimiento inválido",
		}
		_ = c.ServeJSON()
		return
	}

	o := orm.NewOrm()
	if _, err := o.Delete(&models.Mantenimiento{IdMantenimiento: id}); err != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  http.StatusInternalServerError,
			"Message": "Error al eliminar el mantenimiento",
		}
		_ = c.ServeJSON()
		return
	}

	c.Ctx.Output.SetStatus(http.StatusOK)
	c.Data["json"] = map[string]interface{}{
		"Success": true,
		"Status":  http.StatusOK,
		"Message": "Mantenimiento eliminado con éxito",
	}
	_ = c.ServeJSON()
}