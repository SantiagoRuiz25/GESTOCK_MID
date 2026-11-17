
package controllers

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"time"

	beego "github.com/beego/beego/v2/server/web"
)

type EmpresasMidController struct {
	beego.Controller
}

const crudEmpresaURL = "http://localhost:8081/v1/empresas"

var clienteEmpresas = &http.Client{
	Timeout: 10 * time.Second,
}

// Función para comunicarse con GESTOCK_CRUD
func (c *EmpresasMidController) enviarCRUD(metodo string, conID bool) {
	url := crudEmpresaURL

	if conID {
		idStr := c.Ctx.Input.Param(":id")
		id, err := strconv.Atoi(idStr)

		if err != nil || id <= 0 {
			c.Ctx.Output.SetStatus(400)
			c.Data["json"] = map[string]interface{}{
				"success": false,
				"message": "ID de empresa inválido",
			}
			c.ServeJSON()
			return
		}

		url += "/" + strconv.Itoa(id)
	}

	var cuerpo io.Reader

	if metodo == http.MethodPost || metodo == http.MethodPut {
		datos := c.Ctx.Input.RequestBody

		if len(datos) == 0 && c.Ctx.Request.Body != nil {
			var err error
			datos, err = io.ReadAll(c.Ctx.Request.Body)
			if err != nil {
				c.Ctx.Output.SetStatus(400)
				c.Data["json"] = map[string]interface{}{
					"success": false,
					"message": "No se pudo leer el JSON",
				}
				c.ServeJSON()
				return
			}
		}

		if !json.Valid(datos) {
			c.Ctx.Output.SetStatus(400)
			c.Data["json"] = map[string]interface{}{
				"success": false,
				"message": "JSON inválido",
			}
			c.ServeJSON()
			return
		}

		cuerpo = bytes.NewReader(datos)
	}

	req, err := http.NewRequest(metodo, url, cuerpo)
	if err != nil {
		c.responderError("Error al preparar la petición al CRUD")
		return
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := clienteEmpresas.Do(req)
	if err != nil {
		c.Ctx.Output.SetStatus(502)
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"message": "No se pudo conectar con GESTOCK_CRUD",
			"error":   err.Error(),
		}
		c.ServeJSON()
		return
	}
	defer resp.Body.Close()

	resultado, err := io.ReadAll(resp.Body)
	if err != nil {
		c.responderError("Error al leer la respuesta del CRUD")
		return
	}

	// Conserva el código HTTP y la respuesta original del CRUD
	c.Ctx.Output.SetStatus(resp.StatusCode)

	tipo := resp.Header.Get("Content-Type")
	if tipo == "" {
		tipo = "application/json"
	}
	c.Ctx.Output.Header("Content-Type", tipo)

	_ = c.Ctx.Output.Body(resultado)
}

func (c *EmpresasMidController) responderError(mensaje string) {
	c.Ctx.Output.SetStatus(500)
	c.Data["json"] = map[string]interface{}{
		"success": false,
		"message": mensaje,
	}
	c.ServeJSON()
}

// GET - Listar empresas
func (c *EmpresasMidController) GetAll() {
	c.enviarCRUD(http.MethodGet, false)
}

// GET - Buscar empresa por ID
func (c *EmpresasMidController) GetById() {
	c.enviarCRUD(http.MethodGet, true)
}

// POST - Crear empresa
func (c *EmpresasMidController) Post() {
	c.enviarCRUD(http.MethodPost, false)
}

// PUT - Actualizar empresa
func (c *EmpresasMidController) Put() {
	c.enviarCRUD(http.MethodPut, true)
}

// DELETE - Eliminar empresa
func (c *EmpresasMidController) Delete() {
	c.enviarCRUD(http.MethodDelete, true)
}
