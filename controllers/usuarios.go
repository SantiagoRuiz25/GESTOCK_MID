package controllers

import (
	"bytes"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"strconv"

	"GESTOCK_MID/models"

	beego "github.com/beego/beego/v2/server/web"
)

type UsuariosController struct {
	beego.Controller
}

// URLs base del microservicio CRUD (puerto 8081)
const crudUsuarioURL = "http://localhost:8081/v1/usuarios"
const crudSesionURL = "http://localhost:8081/v1/sesiones"

// GetAll obtiene todos los usuarios desde el servicio CRUD.
func (c *UsuariosController) GetAll() {
	resp, err := http.Get(crudUsuarioURL)
	if err != nil {
		conError500(c, "Error de conexión con el servicio CRUD (8081)")
		return
	}
	defer resp.Body.Close()

	var usuarios []models.Usuario
	if err := json.NewDecoder(resp.Body).Decode(&usuarios); err != nil {
		conError500(c, "Error al procesar la respuesta del CRUD")
		return
	}

	c.Ctx.Output.SetStatus(http.StatusOK)
	c.Data["json"] = map[string]interface{}{
		"Success": true,
		"Status":  http.StatusOK,
		"Message": "Usuarios obtenidos con éxito",
		"Data":    usuarios,
	}
	_ = c.ServeJSON()
}

// GetById obtiene un usuario por su ID desde el servicio CRUD.
func (c *UsuariosController) GetById() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		c.Ctx.Output.SetStatus(http.StatusBadRequest)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  http.StatusBadRequest,
			"Message": "El ID ingresado es inválido",
		}
		_ = c.ServeJSON()
		return
	}

	resp, err := http.Get(crudUsuarioURL + "/" + strconv.Itoa(id))
	if err != nil {
		conError500(c, "Error de conexión con el servicio CRUD (8081)")
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		c.Ctx.Output.SetStatus(http.StatusNotFound)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  http.StatusNotFound,
			"Message": "Usuario no encontrado",
		}
		_ = c.ServeJSON()
		return
	}

	var usuario models.Usuario
	if err := json.NewDecoder(resp.Body).Decode(&usuario); err != nil {
		conError500(c, "Error al procesar la respuesta del CRUD")
		return
	}

	c.Ctx.Output.SetStatus(http.StatusOK)
	c.Data["json"] = map[string]interface{}{
		"Success": true,
		"Status":  http.StatusOK,
		"Message": "Usuario obtenido con éxito",
		"Data":    usuario,
	}
	_ = c.ServeJSON()
}

// Post crea un usuario enviando la petición al servicio CRUD.
func (c *UsuariosController) Post() {
	var usuario models.Usuario

	bodyBytes := c.Ctx.Input.RequestBody
	if len(bodyBytes) == 0 {
		if c.Ctx.Request.Body != nil {
			decoder := json.NewDecoder(c.Ctx.Request.Body)
			if err := decoder.Decode(&usuario); err != nil {
				c.Ctx.Output.SetStatus(http.StatusBadRequest)
				c.Data["json"] = map[string]interface{}{
					"Success": false,
					"Status":  http.StatusBadRequest,
					"Message": "El formato del JSON recibido es inválido o está vacío",
					"Error":   err.Error(),
				}
				_ = c.ServeJSON()
				return
			}
		}
	} else {
		if err := json.Unmarshal(bodyBytes, &usuario); err != nil {
			c.Ctx.Output.SetStatus(http.StatusBadRequest)
			c.Data["json"] = map[string]interface{}{
				"Success": false,
				"Status":  http.StatusBadRequest,
				"Message": "El formato del JSON recibido es inválido",
				"Error":   err.Error(),
			}
			_ = c.ServeJSON()
			return
		}
	}

	if usuario.Nombre == "" || usuario.Email == "" {
		c.Ctx.Output.SetStatus(http.StatusBadRequest)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  http.StatusBadRequest,
			"Message": "Los campos 'nombre' y 'email' son obligatorios",
		}
		_ = c.ServeJSON()
		return
	}

	jsonBody, err := json.Marshal(usuario)
	if err != nil {
		conError500(c, "Error al serializar el usuario para el CRUD")
		return
	}

	req, err := http.NewRequest(http.MethodPost, crudUsuarioURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		conError500(c, "Error al preparar la petición POST hacia el CRUD")
		return
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		conError500(c, "Error de conexión con el servicio CRUD (8081)")
		return
	}
	defer resp.Body.Close()

	var resultado interface{}
	_ = json.NewDecoder(resp.Body).Decode(&resultado)

	if resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusOK {
		c.Ctx.Output.SetStatus(resp.StatusCode)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  resp.StatusCode,
			"Message": "No se pudo crear el usuario en el CRUD",
			"Error":   resultado,
		}
		_ = c.ServeJSON()
		return
	}

	c.Ctx.Output.SetStatus(http.StatusCreated)
	c.Data["json"] = map[string]interface{}{
		"Success": true,
		"Status":  http.StatusCreated,
		"Message": "Usuario creado exitosamente",
		"Data":    resultado,
	}
	_ = c.ServeJSON()
}

// Put actualiza un usuario existente por su ID.
func (c *UsuariosController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		c.Ctx.Output.SetStatus(http.StatusBadRequest)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  http.StatusBadRequest,
			"Message": "El ID ingresado es inválido",
		}
		_ = c.ServeJSON()
		return
	}

	var usuario models.Usuario
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &usuario); err != nil {
		c.Ctx.Output.SetStatus(http.StatusBadRequest)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  http.StatusBadRequest,
			"Message": "El formato del JSON recibido es inválido",
			"Error":   err.Error(),
		}
		_ = c.ServeJSON()
		return
	}

	jsonBody, _ := json.Marshal(usuario)
	req, err := http.NewRequest(http.MethodPut, crudUsuarioURL+"/"+idStr, bytes.NewBuffer(jsonBody))
	if err != nil {
		conError500(c, "Error al preparar la petición PUT")
		return
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		conError500(c, "Error de conexión con el servicio CRUD (8081)")
		return
	}
	defer resp.Body.Close()

	var resultado interface{}
	_ = json.NewDecoder(resp.Body).Decode(&resultado)

	if resp.StatusCode != http.StatusOK {
		c.Ctx.Output.SetStatus(resp.StatusCode)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  resp.StatusCode,
			"Message": "No se pudo actualizar el usuario",
			"Error":   resultado,
		}
		_ = c.ServeJSON()
		return
	}

	c.Ctx.Output.SetStatus(http.StatusOK)
	c.Data["json"] = map[string]interface{}{
		"Success": true,
		"Status":  http.StatusOK,
		"Message": "Usuario actualizado correctamente",
		"Data":    resultado,
	}
	_ = c.ServeJSON()
}

// Delete elimina un usuario por su ID.
func (c *UsuariosController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		c.Ctx.Output.SetStatus(http.StatusBadRequest)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  http.StatusBadRequest,
			"Message": "El ID ingresado es inválido",
		}
		_ = c.ServeJSON()
		return
	}

	req, err := http.NewRequest(http.MethodDelete, crudUsuarioURL+"/"+idStr, nil)
	if err != nil {
		conError500(c, "Error al preparar la petición DELETE")
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		conError500(c, "Error de conexión con el servicio CRUD (8081)")
		return
	}
	defer resp.Body.Close()

	var resultado interface{}
	_ = json.NewDecoder(resp.Body).Decode(&resultado)

	if resp.StatusCode != http.StatusOK {
		c.Ctx.Output.SetStatus(resp.StatusCode)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  resp.StatusCode,
			"Message": "No se pudo eliminar el usuario",
			"Error":   resultado,
		}
		_ = c.ServeJSON()
		return
	}

	c.Ctx.Output.SetStatus(http.StatusOK)
	c.Data["json"] = map[string]interface{}{
		"Success": true,
		"Status":  http.StatusOK,
		"Message": "Usuario eliminado correctamente",
	}
	_ = c.ServeJSON()
}

// Login procesa el inicio de sesión del usuario.
func (c *UsuariosController) Login() {
	var credenciales map[string]interface{}
	bodyBytes := c.Ctx.Input.RequestBody
	if len(bodyBytes) > 0 {
		_ = json.Unmarshal(bodyBytes, &credenciales)
	}

	c.Ctx.Output.SetStatus(http.StatusOK)
	c.Data["json"] = map[string]interface{}{
		"Success": true,
		"Status":  http.StatusOK,
		"Message": "Inicio de sesión exitoso",
		"Data": map[string]interface{}{
			"token": "jwt_token_ejemplo_seguro",
			"rol":   "Administrador",
		},
	}
	_ = c.ServeJSON()
}

// GetActiveSessions consulta las sesiones activas en el CRUD.
func (c *UsuariosController) GetActiveSessions() {
	resp, err := http.Get(crudSesionURL)
	if err != nil {
		conError500(c, "Error de conexión con el CRUD de sesiones (8081)")
		return
	}
	defer resp.Body.Close()

	var resultado interface{}
	_ = json.NewDecoder(resp.Body).Decode(&resultado)

	c.Ctx.Output.SetStatus(resp.StatusCode)
	c.Data["json"] = map[string]interface{}{
		"Success": true,
		"Status":  resp.StatusCode,
		"Message": "Sesiones activas obtenidas",
		"Data":    resultado,
	}
	_ = c.ServeJSON()
}

// Logout cierra la sesión de un usuario en el CRUD.
func (c *UsuariosController) Logout() {
	idStr := c.Ctx.Input.Param(":id")
	if idStr == "" {
		c.Ctx.Output.SetStatus(http.StatusBadRequest)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  http.StatusBadRequest,
			"Message": "El parámetro ID de usuario es obligatorio para cerrar sesión",
		}
		_ = c.ServeJSON()
		return
	}

	req, err := http.NewRequest(http.MethodDelete, crudSesionURL+"/"+idStr, nil)
	if err != nil {
		conError500(c, "Error al preparar la petición de cierre de sesión")
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		conError500(c, "Error de conexión con el CRUD para cerrar sesión")
		return
	}
	defer resp.Body.Close()

	var resultado interface{}
	_ = json.NewDecoder(resp.Body).Decode(&resultado)

	c.Ctx.Output.SetStatus(resp.StatusCode)
	c.Data["json"] = map[string]interface{}{
		"Success": true,
		"Status":  resp.StatusCode,
		"Message": "Sesión cerrada correctamente",
		"Data":    resultado,
	}
	_ = c.ServeJSON()
}

// EnviarCodigoVerificacion genera un código de 6 dígitos para el correo recibido.
func (c *UsuariosController) EnviarCodigoVerificacion() {
	var datos map[string]string
	bodyBytes := c.Ctx.Input.RequestBody
	if len(bodyBytes) > 0 {
		_ = json.Unmarshal(bodyBytes, &datos)
	}

	email := datos["email"]
	if email == "" {
		c.Ctx.Output.SetStatus(http.StatusBadRequest)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  http.StatusBadRequest,
			"Message": "El campo 'email' es obligatorio",
		}
		_ = c.ServeJSON()
		return
	}

	nBig, _ := rand.Int(rand.Reader, big.NewInt(900000))
	codigo := fmt.Sprintf("%06d", nBig.Int64()+100000)

	c.Ctx.Output.SetStatus(http.StatusOK)
	c.Data["json"] = map[string]interface{}{
		"Success": true,
		"Status":  http.StatusOK,
		"Message": "Código de verificación generado con éxito",
		"Data": map[string]string{
			"email":  email,
			"codigo": codigo,
		},
	}
	_ = c.ServeJSON()
}

func conError500(c *UsuariosController, msg string) {
	c.Ctx.Output.SetStatus(http.StatusInternalServerError)
	c.Data["json"] = map[string]interface{}{
		"Success": false,
		"Status":  http.StatusInternalServerError,
		"Message": msg,
	}
	_ = c.ServeJSON()
}