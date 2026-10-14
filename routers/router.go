package routers

import (
	"GESTOCK_MID/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/v1/usuarios", &controllers.UsuariosController{}, "get:GetAll;post:Post")
	beego.Router("/v1/usuarios/:id", &controllers.UsuariosController{}, "put:Put;delete:Delete")
	beego.Router("/v1/usuarios/login", &controllers.UsuariosController{}, "post:Login")
	beego.Router("/v1/sesiones", &controllers.UsuariosController{}, "get:GetActiveSessions")
	beego.Router("/v1/logout/:id", &controllers.UsuariosController{}, "post:Logout")
	
	beego.Router("/v1/bodega", &controllers.BodegasController{}, "get:GetAll;post:Post")
	beego.Router("/v1/bodega/:id", &controllers.BodegasController{}, "get:GetById;put:Put;delete:Delete")
}