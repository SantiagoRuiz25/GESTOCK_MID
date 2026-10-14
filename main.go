package main

import (
	_ "GESTOCK_MID/routers"

	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/lib/pq"
	_ "GESTOCK_MID/models" // Importamos los modelos para que ejecuten su propio init()
)

func main() {
	beego.Run()
}