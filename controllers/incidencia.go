package controllers

import(
	"GESTOCK_MID/models"
	"github.com/gin-gonic/gin"
	"GESTOCK_MID/db"
	
)

func GetIncidencias(c *gin.Context) {
	var incidencias []models.Incidencia
	err := db.GetAllIncidencias(&incidencias)
	if err != nil {
		c.JSON(500, gin.H{"error": "Error al obtener las incidencias"})
		return
	}
	c.JSON(200, incidencias)
}

func CreateIncidencia(c *gin.Context) {
	var incidencia models.Incidencia
	err := c.BindJSON(&incidencia)
	if err != nil {
		c.JSON(400, gin.H{"error": "Error al parsear la incidencia"})
		return
	}

	err = db.CreateIncidencia(&incidencia)
	if err != nil {
		c.JSON(500, gin.H{"error": "Error al crear la incidencia"})
		return
	}
	c.JSON(201, incidencia)
}