package controllers

import "database/sql"

import "strings"
import "github.com/lib/pq"

import "log"

import (
	"encoding/json"
	"net/http"
	"strconv"

	"GESTOCK_MID/models"

	beego "github.com/beego/beego/v2/server/web"
)

type BodegasController struct {
	beego.Controller
}

func (c *BodegasController) parseID() (int, bool) {
	rawID := c.Ctx.Input.Param(":id")
	if rawID == "" {
		c.Ctx.Output.SetStatus(http.StatusBadRequest)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  http.StatusBadRequest,
			"Message": "El ID es requerido",
		}
		_ = c.ServeJSON()
		return 0, false
	}

	id, err := strconv.Atoi(rawID)
	if err != nil || id <= 0 {
		c.Ctx.Output.SetStatus(http.StatusBadRequest)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  http.StatusBadRequest,
			"Message": "El ID ingresado es inválido",
		}
		_ = c.ServeJSON()
		return 0, false
	}

	return id, true
}

func (c *BodegasController) GetAll() {
	bodegas, err := models.ListarBodegas(c.Ctx.Request.Context())
	if err != nil {
		log.Printf("Error al consultar bodegas: %v", err)
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  http.StatusInternalServerError,
			"Message": "No se pudieron consultar las bodegas",
		}
		_ = c.ServeJSON()
		return
	}

	c.Data["json"] = map[string]interface{}{
		"Success": true,
		"Status":  http.StatusOK,
		"Message": "Lista de bodegas obtenida con éxito",
		"Data":    bodegas,
	}
	_ = c.ServeJSON()
}

func (c *BodegasController) GetById() {
	id, ok := c.parseID()
	if !ok {
		return
	}
	if id > 2147483647 {
		c.responderBodega(http.StatusBadRequest, "El ID ingresado es inválido", nil)
		return
	}

	bodega, err := models.BuscarBodegaPorID(c.Ctx.Request.Context(), id)
	switch {
	case err == sql.ErrNoRows:
		c.responderBodega(http.StatusNotFound, "La bodega solicitada no existe", nil)
	case err != nil:
		log.Printf("Error al consultar bodega %d: %v", id, err)
		c.responderBodega(http.StatusInternalServerError, "No se pudo consultar la bodega", nil)
	default:
		c.responderBodega(http.StatusOK, "Bodega obtenida con éxito", &bodega)
	}
}

func (c *BodegasController) Post() {
	var bodega models.BodegaPostgres
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &bodega); err != nil {
		c.responderBodega(http.StatusBadRequest, "El JSON recibido es inválido", nil)
		return
	}

	for _, campo := range []*string{
		&bodega.Codigo, &bodega.Nombre, &bodega.Direccion,
		&bodega.Ciudad, &bodega.Telefono, &bodega.Responsable,
	} {
		*campo = strings.TrimSpace(*campo)
		if *campo == "" {
			c.responderBodega(http.StatusBadRequest, "Completa los campos obligatorios de la bodega", nil)
			return
		}
	}
	if bodega.Capacidad <= 0 || bodega.Capacidad > 2147483647 {
		c.responderBodega(http.StatusBadRequest, "La capacidad máxima está fuera del rango admitido", nil)
		return
	}

	err := models.CrearBodega(c.Ctx.Request.Context(), &bodega)
	switch fallo, esErrorPG := err.(*pq.Error); {
	case err == nil:
		c.responderBodega(http.StatusCreated, "Bodega creada con éxito", &bodega)
	case esErrorPG && fallo.Code == "23505":
		c.responderBodega(http.StatusConflict, "Ya existe una bodega con ese código", nil)
	default:
		log.Printf("Error al crear bodega: %v", err)
		c.responderBodega(http.StatusInternalServerError, "No se pudo guardar la bodega", nil)
	}
}

func (c *BodegasController) responderBodega(estado int, mensaje string, bodega *models.BodegaPostgres) {
	respuesta := map[string]interface{}{
		"Success": estado >= 200 && estado < 300,
		"Status":  estado,
		"Message": mensaje,
	}
	if bodega != nil {
		respuesta["Data"] = bodega
	}

	c.Ctx.Output.SetStatus(estado)
	c.Data["json"] = respuesta
	_ = c.ServeJSON()
}

func (c *BodegasController) Put() {
	id, ok := c.parseID()
	if !ok {
		return
	}
	if id > 2147483647 {
		c.responderBodega(http.StatusBadRequest, "El ID ingresado es inválido", nil)
		return
	}

	var bodega models.BodegaPostgres
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &bodega); err != nil {
		c.responderBodega(http.StatusBadRequest, "El JSON recibido es inválido", nil)
		return
	}
	bodega.ID = id

	for _, campo := range []*string{
		&bodega.Codigo, &bodega.Nombre, &bodega.Direccion,
		&bodega.Ciudad, &bodega.Telefono, &bodega.Responsable,
	} {
		*campo = strings.TrimSpace(*campo)
		if *campo == "" {
			c.responderBodega(http.StatusBadRequest, "Completa los campos obligatorios de la bodega", nil)
			return
		}
	}
	if bodega.Capacidad <= 0 || bodega.Capacidad > 2147483647 {
		c.responderBodega(http.StatusBadRequest, "La capacidad máxima está fuera del rango admitido", nil)
		return
	}

	err := models.ActualizarBodega(c.Ctx.Request.Context(), &bodega)
	switch fallo, esErrorPG := err.(*pq.Error); {
	case err == nil:
		c.responderBodega(http.StatusOK, "Bodega actualizada con éxito", &bodega)
	case err == sql.ErrNoRows:
		c.responderBodega(http.StatusNotFound, "La bodega solicitada no existe", nil)
	case esErrorPG && fallo.Code == "23505":
		c.responderBodega(http.StatusConflict, "Ya existe una bodega con ese código", nil)
	default:
		log.Printf("Error al actualizar bodega %d: %v", id, err)
		c.responderBodega(http.StatusInternalServerError, "No se pudo actualizar la bodega", nil)
	}
}

func (c *BodegasController) Delete() {
	id, ok := c.parseID()
	if !ok {
		return
	}
	if id > 2147483647 {
		c.responderBodega(http.StatusBadRequest, "El ID ingresado es inválido", nil)
		return
	}

	err := models.DesactivarBodega(c.Ctx.Request.Context(), id)
	switch {
	case err == sql.ErrNoRows:
		c.responderBodega(http.StatusNotFound, "No existe una bodega activa con ese ID", nil)
	case err != nil:
		log.Printf("Error al desactivar bodega %d: %v", id, err)
		c.responderBodega(http.StatusInternalServerError, "No se pudo desactivar la bodega", nil)
	default:
		c.responderBodega(http.StatusOK, "Bodega desactivada con éxito", nil)
	}
}
