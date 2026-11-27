package main

import "GESTOCK_MID/models"

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
	"github.com/beego/beego/v2/server/web/filter/cors"
	_ "github.com/lib/pq"
	_ "GESTOCK_MID/routers"
	_ "GESTOCK_MID/models"
)



func main() {
	if err := models.ConectarPostgres(); err != nil {
		panic(err)
	}
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins: []string{"http://localhost:4200"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders: []string{
			"Origin",
			"Content-Type",
			"Authorization",
			"Accept",
			"X-Requested-With",
		},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Responder rápido a peticiones preflight OPTIONS
	beego.InsertFilter("*", beego.BeforeRouter, func(ctx *context.Context) {
		if ctx.Input.Method() == "OPTIONS" {
			ctx.Output.SetStatus(200)
			_ = ctx.Output.Body([]byte("ok"))
		}
	})

	beego.Run()
}
