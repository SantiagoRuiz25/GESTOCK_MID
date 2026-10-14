package main

import (
<<<<<<< HEAD
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

=======
	_ "GESTOCK_MID/routers"

	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/lib/pq"
	_ "GESTOCK_MID/models" // Importamos los modelos para que ejecuten su propio init()
)

>>>>>>> 7c4a7ed89d2b84a7d26255830a656b0b2a4bbbc3
func main() {
	beego.Run()
}