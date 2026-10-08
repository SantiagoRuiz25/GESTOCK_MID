package controllers

import (
	"net/http"
	"GESTOCK_MID/models"
	"github.com/gin-gonic/gin"
)

// Obtener todas las incidencias
func GetIncidencias(c *gin.Context) {
	incidencias, err := models.GetAllIncidencias()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener las incidencias"})
		return
	}
	c.JSON(http.StatusOK, incidencias)
}

// Crear una incidencia
func CreateIncidencia(c *gin.Context) {
	var incidencia models.Incidencia
	
	if err := c.ShouldBindJSON(&incidencia); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error al parsear la incidencia"})
		return
	}

	if err := incidencia.Create(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear la incidencia"})
		return
	}
	
	c.JSON(http.StatusCreated, incidencia)
}