package controllers

import (
	"net/http"
	"GESTOCK_MID/models"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetOperaciones(c *gin.Context) {
	operaciones, err := models.GetAllOperaciones()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, operaciones)
}

func CreateOperacion(c *gin.Context) {
	var nuevaOperacion models.Operacion
	if err := c.ShouldBindJSON(&nuevaOperacion); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := models.CreateOperacion(&nuevaOperacion); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, nuevaOperacion)
}

func DeleteOperacion(c *gin.Context) {
	id := c.Param("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid operation ID"})
		return
	}
	if err := models.DeleteOperacion(intId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Operation deleted successfully"})
}