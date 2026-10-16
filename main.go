package main

import (
	"GESTOCK_MID/db"
	_ "GESTOCK_MID/models" // Para que ejecuten su init()
	"GESTOCK_MID/routers"

	_ "github.com/lib/pq"
)

func main() {
	// 1. Inicializa la conexión a la base de datos usando tu paquete local 'db'
	db.ConnectDB()

	// 2. Configura e inicia las rutas de tu API (usando tu enrutador)
	r := routers.SetupRouter()
	
	// 3. Arranca el servidor local en el puerto 8080
	r.Run(":8080")
}