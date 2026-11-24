package routers

import (
	"GESTOCK_MID/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	// Usuarios y sesión
	beego.Router("/v1/usuarios", &controllers.UsuariosController{}, "get:GetAll;post:Post")
	beego.Router("/v1/usuarios/login", &controllers.UsuariosController{}, "post:Login")
	beego.Router("/v1/sesiones", &controllers.UsuariosController{}, "get:GetActiveSessions")
	beego.Router("/v1/logout/:id", &controllers.UsuariosController{}, "post:Logout")
	beego.Router("/v1/usuarios/:id", &controllers.UsuariosController{}, "put:Put;delete:Delete")

	// Bodegas
	beego.Router("/v1/bodega", &controllers.BodegasController{}, "get:GetAll;post:Post")
	beego.Router("/v1/bodega/:id", &controllers.BodegasController{}, "get:GetById;put:Put;delete:Delete")

	// Categorías
	beego.Router("/v1/categorias", &controllers.CategoriasController{}, "get:GetAll;post:Post")
	beego.Router("/v1/categorias/:id", &controllers.CategoriasController{}, "get:GetById;put:Put;delete:Delete")

	// Productos
	beego.Router("/v1/productos", &controllers.ProductosController{}, "get:GetAll;post:Post")
	beego.Router("/v1/productos/:id", &controllers.ProductosController{}, "get:GetById;put:Put;delete:Delete")

	// Operación
	beego.Router("/v1/operacion", &controllers.OperacionController{}, "get:GetAll;post:Post")
	beego.Router("/v1/operacion/:id", &controllers.OperacionController{}, "get:GetById;put:Put;delete:Delete")

	// Incidencias
	beego.Router("/v1/incidencia", &controllers.IncidenciaController{}, "get:GetAll;post:Post")
	beego.Router("/v1/incidencia/:id", &controllers.IncidenciaController{}, "get:GetById;put:Put;delete:Delete")

	// Mantenimiento
	beego.Router("/v1/mantenimiento", &controllers.MantenimientoController{}, "get:GetAll;post:Post")
	beego.Router("/v1/mantenimiento/:id", &controllers.MantenimientoController{}, "get:GetById;put:Put;delete:Delete")

	// Auditoría
	beego.Router("/v1/auditoria", &controllers.AuditoriaController{}, "get:GetAll;post:Post")
	beego.Router("/v1/auditoria/:id", &controllers.AuditoriaController{}, "get:GetById;put:Put;delete:Delete")
}