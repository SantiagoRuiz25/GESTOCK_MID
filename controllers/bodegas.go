package controllers

import (
	"context"
	"errors"
	"log"
	"net/http"
	"strconv"

	"GESTOCK_MID/services"

	"github.com/beego/beego/v2/server/web"
)

type BodegasController struct {
	web.Controller
}

func (c *BodegasController) GetAll() {
	c.enviarBodegaAlCRUD(http.MethodGet, 0)
}

func (c *BodegasController) GetById() {
	c.enviarBodegaPorID(http.MethodGet)
}

func (c *BodegasController) Post() {
	c.enviarBodegaAlCRUD(http.MethodPost, 0)
}

func (c *BodegasController) Put() {
	c.enviarBodegaPorID(http.MethodPut)
}

func (c *BodegasController) Delete() {
	c.enviarBodegaPorID(http.MethodDelete)
}

func (c *BodegasController) enviarBodegaPorID(metodo string) {
	id, err := strconv.ParseInt(c.Ctx.Input.Param(":id"), 10, 32)
	if err != nil || id <= 0 {
		c.responderErrorBodega(http.StatusBadRequest, "El ID debe ser un entero positivo")
		return
	}
	c.enviarBodegaAlCRUD(metodo, id)
}

func (c *BodegasController) enviarBodegaAlCRUD(metodo string, id int64) {
	estado, datos, err := services.SolicitarBodega(
		c.Ctx.Request.Context(), metodo, id, c.Ctx.Input.RequestBody,
	)
	if err != nil {
		log.Printf("Error al consultar el CRUD de bodegas: %v", err)
		if errors.Is(err, context.DeadlineExceeded) {
			c.responderErrorBodega(http.StatusGatewayTimeout, "El servicio CRUD tardó demasiado en responder")
			return
		}
		c.responderErrorBodega(http.StatusBadGateway, "No se pudo comunicar con el servicio CRUD")
		return
	}

	c.Ctx.Output.SetStatus(estado)
	c.Data["json"] = datos
	_ = c.ServeJSON()
}

func (c *BodegasController) responderErrorBodega(estado int, mensaje string) {
	c.Ctx.Output.SetStatus(estado)
	c.Data["json"] = map[string]interface{}{
		"Success": false,
		"Status":  estado,
		"Message": mensaje,
	}
	_ = c.ServeJSON()
}
