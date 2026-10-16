package controllers

import (
	"net/http"
	"strconv"

	"GESTOCK_MID/models"

	"github.com/gin-gonic/gin"
)

// GET: Listar todos los mantenimientos
func GetMantenimientos(c *gin.Context) {
	mantenimientos, err := models.GetAllMantenimientos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener los mantenimientos"})
		return
	}
	c.JSON(http.StatusOK, mantenimientos)
}

// POST: Crear un nuevo mantenimiento
func CreateMantenimiento(c *gin.Context) {
	var nuevoMantenimiento models.Mantenimiento

	if err := c.ShouldBindJSON(&nuevoMantenimiento); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos o formato JSON incorrecto"})
		return
	}

	if err := nuevoMantenimiento.Create(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al guardar el mantenimiento"})
		return
	}

	c.JSON(http.StatusCreated, nuevoMantenimiento)
}

// PUT: Actualizar un mantenimiento existente
func UpdateMantenimiento(c *gin.Context) {
	// Opcional: si manejas el ID por la URL (ej: /mantenimientos/:id)
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de mantenimiento inválido"})
		return
	}

	var mantenimiento models.Mantenimiento
	if err := c.ShouldBindJSON(&mantenimiento); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}
	
	// Aseguramos que el ID de la estructura coincida con el de la URL
	mantenimiento.ID = id

	// Nota: Asegúrate de tener implementada la función UpdateMantenimiento en tu modelo si vas a usar esta ruta PUT
	if err := models.UpdateMantenimiento(&mantenimiento); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar el mantenimiento"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Mantenimiento actualizado exitosamente", "data": mantenimiento})
}