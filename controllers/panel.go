package controllers

import (
	"context"
	"errors"
	"log"
	"net/http"
	"time"

	"GESTOCK_MID/services"

	beego "github.com/beego/beego/v2/server/web"
)

type PanelController struct {
	beego.Controller
}

func (c *PanelController) GetResumen() {
	ctx, cancelar := context.WithTimeout(c.Ctx.Request.Context(), 10*time.Second)
	defer cancelar()

	resumen, err := services.ObtenerResumenPanel(ctx)
	if err != nil {
		log.Printf("Error al consultar el Panel: %v", err)
		estado := http.StatusBadGateway
		mensaje := "No se pudo obtener el resumen del Panel"
		if errors.Is(err, context.DeadlineExceeded) {
			estado = http.StatusGatewayTimeout
			mensaje = "El servicio CRUD tardó demasiado en responder"
		}
		c.Ctx.Output.SetStatus(estado)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  estado,
			"Message": mensaje,
		}
		_ = c.ServeJSON()
		return
	}

	c.Data["json"] = map[string]interface{}{
		"Success": true,
		"Status":  http.StatusOK,
		"Message": "Resumen del Panel obtenido con éxito",
		"Data":    resumen,
	}
	_ = c.ServeJSON()
}
