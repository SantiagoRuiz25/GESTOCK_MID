package main

import (
	_ "GESTOCK_MID/routers"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	beego.Run()
}

