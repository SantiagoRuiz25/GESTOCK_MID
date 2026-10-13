package main

import (
	"GESTOCK_MID/routers"
)

func main() {
	// Inicializamos el enrutador de Gin
	r := routers.SetupRouter()

	// Arrancamos el servidor en el puerto 8080
	r.Run(":8080")
}