package controllers

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	beego "github.com/beego/beego/v2/server/web"
)

type IncidenciaController struct {
	beego.Controller
}

// 1. GET ALL: El MID pide la lista al CRUD (puerto 8081)
func (c *IncidenciaController) GetAll() {
	client := &http.Client{}
	req, errReq := http.NewRequest("GET", "http://localhost:8081/incidencias", nil)
	if errReq != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]interface{}{"Success": false, "Message": "Error creando la petición al CRUD"}
		c.ServeJSON()
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]interface{}{"Success": false, "Message": "Error de conexión con el servicio CRUD"}
		c.ServeJSON()
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var resultado map[string]interface{}
	json.Unmarshal(body, &resultado)

	c.Ctx.Output.SetStatus(resp.StatusCode)
	c.Data["json"] = resultado
	c.ServeJSON()
}

// 2. GET BY ID: Consulta una incidencia por ID en el CRUD (8081)
func (c *IncidenciaController) GetById() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)

	client := &http.Client{}
	req, errReq := http.NewRequest("GET", "http://localhost:8081/incidencias/"+strconv.Itoa(id), nil)
	if errReq != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]interface{}{"Success": false, "Message": "Error creando la petición"}
		c.ServeJSON()
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]interface{}{"Success": false, "Message": "Error de conexión con el CRUD"}
		c.ServeJSON()
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var resultado map[string]interface{}
	json.Unmarshal(body, &resultado)

	c.Ctx.Output.SetStatus(resp.StatusCode)
	c.Data["json"] = resultado
	c.ServeJSON()
}

// 3. POST: Recibe del usuario en el MID (8080) y lo reenvía al CRUD (8081) usando el buffer de tu imagen
func (c *IncidenciaController) Post() {
	client := &http.Client{}
	req, errReq := http.NewRequest(
		"POST",
		"http://localhost:8081/incidencias",
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

	resp, err := client.Do(req)
	if err != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  500,
			"Message": "No se pudo comunicar con el servidor CRUD",
		}
		c.ServeJSON()
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var resultado map[string]interface{}
	json.Unmarshal(body, &resultado)

	c.Ctx.Output.SetStatus(resp.StatusCode)
	c.Data["json"] = resultado
	c.ServeJSON()
}

// 4. DELETE: Elimina a través del CRUD (8081)
func (c *IncidenciaController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)

	client := &http.Client{}
	req, errReq := http.NewRequest("DELETE", "http://localhost:8081/incidencias/"+strconv.Itoa(id), nil)
	if errReq != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]interface{}{"Success": false, "Message": "Error en la petición"}
		c.ServeJSON()
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]interface{}{"Success": false, "Message": "Error de conexión con el CRUD"}
		c.ServeJSON()
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var resultado map[string]interface{}
	json.Unmarshal(body, &resultado)

	c.Ctx.Output.SetStatus(resp.StatusCode)
	c.Data["json"] = resultado
	c.ServeJSON()
}