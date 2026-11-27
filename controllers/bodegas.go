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

	"github.com/beego/beego/v2/server/web"
)

const crudBodegaURL = "http://localhost:8081/v1/bodega"

type BodegasController struct {
	web.Controller
}

// BodegaController mantiene compatibilidad si alguna ruta lo referencia directamente.
type BodegaController struct {
	BodegasController
}

func (c *BodegasController) GetAll() {
	if _, _, err := services.SolicitarBodega(
		c.Ctx.Request.Context(), http.MethodGet, 0, nil,
	); err != nil {
		c.reenviarPeticionDirecta(http.MethodGet, "", nil)
		return
	}
	c.enviarBodegaAlCRUD(http.MethodGet, 0)
}

func (c *BodegasController) GetById() {
	c.enviarBodegaPorID(http.MethodGet)
}

func (c *BodegasController) Post() {
	var bodega models.Bodegas
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &bodega); err != nil {
		c.responderErrorBodega(http.StatusBadRequest, "El formato del JSON recibido es inválido: "+err.Error())
		return
	}
	if bodega.Nombre == "" || bodega.Codigo == 0 {
		c.responderErrorBodega(http.StatusBadRequest, "Los campos 'nombre' y 'codigo' son obligatorios")
		return
	}
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

func (c *BodegasController) reenviarPeticionDirecta(metodo string, sufijo string, cuerpo []byte) {
	url := crudBodegaURL + sufijo
	req, err := http.NewRequest(metodo, url, bytes.NewBuffer(cuerpo))
	if err != nil {
		c.responderErrorBodega(http.StatusInternalServerError, "Error al preparar la petición al CRUD")
		return
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.responderErrorBodega(http.StatusBadGateway, "Error de conexión con el servicio CRUD (8081)")
		return
	}
	defer resp.Body.Close()

	var resultado interface{}
	_ = json.NewDecoder(resp.Body).Decode(&resultado)
	c.Ctx.Output.SetStatus(resp.StatusCode)
	c.Data["json"] = resultado
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