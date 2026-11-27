package controllers

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	"GESTOCK_MID/models"
	"GESTOCK_MID/services"

	beego "github.com/beego/beego/v2/server/web"
)

const crudProductoURL = "http://localhost:8081/v1/productos"

type ProductosController struct {
	beego.Controller
}

func (c *ProductosController) GetAll() {
	if _, _, err := services.SolicitarProductoParaCliente(
		c.Ctx.Request.Context(), http.MethodGet, 0, nil,
	); err != nil {
		c.reenviarProductoDirecto(http.MethodGet, "", nil)
		return
	}
	c.enviarAlCRUD(http.MethodGet, 0)
}

func (c *ProductosController) GetById() {
	c.enviarProductoPorID(http.MethodGet)
}

func (c *ProductosController) Post() {
	var producto models.Producto
	bodyBytes := c.Ctx.Input.RequestBody
	var err error
	if len(bodyBytes) == 0 && c.Ctx.Request.Body != nil {
		err = json.NewDecoder(c.Ctx.Request.Body).Decode(&producto)
	} else {
		err = json.Unmarshal(bodyBytes, &producto)
	}
	if err != nil {
		c.responderError(http.StatusBadRequest, "El formato del JSON recibido es inválido: "+err.Error())
		return
	}
	if producto.Nombre == "" || producto.Codigo == "" {
		c.responderError(http.StatusBadRequest, "Los campos 'nombre' y 'codigo' son obligatorios")
		return
	}
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

func (c *ProductosController) reenviarProductoDirecto(metodo string, sufijo string, cuerpo []byte) {
	idStr := c.Ctx.Input.Param(":id")
	if idStr != "" && sufijo == "" {
		sufijo = "/" + idStr
	}

	req, err := http.NewRequest(metodo, crudProductoURL+sufijo, bytes.NewBuffer(cuerpo))
	if err != nil {
		c.responderError(http.StatusInternalServerError, "Error al preparar la petición al CRUD")
		return
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.responderError(http.StatusBadGateway, "Error de conexión con el servicio CRUD backend (puerto 8081)")
		return
	}
	defer resp.Body.Close()

	var resultado interface{}
	_ = json.NewDecoder(resp.Body).Decode(&resultado)
	c.Ctx.Output.SetStatus(resp.StatusCode)
	c.Data["json"] = resultado
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