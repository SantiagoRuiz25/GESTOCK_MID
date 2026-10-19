package main

import (
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/lib/pq"
	_ "GESTOCK_MID/routers"
	_ "GESTOCK_MID/models"
)



func main() {
	beego.Run()
}