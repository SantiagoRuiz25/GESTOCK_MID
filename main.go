package main

import (
	"GESTOCK_MID/models"
	_ "GESTOCK_MID/routers"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
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