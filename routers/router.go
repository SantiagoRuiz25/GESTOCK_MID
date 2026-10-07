package routers

import (
	"GESTOCK_MID/controllers"
	beego "github.com/beego/beego/v2/server/web"
	
)

func init() {
    web.Router("/incidencias", &controllers.IncidenciaController{})
	
}
