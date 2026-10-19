package controllers

import (
	"net/http"
	"strconv"

	"GESTOCK_MID/models"

	"github.com/gin-gonic/gin"
)

// GET: Listar mantenimientos
func GetMantenimientos(c *gin.Context) {
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

// PUT: Actualizar mantenimiento
func UpdateMantenimiento(c *gin.Context) {
	// 1. Obtener y validar el ID de la URL
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de mantenimiento inválido"})
		return
	}

	// 2. Leer el cuerpo de la petición (JSON)
	var mantenimiento models.Mantenimiento
	if err := c.ShouldBindJSON(&mantenimiento); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	// 3. Asignar el ID obtenido de la URL al modelo
	mantenimiento.ID = id

	// 4. Ejecutar la actualización en la base de datos
	if err := models.UpdateMantenimiento(&mantenimiento); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar el mantenimiento"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Mantenimiento actualizado exitosamente",
		"data":    mantenimiento,
	})
}