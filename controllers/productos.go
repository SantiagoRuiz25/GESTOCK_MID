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

type ProductosController struct {
	web.Controller
}

func (c *ProductosController) GetAll() {
	c.enviarAlCRUD(http.MethodGet, 0)
}

func (c *ProductosController) GetById() {
	c.enviarProductoPorID(http.MethodGet)
}

func (c *ProductosController) Post() {
	c.enviarAlCRUD(http.MethodPost, 0)
}

func (c *ProductosController) Put() {
	c.enviarProductoPorID(http.MethodPut)
}

func (c *ProductosController) Delete() {
	c.enviarProductoPorID(http.MethodDelete)
}

func (c *ProductosController) enviarProductoPorID(metodo string) {
	id, err := strconv.ParseInt(c.Ctx.Input.Param(":id"), 10, 32)
	if err != nil || id <= 0 {
		c.responderError(http.StatusBadRequest, "El ID debe ser un entero positivo")
		return
	}
	c.enviarAlCRUD(metodo, id)
}

func (c *ProductosController) enviarAlCRUD(metodo string, id int64) {
	estado, datos, err := services.SolicitarProductoParaCliente(
		c.Ctx.Request.Context(), metodo, id, c.Ctx.Input.RequestBody,
	)
	if err != nil {
		log.Printf("Error al consultar el CRUD de productos: %v", err)
		if errors.Is(err, context.DeadlineExceeded) {
			c.responderError(http.StatusGatewayTimeout, "El servicio CRUD tardó demasiado en responder")
			return
		}
		c.responderError(http.StatusBadGateway, "No se pudo comunicar con el servicio CRUD")
		return
	}

	c.Ctx.Output.SetStatus(estado)
	c.Data["json"] = datos
	_ = c.ServeJSON()
}

func (c *ProductosController) responderError(estado int, mensaje string) {
	c.Ctx.Output.SetStatus(estado)
	c.Data["json"] = map[string]interface{}{
		"Success": false,
		"Status":  estado,
		"Message": mensaje,
	}
	_ = c.ServeJSON()
}
