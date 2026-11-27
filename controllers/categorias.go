package controllers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
	"strings"

	"GESTOCK_MID/models"

	"github.com/beego/beego/v2/server/web"
	"github.com/lib/pq"
)

type CategoriasController struct {
	web.Controller
}

func (c *CategoriasController) GetAll() {
	categorias, err := models.ListarCategorias(c.Ctx.Request.Context())
	if err != nil {
		c.errorCategoria(err)
		return
	}
	c.responderCategoria(http.StatusOK, "Categorías obtenidas con éxito", categorias)
}

func (c *CategoriasController) GetById() {
	id, ok := c.leerIDCategoria()
	if !ok {
		return
	}

	categoria, err := models.BuscarCategoriaPorID(c.Ctx.Request.Context(), id)
	if err != nil {
		c.errorCategoria(err)
		return
	}
	c.responderCategoria(http.StatusOK, "Categoría obtenida con éxito", categoria)
}

func (c *CategoriasController) Post() {
	nombre, ok := c.leerNombreCategoria()
	if !ok {
		return
	}

	categoria, err := models.CrearCategoria(c.Ctx.Request.Context(), nombre)
	if err != nil {
		c.errorCategoria(err)
		return
	}
	c.responderCategoria(http.StatusCreated, "Categoría creada con éxito", categoria)
}

func (c *CategoriasController) Put() {
	id, ok := c.leerIDCategoria()
	if !ok {
		return
	}
	nombre, ok := c.leerNombreCategoria()
	if !ok {
		return
	}

	categoria, err := models.ActualizarCategoria(c.Ctx.Request.Context(), id, nombre)
	if err != nil {
		c.errorCategoria(err)
		return
	}
	c.responderCategoria(http.StatusOK, "Categoría actualizada con éxito", categoria)
}

func (c *CategoriasController) Delete() {
	id, ok := c.leerIDCategoria()
	if !ok {
		return
	}

	err := models.DesactivarCategoria(c.Ctx.Request.Context(), id)
	if errors.Is(err, sql.ErrNoRows) {
		c.responderCategoria(http.StatusNotFound, "No existe una categoría activa con ese ID", nil)
		return
	}
	if err != nil {
		c.errorCategoria(err)
		return
	}
	c.responderCategoria(http.StatusOK, "Categoría desactivada con éxito", nil)
}

func (c *CategoriasController) leerIDCategoria() (int64, bool) {
	id, err := strconv.ParseInt(c.Ctx.Input.Param(":id"), 10, 32)
	if err != nil || id <= 0 {
		c.responderCategoria(http.StatusBadRequest, "El ID debe ser un entero positivo", nil)
		return 0, false
	}
	return id, true
}

func (c *CategoriasController) leerNombreCategoria() (string, bool) {
	var entrada struct {
		Nombre string `json:"nombre"`
	}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &entrada); err != nil {
		c.responderCategoria(http.StatusBadRequest, "El JSON recibido es inválido", nil)
		return "", false
	}

	nombre := strings.TrimSpace(entrada.Nombre)
	if nombre == "" {
		c.responderCategoria(http.StatusBadRequest, "El nombre es obligatorio", nil)
		return "", false
	}
	return nombre, true
}

func (c *CategoriasController) errorCategoria(err error) {
	var fallo *pq.Error
	switch {
	case errors.Is(err, sql.ErrNoRows):
		c.responderCategoria(http.StatusNotFound, "Categoría no encontrada", nil)
	case errors.As(err, &fallo) && fallo.Code == "23505":
		c.responderCategoria(http.StatusConflict, "Ya existe una categoría con ese nombre", nil)
	default:
		log.Printf("Error en Categorías: %v", err)
		c.responderCategoria(http.StatusInternalServerError, "No se pudo completar la operación de categoría", nil)
	}
}

func (c *CategoriasController) responderCategoria(estado int, mensaje string, datos interface{}) {
	respuesta := map[string]interface{}{
		"Success": estado >= 200 && estado < 300,
		"Status":  estado,
		"Message": mensaje,
	}
	if datos != nil {
		respuesta["Data"] = datos
	}
	c.Ctx.Output.SetStatus(estado)
	c.Data["json"] = respuesta
	_ = c.ServeJSON()
}
