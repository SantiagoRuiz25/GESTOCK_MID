package controllers

import (
<<<<<<< HEAD
=======
	"bytes"
>>>>>>> cc7130cc06b3decdb797421a33962ecf6505d501
	"encoding/json"
	"net/http"
	"strconv"

	"GESTOCK_MID/models"

<<<<<<< HEAD
	beego "github.com/beego/beego/v2/server/web"
=======
	"github.com/beego/beego/v2/server/web"
>>>>>>> cc7130cc06b3decdb797421a33962ecf6505d501
)

type BodegasController struct {
	beego.Controller
}

// 1. GET ALL: Consultar todas las bodegas
func (c *BodegasController) GetAll() {
	respuesta, errCrud := http.Get("http://localhost:8080/v1/bodega")
	if errCrud != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  500,
			"Message": "Error de conexión con el servicio CRUD",
		}
		c.ServeJSON()
		return
	}
	defer respuesta.Body.Close()

	if respuesta.StatusCode == 404 {
		c.Ctx.Output.SetStatus(404)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  404,
			"Message": "No se encontraron bodegas registradas",
		}
		c.ServeJSON()
		return
	}

	var listaBodegas []models.Bodegas
	json.NewDecoder(respuesta.Body).Decode(&listaBodegas)

	c.Ctx.Output.SetStatus(200)
	c.Data["json"] = map[string]interface{}{
		"Success": true,
		"Status":  200,
		"Message": "Lista de bodegas obtenida con éxito",
		"Data":    listaBodegas,
	}
	c.ServeJSON()
}

// 2. GET BY ID: Consultar una bodega por ID
func (c *BodegasController) GetById() {
	idBodega := c.Ctx.Input.Param(":id")

	if idBodega == "" {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  400,
			"Message": "El ID es requerido",
		}
		c.ServeJSON()
		return
	}

	id, err := strconv.Atoi(idBodega)
	if err != nil || id <= 0 {
<<<<<<< HEAD
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  404,
			"Mensaje": "Id ingresado es invalido",
=======
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  400,
			"Message": "Id ingresado es invalido",
>>>>>>> cc7130cc06b3decdb797421a33962ecf6505d501
		}
		c.ServeJSON()
		return
	}

	respuesta, errCrud := http.Get("http://localhost:8080/v1/bodega/" + strconv.Itoa(id))
	if errCrud != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  500,
			"Message": "Error de conexión con el servicio CRUD",
		}
		c.ServeJSON()
		return
	}
	defer respuesta.Body.Close()

	if respuesta.StatusCode == 404 {
		c.Ctx.Output.SetStatus(404)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  404,
			"Message": "La bodega solicitada no existe",
		}
		c.ServeJSON()
		return
	}

	var bodega models.Bodegas
	json.NewDecoder(respuesta.Body).Decode(&bodega)

	c.Ctx.Output.SetStatus(200)
	c.Data["json"] = map[string]interface{}{
		"Success": true,
		"Status":  200,
		"Message": "Bodega obtenida con éxito",
		"Data":    bodega,
	}
	c.ServeJSON()
}

<<<<<<< HEAD
=======
// 3. POST: Crear nueva bodega
>>>>>>> cc7130cc06b3decdb797421a33962ecf6505d501
func (c *BodegasController) Post() {
	var nuevaBodega models.Bodegas

	err := json.Unmarshal(c.Ctx.Input.RequestBody, &nuevaBodega)
	if err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  400,
			"Message": "El formato del JSON recibido es invalido",
		}
		c.ServeJSON()
		return
	}

	if nuevaBodega.Nombre == "" || nuevaBodega.Codigo <= 0 {
<<<<<<< HEAD
=======
		c.Ctx.Output.SetStatus(400)
>>>>>>> cc7130cc06b3decdb797421a33962ecf6505d501
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  400,
			"Message": "Campos obligatorios incompletos",
		}
		c.ServeJSON()
		return
	}

	respuesta, errCrud := http.Post(
		"http://localhost:8081/v1/bodega",
		"application/json",
		bytes.NewBuffer(c.Ctx.Input.RequestBody),
	)
	if errCrud != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  500,
			"Message": "Error de conexión con el servicio CRUD",
		}
		c.ServeJSON()
		return
	}
	defer respuesta.Body.Close()

	var bodegaGuardada models.Bodegas
	json.NewDecoder(respuesta.Body).Decode(&bodegaGuardada)

	c.Ctx.Output.SetStatus(201)
	c.Data["json"] = map[string]interface{}{
		"Success": true,
		"Status":  201,
		"Message": "Bodega creada exitosamente",
		"Data":    bodegaGuardada,
	}
	c.ServeJSON()
}
<<<<<<< HEAD
=======

// 4. PUT: Actualizar bodega por ID
func (c *BodegasController) Put() {
	idBodega := c.Ctx.Input.Param(":id")

	if idBodega == "" {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  400,
			"Message": "El ID es requerido",
		}
		c.ServeJSON()
		return
	}

	id, err := strconv.Atoi(idBodega)
	if err != nil || id <= 0 {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  400,
			"Message": "Id ingresado es invalido",
		}
		c.ServeJSON()
		return
	}

	var bodegaActualizada models.Bodegas
	errParse := json.Unmarshal(c.Ctx.Input.RequestBody, &bodegaActualizada)
	if errParse != nil {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  400,
			"Message": "El formato del JSON recibido es invalido",
		}
		c.ServeJSON()
		return
	}

	client := &http.Client{}
	req, errReq := http.NewRequest(
		"PUT",
		"http://localhost:8080/v1/bodega/"+strconv.Itoa(id),
		bytes.NewBuffer(c.Ctx.Input.RequestBody),
	)
	if errReq != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  500,
			"Message": "Error de conexión con el servicio CRUD",
		}
		c.ServeJSON()
		return
	}
	req.Header.Set("Content-Type", "application/json")

	respuesta, errCrud := client.Do(req)
	if errCrud != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  500,
			"Message": "Error de conexión con el servicio CRUD",
		}
		c.ServeJSON()
		return
	}
	defer respuesta.Body.Close()

	var bodegaRespuesta models.Bodegas
	json.NewDecoder(respuesta.Body).Decode(&bodegaRespuesta)

	c.Ctx.Output.SetStatus(200)
	c.Data["json"] = map[string]interface{}{
		"Success": true,
		"Status":  200,
		"Message": "Bodega actualizada exitosamente",
		"Data":    bodegaRespuesta,
	}
	c.ServeJSON()
}

// 5. DELETE: Eliminar bodega por ID
func (c *BodegasController) Delete() {
	idBodega := c.Ctx.Input.Param(":id")

	if idBodega == "" {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  400,
			"Message": "El ID es requerido",
		}
		c.ServeJSON()
		return
	}

	id, err := strconv.Atoi(idBodega)
	if err != nil || id <= 0 {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  400,
			"Message": "Id ingresado es invalido",
		}
		c.ServeJSON()
		return
	}

	client := &http.Client{}
	req, errReq := http.NewRequest("DELETE", "http://localhost:8080/v1/bodega/"+strconv.Itoa(id), nil)
	if errReq != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  500,
			"Message": "Error de conexión con el servicio CRUD",
		}
		c.ServeJSON()
		return
	}

	respuesta, errCrud := client.Do(req)
	if errCrud != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  500,
			"Message": "Error de conexión con el servicio CRUD",
		}
		c.ServeJSON()
		return
	}
	defer respuesta.Body.Close()

	c.Ctx.Output.SetStatus(200)
	c.Data["json"] = map[string]interface{}{
		"Success": true,
		"Status":  200,
		"Message": "Bodega eliminada correctamente",
	}
	c.ServeJSON()
}
>>>>>>> cc7130cc06b3decdb797421a33962ecf6505d501
