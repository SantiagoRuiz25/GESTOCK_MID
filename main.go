package main

import (
	_ "GESTOCK_MID/routers"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/client/orm"
	"GESTOCK_MID/models"
	
	_ "github.com/lib/pq"
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
	beego.Run()
}