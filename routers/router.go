package routers

import (
	"GESTOCK_MID/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	// 1. Roles (Ya tiene todos sus métodos implementados en su controlador y servicio)
	beego.Router("/v1/rol", &controllers.RolMidController{}, "get:GetAll;post:Post")
	beego.Router("/v1/rol/:id", &controllers.RolMidController{}, "get:GetOne;delete:Delete")
	beego.Router("/v1/roles", &controllers.RolMidController{}, "get:GetAll;post:Post")
	beego.Router("/v1/roles/:id", &controllers.RolMidController{}, "get:GetOne;delete:Delete")

	// 2. Sesiones / Sesiones Activas
	beego.Router("/v1/sesion", &controllers.SesionController{}, "get:GetAll")
	beego.Router("/v1/sesion/:id", &controllers.SesionController{}, "delete:Delete")
	beego.Router("/v1/sesiones", &controllers.SesionController{}, "get:GetAll")
	beego.Router("/v1/sesiones/:id", &controllers.SesionController{}, "delete:Delete")

	// 3. Usuarios (Mapeado estrictamente a los métodos que tu controlador de usuarios soporta por ahora)
	beego.Router("/v1/usuario", &controllers.UsuariosController{}, "get:GetAll;post:Post")
	beego.Router("/v1/usuarios", &controllers.UsuariosController{}, "get:GetAll;post:Post")

	// 4. Bodegas (Ajustado a los métodos disponibles)
	beego.Router("/v1/bodega", &controllers.BodegaController{}, "get:GetAll;post:Post")
	beego.Router("/v1/bodegas", &controllers.BodegaController{}, "get:GetAll;post:Post")

	// 5. Productos (Ajustado a los métodos disponibles)
	beego.Router("/v1/producto", &controllers.ProductosController{}, "get:GetAll;post:Post")
	beego.Router("/v1/productos", &controllers.ProductosController{}, "get:GetAll;post:Post")

	// 6. Categorías (Ajustado a los métodos disponibles)
	beego.Router("/v1/categoria", &controllers.CategoriasController{}, "get:GetAll;post:Post")
	beego.Router("/v1/categorias", &controllers.CategoriasController{}, "get:GetAll;post:Post")
}