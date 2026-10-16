package main

import (
	_ "GESTOCK_MID/routers"

	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/lib/pq"
	_ "GESTOCK_MID/models" 
	"GESTOCK_MID/db"// Importamos los modelos para que ejecuten su propio init()
	"GESTOCK_MID/routers"
)

func main() {
	beego.Run()
	db.ConnectDB()

	// 2. Configura las rutas de tu API
	r := routers.SetupRouter()

	// 3. Arranca el servidor local
	r.Run(":8080")
}




