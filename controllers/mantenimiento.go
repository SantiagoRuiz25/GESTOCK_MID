package controllers

import (
	"net/http"
	"GESTOCK_MID/models"
	"github.com/gin-gonic/gin"
	
)

// GET: Listar mantenimientos
func GetMantenimientos(c *gin.Context) {
	// Nota: Reemplaza db.DB por tu variable de conexión real si la tienes global
	mantenimientos, err := models.GetAllMantenimientos() 
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener los mantenimientos"})
		return
	}
	c.JSON(http.StatusOK, mantenimientos)
}

// POST: Crear mantenimiento
func CreateMantenimiento(c *gin.Context) {
	var nuevoMantenimiento models.Mantenimiento

	if err := c.ShouldBindJSON(&nuevoMantenimiento); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	if err := nuevoMantenimiento.Create(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al guardar el mantenimiento"})
		return
	}

	c.JSON(http.StatusCreated, nuevoMantenimiento)
}