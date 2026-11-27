package controllers

import (
	"context"
	"errors"
	"log"
	"net/http"
	"strconv"

	"GESTOCK_MID/services"

	beego "github.com/beego/beego/v2/server/web"
)

type EmpresasMidController struct {
	beego.Controller
}

func (c *EmpresasMidController) GetAll() {
	c.enviarCRUD(http.MethodGet, false)
}

func (c *EmpresasMidController) GetById() {
	c.enviarCRUD(http.MethodGet, true)
}

func (c *EmpresasMidController) Post() {
	c.enviarCRUD(http.MethodPost, false)
}

func (c *EmpresasMidController) Put() {
	c.enviarCRUD(http.MethodPut, true)
}

func (c *EmpresasMidController) Delete() {
	c.enviarCRUD(http.MethodDelete, true)
}

func (c *EmpresasMidController) enviarCRUD(metodo string, conID bool) {
	var id int64
	if conID {
		valor, err := strconv.ParseInt(c.Ctx.Input.Param(":id"), 10, 32)
		if err != nil || valor <= 0 {
			c.responderErrorEmpresa(http.StatusBadRequest, "El ID de empresa debe ser un entero positivo")
			return
		}
		id = valor
	}

	estado, datos, err := services.SolicitarEmpresa(
		c.Ctx.Request.Context(), metodo, id, c.Ctx.Input.RequestBody,
	)
	if err != nil {
		log.Printf("Error al consultar el CRUD de empresas: %v", err)
		if errors.Is(err, context.DeadlineExceeded) {
			c.responderErrorEmpresa(http.StatusGatewayTimeout, "El servicio CRUD tardó demasiado en responder")
			return
		}
		c.responderErrorEmpresa(http.StatusBadGateway, "No se pudo comunicar con el servicio CRUD")
		return
	}

	c.Ctx.Output.SetStatus(estado)
	c.Data["json"] = datos
	_ = c.ServeJSON()
}

func (c *EmpresasMidController) responderErrorEmpresa(estado int, mensaje string) {
	c.Ctx.Output.SetStatus(estado)
	c.Data["json"] = map[string]interface{}{
		"Success": false,
		"Status":  estado,
		"Message": mensaje,
	}
	_ = c.ServeJSON()
}
