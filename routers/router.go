package routers

import (
	"GESTOCK_MID/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	// 1. Usuarios y Sesiones
	beego.Router("/v1/usuarios", &controllers.UsuariosController{}, "get:GetAll;post:Post")
	beego.Router("/v1/usuarios/:id", &controllers.UsuariosController{}, "put:Put;delete:Delete")
	beego.Router("/v1/usuarios/login", &controllers.UsuariosController{}, "post:Login")
	beego.Router("/v1/sesiones", &controllers.UsuariosController{}, "get:GetActiveSessions")
	beego.Router("/v1/logout/:id", &controllers.UsuariosController{}, "post:Logout")

	// 2. Bodegas
	beego.Router("/v1/bodegas", &controllers.BodegasController{}, "get:GetAll;post:Post")
	beego.Router("/v1/bodegas/:id", &controllers.BodegasController{}, "get:GetById;put:Put;delete:Delete")

	// 3. Productos
	beego.Router("/v1/productos", &controllers.ProductosController{}, "get:GetAll;post:Post")
	beego.Router("/v1/productos/:id", &controllers.ProductosController{}, "get:GetById;put:Put;delete:Delete")

	// 4. Categorías
	beego.Router("/v1/categorias", &controllers.CategoriasController{}, "get:GetAll;post:Post")
	beego.Router("/v1/categorias/:id", &controllers.CategoriasController{}, "get:GetById;put:Put;delete:Delete")

	// 5. Auditorías
	beego.Router("/v1/auditorias", &controllers.AuditoriaController{}, "get:GetAll;post:Post")
	beego.Router("/v1/auditorias/:id", &controllers.AuditoriaController{}, "get:GetById;put:Put;delete:Delete")
}