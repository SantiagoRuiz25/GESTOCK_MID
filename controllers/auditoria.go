package controllers

import (
	"net/http"
	"strconv"

	"GESTOCK_MID/models"

	"github.com/gin-gonic/gin"
)

// Buscar todas las auditorías
func GetAuditorias(c *gin.Context) {
	auditorias, err := models.GetAllAuditorias()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, auditorias)
}

// Crear una auditoría
func CreateAuditoria(c *gin.Context) {
	var auditoria models.Auditoria
	if err := c.ShouldBindJSON(&auditoria); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := auditoria.Create(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, auditoria)
}

// Actualizar una auditoría
func UpdateAuditoria(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var auditoriaActualizada models.Auditoria
	if err := c.ShouldBindJSON(&auditoriaActualizada); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Llamamos al método Update pasando el ID de los parámetros de la ruta
	if err := auditoriaActualizada.Update(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Auditoría actualizada correctamente"})
}

// Eliminar una auditoría
func DeleteAuditoria(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	if err := models.DeleteAuditoria(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Auditoría eliminada correctamente"})
}