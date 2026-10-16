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

	// Rutas para mantenimientos
	r.GET("/mantenimientos", controllers.GetMantenimientos)
	r.POST("/mantenimientos", controllers.CreateMantenimiento)
	r.PUT("/mantenimientos/:id", controllers.UpdateMantenimiento)

	// Rutas para operaciones
	r.GET("/operaciones", controllers.GetOperaciones)
	r.POST("/operaciones", controllers.CreateOperacion)
	r.DELETE("/operaciones/:id", controllers.DeleteOperacion)

	return r
}