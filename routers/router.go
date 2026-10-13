package routers

import (
	"GESTOCK_MID/controllers"

	web "github.com/github.com/beego/beego/v2/server/web"
)

func init() {
	ns := web.NewNamespace("/v1",
		// Rutas para bodegas
		web.NSNamespace("/bodegas",
			web.NSRouter("/", &controllers.BodegasController{}, "get:GetAll;post:Post"),
			web.NSRouter("/:id", &controllers.BodegasController{}, "get:GetOne;put:Put;delete:Delete"),
		),

		// Rutas para incidencias
		web.NSNamespace("/incidencias",
			web.NSRouter("/", &controllers.IncidenciasController{}, "get:GetAll;post:Post"),
		),

		// Rutas para auditorias
		web.NSNamespace("/auditorias",
			web.NSRouter("/", &controllers.AuditoriasController{}, "get:GetAll;post:Post"),
			web.NSRouter("/:id", &controllers.AuditoriasController{}, "put:Put;delete:Delete"),
		),

		// Rutas para mantenimientos
		web.NSNamespace("/mantenimientos",
			web.NSRouter("/", &controllers.MantenimientosController{}, "get:GetAll;post:Post"),
			web.NSRouter("/:id", &controllers.MantenimientosController{}, "put:Put"),
		),

		// Rutas para operaciones
		web.NSNamespace("/operaciones",
			web.NSRouter("/", &controllers.OperacionesController{}, "get:GetAll;post:Post"),
			web.NSRouter("/:id", &controllers.OperacionesController{}, "delete:Delete"),
		),
	)

	web.AddNamespace(ns)
}