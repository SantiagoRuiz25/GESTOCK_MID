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
	"github.com/beego/beego/v2/server/web"
)

type CategoriasController struct {
	web.Controller
}

func (c *CategoriasController) GetAll() {
	categorias, err := models.ListarCategorias(c.Ctx.Request.Context())
	if err != nil {
		log.Printf("Error al listar categorías: %v", err)
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  http.StatusInternalServerError,
			"Message": "No se pudieron consultar las categorías",
		}
		_ = c.ServeJSON()
		return
	}

	c.Data["json"] = map[string]interface{}{
		"Success": true,
		"Status":  http.StatusOK,
		"Message": "Categorías obtenidas con éxito",
		"Data":    categorias,
	}
	_ = c.ServeJSON()
}

func (c *CategoriasController) GetById() {
	id, err := strconv.ParseInt(c.Ctx.Input.Param(":id"), 10, 32)
	if err != nil || id <= 0 {
		c.responderCategoria(http.StatusBadRequest, "El ID debe ser un entero positivo", nil)
		return
	}

	categoria, err := models.BuscarCategoriaPorID(c.Ctx.Request.Context(), id)
	switch {
	case err == sql.ErrNoRows:
		c.responderCategoria(http.StatusNotFound, "Categoría no encontrada", nil)
	case err != nil:
		log.Printf("Error al consultar categoría %d: %v", id, err)
		c.responderCategoria(http.StatusInternalServerError, "No se pudo consultar la categoría", nil)
	default:
		c.responderCategoria(http.StatusOK, "Categoría obtenida con éxito", &categoria)
	}
}

func (c *CategoriasController) Post() {
	var entrada struct {
		Nombre string `json:"nombre"`
	}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &entrada); err != nil {
		c.responderCategoria(http.StatusBadRequest, "El JSON recibido es inválido", nil)
		return
	}

	nombre := strings.TrimSpace(entrada.Nombre)
	if nombre == "" {
		c.responderCategoria(http.StatusBadRequest, "El nombre es obligatorio", nil)
		return
	}

	categoria, err := models.CrearCategoria(c.Ctx.Request.Context(), nombre)
	switch fallo, esErrorPG := err.(*pq.Error); {
	case err == nil:
		c.responderCategoria(http.StatusCreated, "Categoría creada con éxito", &categoria)
	case esErrorPG && fallo.Code == "23505":
		c.responderCategoria(http.StatusConflict, "Ya existe una categoría con ese nombre", nil)
	default:
		log.Printf("Error al crear categoría: %v", err)
		c.responderCategoria(http.StatusInternalServerError, "No se pudo guardar la categoría", nil)
	}
}

func (c *CategoriasController) responderCategoria(estado int, mensaje string, categoria *models.CategoriaPostgres) {
	respuesta := map[string]interface{}{
		"Success": estado >= 200 && estado < 300,
		"Status":  estado,
		"Message": mensaje,
	}
	if categoria != nil {
		respuesta["Data"] = categoria
	}
	c.Ctx.Output.SetStatus(estado)
	c.Data["json"] = respuesta
	_ = c.ServeJSON()
}

func (c *CategoriasController) Put() {
	id, err := strconv.ParseInt(c.Ctx.Input.Param(":id"), 10, 32)
	if err != nil || id <= 0 {
		c.responderCategoria(http.StatusBadRequest, "El ID debe ser un entero positivo", nil)
		return
	}

	var entrada struct {
		Nombre string `json:"nombre"`
	}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &entrada); err != nil {
		c.responderCategoria(http.StatusBadRequest, "El JSON recibido es inválido", nil)
		return
	}

	nombre := strings.TrimSpace(entrada.Nombre)
	if nombre == "" {
		c.responderCategoria(http.StatusBadRequest, "El nombre es obligatorio", nil)
		return
	}

	categoria, err := models.ActualizarCategoria(c.Ctx.Request.Context(), id, nombre)
	switch fallo, esErrorPG := err.(*pq.Error); {
	case err == nil:
		c.responderCategoria(http.StatusOK, "Categoría actualizada con éxito", &categoria)
	case err == sql.ErrNoRows:
		c.responderCategoria(http.StatusNotFound, "Categoría no encontrada", nil)
	case esErrorPG && fallo.Code == "23505":
		c.responderCategoria(http.StatusConflict, "Ya existe una categoría con ese nombre", nil)
	default:
		log.Printf("Error al actualizar categoría %d: %v", id, err)
		c.responderCategoria(http.StatusInternalServerError, "No se pudo actualizar la categoría", nil)
	}
}

func (c *CategoriasController) Delete() {
	id, err := strconv.ParseInt(c.Ctx.Input.Param(":id"), 10, 32)
	if err != nil || id <= 0 {
		c.responderCategoria(http.StatusBadRequest, "El ID debe ser un entero positivo", nil)
		return
	}

	err = models.DesactivarCategoria(c.Ctx.Request.Context(), id)
	switch {
	case err == sql.ErrNoRows:
		c.responderCategoria(http.StatusNotFound, "No existe una categoría activa con ese ID", nil)
	case err != nil:
		log.Printf("Error al desactivar categoría %d: %v", id, err)
		c.responderCategoria(http.StatusInternalServerError, "No se pudo desactivar la categoría", nil)
	default:
		c.responderCategoria(http.StatusOK, "Categoría desactivada con éxito", nil)
	}
}
