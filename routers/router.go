package routers

import (
	"GESTOCK_MID/controllers"
	"github.com/gin-gonic/gin"
	
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	// Rutas para incidencias
	r.GET("/incidencias", controllers.GetIncidencias)
	r.POST("/incidencias", controllers.CreateIncidencia)
	// Rutas para auditorías
	r.GET("/auditorias", controllers.GetAuditorias)
	r.POST("/auditorias", controllers.CreateAuditoria)
	r.PUT("/auditorias/:id", controllers.UpdateAuditoria)
	r.DELETE("/auditorias/:id", controllers.DeleteAuditoria)
	return r
}