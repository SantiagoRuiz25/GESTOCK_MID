package routers

import (
	"GESTOCK_MID/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	// 1. Incidencias
	beego.Router("/incidencias", &controllers.IncidenciaController{}, "get:GetAll;post:Post")
	beego.Router("/incidencias/:id", &controllers.IncidenciaController{}, "get:GetById;delete:Delete")

	// 2. Mantenimientos
	beego.Router("/mantenimientos", &controllers.MantenimientoController{}, "get:GetAll;post:Post")
	beego.Router("/mantenimientos/:id", &controllers.MantenimientoController{}, "get:GetById")

	// 3. Bodegas
	beego.Router("/v1/bodegas", &controllers.BodegaController{}, "get:GetAll;post:Post")
	beego.Router("/v1/bodegas/:id", &controllers.BodegaController{}, "get:GetById;put:Put;delete:Delete")

	// 4. Operaciones
	beego.Router("/operaciones", &controllers.OperacionController{}, "get:GetAll;post:Post")
	beego.Router("/operaciones/:id", &controllers.OperacionController{}, "get:GetById;put:Put;delete:Delete")

	// 5. Auditorías
	beego.Router("/auditorias", &controllers.AuditoriaController{}, "get:GetAll;post:Post")
	beego.Router("/auditorias/:id", &controllers.AuditoriaController{}, "get:GetById")
}