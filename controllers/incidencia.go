package controllers

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	beego "github.com/beego/beego/v2/server/web"
)

const crudIncidenciaURL = "http://localhost:8081/incidencias"

type IncidenciaController struct {
	beego.Controller
}

// GetAll consulta todas las incidencias desde el servicio CRUD (puerto 8081).
func (c *IncidenciaController) GetAll() {
	client := &http.Client{}
	req, errReq := http.NewRequest(http.MethodGet, crudIncidenciaURL, nil)
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

// GetById consulta una incidencia por ID en el servicio CRUD (puerto 8081).
func (c *IncidenciaController) GetById() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		c.Ctx.Output.SetStatus(http.StatusBadRequest)
		c.Data["json"] = map[string]interface{}{"Success": false, "Message": "ID de incidencia inválido"}
		_ = c.ServeJSON()
		return
	}

	client := &http.Client{}
	req, errReq := http.NewRequest(http.MethodGet, crudIncidenciaURL+"/"+strconv.Itoa(id), nil)
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

// Post reenvía la creación de una incidencia al servicio CRUD (puerto 8081).
func (c *IncidenciaController) Post() {
	client := &http.Client{}
	req, errReq := http.NewRequest(
		http.MethodPost,
		crudIncidenciaURL,
		bytes.NewBuffer(c.Ctx.Input.RequestBody),
	)
	if errReq != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  http.StatusInternalServerError,
			"Message": "Error de conexión con el servicio CRUD",
		}
		_ = c.ServeJSON()
		return
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  http.StatusInternalServerError,
			"Message": "No se pudo comunicar con el servidor CRUD",
		}
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

// Put actualiza una incidencia por ID en el servicio CRUD (puerto 8081).
func (c *IncidenciaController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		c.Ctx.Output.SetStatus(http.StatusBadRequest)
		c.Data["json"] = map[string]interface{}{"Success": false, "Message": "ID de incidencia inválido"}
		_ = c.ServeJSON()
		return
	}

	client := &http.Client{}
	req, errReq := http.NewRequest(
		http.MethodPut,
		crudIncidenciaURL+"/"+strconv.Itoa(id),
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

// Delete elimina una incidencia por ID a través del servicio CRUD (puerto 8081).
func (c *IncidenciaController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		c.Ctx.Output.SetStatus(http.StatusBadRequest)
		c.Data["json"] = map[string]interface{}{"Success": false, "Message": "ID de incidencia inválido"}
		_ = c.ServeJSON()
		return
	}

	client := &http.Client{}
	req, errReq := http.NewRequest(http.MethodDelete, crudIncidenciaURL+"/"+strconv.Itoa(id), nil)
	if errReq != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.Data["json"] = map[string]interface{}{"Success": false, "Message": "Error en la petición"}
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