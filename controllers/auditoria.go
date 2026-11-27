package controllers

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	beego "github.com/beego/beego/v2/server/web"
)

type AuditoriaController struct {
	beego.Controller
}

func (c *AuditoriaController) GetAll() {
	client := &http.Client{}
	req, errReq := http.NewRequest("GET", "http://localhost:8081/auditorias", nil)
	if errReq != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.Data["json"] = map[string]interface{}{"Success": false, "Message": "Error creando la petición al CRUD"}
		_ = c.ServeJSON()
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.Data["json"] = map[string]interface{}{"Success": false, "Message": "Error de conexión con el servicio CRUD"}
		_ = c.ServeJSON()
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var resultado map[string]interface{}
	_ = json.Unmarshal(body, &resultado)

	c.Ctx.Output.SetStatus(resp.StatusCode)
	c.Data["json"] = resultado
	_ = c.ServeJSON()
}

func (c *AuditoriaController) GetById() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		c.Ctx.Output.SetStatus(http.StatusBadRequest)
		c.Data["json"] = map[string]interface{}{"Success": false, "Message": "ID inválido"}
		_ = c.ServeJSON()
		return
	}

	client := &http.Client{}
	req, errReq := http.NewRequest("GET", "http://localhost:8081/auditorias/"+strconv.Itoa(id), nil)
	if errReq != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.Data["json"] = map[string]interface{}{"Success": false, "Message": "Error creando la petición"}
		_ = c.ServeJSON()
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.Data["json"] = map[string]interface{}{"Success": false, "Message": "Error de conexión con el CRUD"}
		_ = c.ServeJSON()
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var resultado map[string]interface{}
	_ = json.Unmarshal(body, &resultado)

	c.Ctx.Output.SetStatus(resp.StatusCode)
	c.Data["json"] = resultado
	_ = c.ServeJSON()
}

func (c *AuditoriaController) Post() {
	client := &http.Client{}
	req, errReq := http.NewRequest(
		"POST",
		"http://localhost:8081/auditorias",
		bytes.NewBuffer(c.Ctx.Input.RequestBody),
	)
	if errReq != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.Data["json"] = map[string]interface{}{"Success": false, "Message": "Error de conexión con el servicio CRUD"}
		_ = c.ServeJSON()
		return
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.Data["json"] = map[string]interface{}{"Success": false, "Message": "No se pudo comunicar con el servidor CRUD"}
		_ = c.ServeJSON()
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var resultado map[string]interface{}
	_ = json.Unmarshal(body, &resultado)

	c.Ctx.Output.SetStatus(resp.StatusCode)
	c.Data["json"] = resultado
	_ = c.ServeJSON()
}

func (c *AuditoriaController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		c.Ctx.Output.SetStatus(http.StatusBadRequest)
		c.Data["json"] = map[string]interface{}{"Success": false, "Message": "ID inválido"}
		_ = c.ServeJSON()
		return
	}

	client := &http.Client{}
	req, errReq := http.NewRequest(
		"PUT",
		"http://localhost:8081/auditorias/"+strconv.Itoa(id),
		bytes.NewBuffer(c.Ctx.Input.RequestBody),
	)
	if errReq != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.Data["json"] = map[string]interface{}{"Success": false, "Message": "Error creando la petición"}
		_ = c.ServeJSON()
		return
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.Data["json"] = map[string]interface{}{"Success": false, "Message": "Error de conexión con el CRUD"}
		_ = c.ServeJSON()
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var resultado map[string]interface{}
	_ = json.Unmarshal(body, &resultado)

	c.Ctx.Output.SetStatus(resp.StatusCode)
	c.Data["json"] = resultado
	_ = c.ServeJSON()
}

func (c *AuditoriaController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		c.Ctx.Output.SetStatus(http.StatusBadRequest)
		c.Data["json"] = map[string]interface{}{"Success": false, "Message": "ID inválido"}
		_ = c.ServeJSON()
		return
	}

	client := &http.Client{}
	req, errReq := http.NewRequest("DELETE", "http://localhost:8081/auditorias/"+strconv.Itoa(id), nil)
	if errReq != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.Data["json"] = map[string]interface{}{"Success": false, "Message": "Error creando la petición"}
		_ = c.ServeJSON()
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.Data["json"] = map[string]interface{}{"Success": false, "Message": "Error de conexión con el CRUD"}
		_ = c.ServeJSON()
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var resultado map[string]interface{}
	_ = json.Unmarshal(body, &resultado)

	c.Ctx.Output.SetStatus(resp.StatusCode)
	c.Data["json"] = resultado
	_ = c.ServeJSON()
}