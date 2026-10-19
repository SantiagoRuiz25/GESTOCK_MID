package main

import (
	// Importa tus routers y modelos según la estructura que prefieras
	_ "GESTOCK_MID/routers" 
	
	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/client/orm"
	_ "github.com/lib/pq" // Driver de postgresql
)

func init() {
    // 1. Registra la conexión a PostgreSQL usando la configuración de tu carpeta conf/app.conf
    orm.RegisterDataBase("default", "postgres", "postgres://postgres:tu_contraseña@127.0.0.1:5432/gestock_db?sslmode=disable")
}

func main() {
	// 2. Si usas Beego nativo, las rutas se cargan automáticamente con el import de _ "GESTOCK_MID/routers".
	// Si prefieres inicializar un enrutador manual (como Gin o similar), lo declararías aquí:
	// r := routers.SetupRouter()
	// r.Run(":8080")

	// 3. Ejecuta el servidor web de Beego
	web.Run()
}