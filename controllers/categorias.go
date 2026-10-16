package controllers
import (
	"encoding/json"
	"net/http"
	"strconv"

	"GESTOCK_MID/models"
	"github.com/beego/beego/v2/server/web"
)

type CategoriasController struct {
	web.Controller
}

func (c *CategoriasController) GetById() {
	idCategoria := c.Ctx.Input.Param(":id")

	if idCategoria == "" {
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  400,
			"Message": "El ID es requerido",
		}
		c.ServeJSON()
		return
	}

	id, err := strconv.Atoi(idCategoria)
	if err != nil || id <= 0 {
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  400,
			"Message": "Id ingresado es invalido",
		}
		c.ServeJSON()
		return
	}

	respuesta, errCrud := http.Get("http://localhost:8081/v1/categoria/" + strconv.Itoa(id))
	if errCrud != nil {
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  500,
			"Message": "Error de conexión con la base de datos o servicio CRUD",
		}
		c.ServeJSON()
		return
	}
	defer respuesta.Body.Close()

	if respuesta.StatusCode == 404 {
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  404,
			"Message": "No se encontró la categoría solicitada en el sistema",
		}
		c.ServeJSON()
		return
	}

	var categoria models.Categoria
	json.NewDecoder(respuesta.Body).Decode(&categoria)

	c.Data["json"] = map[string]interface{}{
		"Success": true,
		"Status":  200,
		"Message": "Categoría obtenida con éxito",
		"Data":    categoria,
	}
	c.ServeJSON()
}

func (c *CategoriasController) Post() {
	var nuevaCategoria models.Categoria

	err := json.Unmarshal(c.Ctx.Input.RequestBody, &nuevaCategoria)
	if err != nil {
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  400,
			"Message": "El formato del JSON recibido es invalido",
		}
		c.ServeJSON()
		return
	}

	if nuevaCategoria.Nombre == "" {
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  400,
			"Message": "Campos obligatorios incompletos",
		}
		c.ServeJSON()
		return
	}

	c.Data["json"] = map[string]interface{}{
		"Success": true,
		"Status":  201,
		"Message": "Categoría creada exitosamente",
		"Data":    nuevaCategoria,
	}
	c.ServeJSON()
}