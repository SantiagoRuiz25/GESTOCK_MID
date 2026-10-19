package controllers

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	beego "github.com/beego/beego/v2/server/web"
)

type OperacionController struct {
	beego.Controller
}

func (c *OperacionController) GetAll() {
	client := &http.Client{}
	req, errReq := http.NewRequest("GET", "http://localhost:8081/operaciones", nil)
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

func (c *OperacionController) GetById() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)

	client := &http.Client{}
	req, errReq := http.NewRequest("GET", "http://localhost:8081/operaciones/"+strconv.Itoa(id), nil)
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

func (c *OperacionController) Post() {
	client := &http.Client{}
	req, errReq := http.NewRequest(
		"POST",
		"http://localhost:8081/operaciones",
		bytes.NewBuffer(c.Ctx.Input.RequestBody),
	)

	if errReq != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]interface{}{"Success": false, "Message": "Error de conexión con el servicio CRUD"}
		c.ServeJSON()
		return
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]interface{}{"Success": false, "Message": "No se pudo comunicar con el servidor CRUD"}
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

func (c *OperacionController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)

	client := &http.Client{}
	req, errReq := http.NewRequest(
		"PUT",
		"http://localhost:8081/operaciones/"+strconv.Itoa(id),
		bytes.NewBuffer(c.Ctx.Input.RequestBody),
	)

	if errReq != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]interface{}{"Success": false, "Message": "Error de conexión con el servicio CRUD"}
		c.ServeJSON()
		return
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]interface{}{"Success": false, "Message": "No se pudo comunicar con el servidor CRUD"}
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

func (c *OperacionController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)

	client := &http.Client{}
	req, errReq := http.NewRequest("DELETE", "http://localhost:8081/operaciones/"+strconv.Itoa(id), nil)
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