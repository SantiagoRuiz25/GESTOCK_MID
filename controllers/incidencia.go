package controllers

import (
	"net/http"

	"GESTOCK_MID/models"

	"github.com/gin-gonic/gin"
)

// GET: Buscar / Listar incidencias
func GetIncidencias(c *gin.Context) {
	incidencias, err := models.GetAllIncidencias()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener las incidencias: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, incidencias)
}

// POST: Crear incidencia
func CreateIncidencia(c *gin.Context) {
	var incidencia models.Incidencia

	if err := c.ShouldBindJSON(&incidencia); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos o formato JSON incorrecto"})
		return
	}

	if err := incidencia.Create(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear la incidencia: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, incidencia)
}