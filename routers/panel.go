package routers

import (
	"GESTOCK_MID/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/v1/panel/resumen", &controllers.PanelController{}, "get:GetResumen")
}
