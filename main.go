package main

import (
<<<<<<< HEAD
	_ "GESTOCK_MID/routers"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/client/orm"
	"GESTOCK_MID/models"
	
	_ "github.com/lib/pq"
=======
	"GESTOCK_MID/routers"
>>>>>>> cc7130cc06b3decdb797421a33962ecf6505d501
)

func init() {

	orm.RegisterModel(
		new(models.Usuario),
		new(models.SesionActiva),
		new(models.Bodegas),
		new(models.Auditoria),
		new(models.Incidencia),
		new(models.Mantenimiento),
	)
}

func main() {
<<<<<<< HEAD
	beego.Run()
=======
	// Inicializamos el enrutador de Gin
	r := routers.SetupRouter()

	// Arrancamos el servidor en el puerto 8080
	r.Run(":8080")
>>>>>>> cc7130cc06b3decdb797421a33962ecf6505d501
}