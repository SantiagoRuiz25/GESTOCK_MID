package routers

import (
	"GESTOCK_MID/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	// 1. Usuarios, autenticación y sesiones
	beego.Router("/v1/usuarios", &controllers.UsuariosController{}, "get:GetAll;post:Post")
	beego.Router("/v1/usuarios/login", &controllers.UsuariosController{}, "post:Login")
	beego.Router("/v1/usuarios/:id", &controllers.UsuariosController{}, "put:Put;delete:Delete")
	beego.Router("/v1/sesiones", &controllers.UsuariosController{}, "get:GetActiveSessions")
	beego.Router("/v1/logout/:id", &controllers.UsuariosController{}, "post:Logout")

	// 2. Roles
	beego.Router("/v1/roles", &controllers.RolMidController{}, "get:GetAll;post:Post")
	beego.Router("/v1/roles/:id", &controllers.RolMidController{}, "get:GetOne;delete:Delete")

	// 3. Empresas
	beego.Router("/v1/empresa", &controllers.EmpresasMidController{}, "get:GetAll;post:Post")
	beego.Router("/v1/empresa/:id", &controllers.EmpresasMidController{}, "get:GetById;put:Put;delete:Delete")
	beego.Router("/v1/empresas", &controllers.EmpresasMidController{}, "get:GetAll;post:Post")
	beego.Router("/v1/empresas/:id", &controllers.EmpresasMidController{}, "get:GetById;put:Put;delete:Delete")

	// 4. Bodegas
	beego.Router("/v1/bodega", &controllers.BodegasController{}, "get:GetAll;post:Post")
	beego.Router("/v1/bodega/:id", &controllers.BodegasController{}, "get:GetById;put:Put;delete:Delete")
	beego.Router("/v1/bodegas", &controllers.BodegaController{}, "get:GetAll;post:Post")
	beego.Router("/v1/bodegas/:id", &controllers.BodegaController{}, "get:GetById;put:Put;delete:Delete")

	// 5. Categorías
	beego.Router("/v1/categorias", &controllers.CategoriasController{}, "get:GetAll;post:Post")
	beego.Router("/v1/categorias/:id", &controllers.CategoriasController{}, "get:GetById;put:Put;delete:Delete")

	// 6. Productos
	beego.Router("/v1/productos", &controllers.ProductosController{}, "get:GetAll;post:Post")
	beego.Router("/v1/productos/:id", &controllers.ProductosController{}, "get:GetById;put:Put;delete:Delete")

	// 7. Operaciones
	beego.Router("/v1/operacion", &controllers.OperacionController{}, "get:GetAll;post:Post")
	beego.Router("/v1/operacion/:id", &controllers.OperacionController{}, "get:GetById;put:Put;delete:Delete")
	beego.Router("/v1/operaciones", &controllers.OperacionController{}, "get:GetAll;post:Post")
	beego.Router("/v1/operaciones/:id", &controllers.OperacionController{}, "get:GetById;put:Put;delete:Delete")

	// 8. Incidencias
	beego.Router("/v1/incidencia", &controllers.IncidenciaController{}, "get:GetAll;post:Post")
	beego.Router("/v1/incidencia/:id", &controllers.IncidenciaController{}, "get:GetById;put:Put;delete:Delete")
	beego.Router("/v1/incidencias", &controllers.IncidenciaController{}, "get:GetAll;post:Post")
	beego.Router("/v1/incidencias/:id", &controllers.IncidenciaController{}, "get:GetById;delete:Delete")

	// 9. Mantenimientos
	beego.Router("/v1/mantenimiento", &controllers.MantenimientoController{}, "get:GetAll;post:Post")
	beego.Router("/v1/mantenimiento/:id", &controllers.MantenimientoController{}, "get:GetById;put:Put;delete:Delete")
	beego.Router("/v1/mantenimientos", &controllers.MantenimientoController{}, "get:GetAll;post:Post")
	beego.Router("/v1/mantenimientos/:id", &controllers.MantenimientoController{}, "get:GetById")

	// 10. Auditorías
	beego.Router("/v1/auditoria", &controllers.AuditoriaController{}, "get:GetAll;post:Post")
	beego.Router("/v1/auditoria/:id", &controllers.AuditoriaController{}, "get:GetById;put:Put;delete:Delete")
	beego.Router("/v1/auditorias", &controllers.AuditoriaController{}, "get:GetAll;post:Post")
	beego.Router("/v1/auditorias/:id", &controllers.AuditoriaController{}, "get:GetById")
}