package routers

import (
	"GESTOCK_MID/controllers"

	web "github.com/beego/beego/v2/server/web"
)

func init() {
	// Rutas de bodegas e incidencias
	web.Router("/v1/bodegas/:id", &controllers.BodegasController{}, "get:GetById")
	web.Router("/v1/bodegas", &controllers.BodegasController{}, "post:Post")
}