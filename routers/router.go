package routers

import (
	"GESTOCK_MID/controllers"

	web "github.com/beego/beego/v2/server/web"
)

func init() {
	web.Router("/incidencias", &controllers.IncidenciaController{})
	web.Router("/bodegas/:id", &controllers.BodegasController{}, "get:GetById")
	web.Router("/bodegas", &controllers.BodegasController{}, "post:Post")
}