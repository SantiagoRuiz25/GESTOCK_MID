package controllers

import (
	"GESTOCK_MID/models"
	"net/http"
	"github.com/gin-gonic/gin"
	"strconv"
)

// buscar

func GetAuditorias(c *gin.Context) {
	auditorias, err := models.GetAllAuditorias()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, auditorias)
}


// crear

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

// actualizar

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

	if err := auditoriaActualizada.Update(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Auditoría actualizada correctamente"})
}

// eliminar

func DeleteAuditoria(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var auditoria models.Auditoria
	if err := auditoria.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Auditoría eliminada correctamente"})
}