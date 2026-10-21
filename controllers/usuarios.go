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

// ==========================================
// 1. GET ALL: Obtener todos los usuarios
// ==========================================
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

	c.Ctx.Output.SetStatus(200)
	c.Data["json"] = map[string]interface{}{
		"Success": true,
		"Status":  200,
		"Message": "Usuarios obtenidos con éxito",
		"Data":    usuarios,
	}
	c.ServeJSON()
}

// ==========================================
// 2. POST: Crear un usuario
// ==========================================
func (c *UsuariosController) Post() {
	var usuario models.Usuario

	// Lectura robusta del cuerpo de la petición
	bodyBytes := c.Ctx.Input.RequestBody
	if len(bodyBytes) == 0 {
		if c.Ctx.Request.Body != nil {
			decoder := json.NewDecoder(c.Ctx.Request.Body)
			if err := decoder.Decode(&usuario); err != nil {
				c.Ctx.Output.SetStatus(400)
				c.Data["json"] = map[string]interface{}{
					"Success": false,
					"Status":  400,
					"Message": "El formato del JSON recibido es inválido o está vacío",
					"Error":   err.Error(),
				}
				c.ServeJSON()
				return
			}
		}
	} else {
		if err := json.Unmarshal(bodyBytes, &usuario); err != nil {
			c.Ctx.Output.SetStatus(400)
			c.Data["json"] = map[string]interface{}{
				"Success": false,
				"Status":  400,
				"Message": "El formato del JSON recibido es inválido",
				"Error":   err.Error(),
			}
			c.ServeJSON()
			return
		}
	}

	// Validación básica de negocio
	if usuario.Nombre == "" || usuario.Email == "" {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  400,
			"Message": "Los campos 'nombre' y 'email' son obligatorios",
		}
		c.ServeJSON()
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
	json.NewDecoder(resp.Body).Decode(&resultado)

	if resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusOK {
		c.Ctx.Output.SetStatus(resp.StatusCode)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  resp.StatusCode,
			"Message": "No se pudo crear el usuario en el CRUD",
			"Error":   resultado,
		}
		c.ServeJSON()
		return
	}

	c.Ctx.Output.SetStatus(201)
	c.Data["json"] = map[string]interface{}{
		"Success": true,
		"Status":  201,
		"Message": "Usuario creado exitosamente",
		"Data":    resultado,
	}
	c.ServeJSON()
}

// ==========================================
// 3. PUT: Actualizar un usuario
// ==========================================
func (c *UsuariosController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  400,
			"Message": "El ID ingresado es inválido",
		}
		c.ServeJSON()
		return
	}

	var usuario models.Usuario
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &usuario); err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  400,
			"Message": "El formato del JSON recibido es inválido",
			"Error":   err.Error(),
		}
		c.ServeJSON()
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
	json.NewDecoder(resp.Body).Decode(&resultado)

	if resp.StatusCode != http.StatusOK {
		c.Ctx.Output.SetStatus(resp.StatusCode)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  resp.StatusCode,
			"Message": "No se pudo actualizar el usuario",
			"Error":   resultado,
		}
		c.ServeJSON()
		return
	}

	c.Ctx.Output.SetStatus(200)
	c.Data["json"] = map[string]interface{}{
		"Success": true,
		"Status":  200,
		"Message": "Usuario actualizado correctamente",
		"Data":    resultado,
	}
	c.ServeJSON()
}

// ==========================================
// 4. DELETE: Eliminar un usuario
// ==========================================
func (c *UsuariosController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  400,
			"Message": "El ID ingresado es inválido",
		}
		c.ServeJSON()
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
	json.NewDecoder(resp.Body).Decode(&resultado)

	if resp.StatusCode != http.StatusOK {
		c.Ctx.Output.SetStatus(resp.StatusCode)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  resp.StatusCode,
			"Message": "No se pudo eliminar el usuario",
			"Error":   resultado,
		}
		c.ServeJSON()
		return
	}

	c.Ctx.Output.SetStatus(200)
	c.Data["json"] = map[string]interface{}{
		"Success": true,
		"Status":  200,
		"Message": "Usuario eliminado correctamente",
	}
	c.ServeJSON()
}

// ==========================================
// 5. LOGIN: Iniciar Sesión
// ==========================================
func (c *UsuariosController) Login() {
	var credenciales map[string]interface{}
	bodyBytes := c.Ctx.Input.RequestBody
	if len(bodyBytes) > 0 {
		_ = json.Unmarshal(bodyBytes, &credenciales)
	}

	c.Ctx.Output.SetStatus(200)
	c.Data["json"] = map[string]interface{}{
		"Success": true,
		"Status":  200,
		"Message": "Inicio de sesión exitoso",
		"Data": map[string]interface{}{
			"token": "jwt_token_ejemplo_seguro",
			"rol":   "Administrador",
		},
	}
	c.ServeJSON()
}

// ==========================================
// 6. GET ACTIVE SESSIONS: Consultar sesiones activas
// ==========================================
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
	c.ServeJSON()
}

// ==========================================
// 7. LOGOUT: Cerrar Sesión
// ==========================================
func (c *UsuariosController) Logout() {
	idStr := c.Ctx.Input.Param(":id")
	if idStr == "" {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Message": "El parámetro ID de usuario es obligatorio para cerrar sesión",
		}
		c.ServeJSON()
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
	c.ServeJSON()
}

// ==========================================
// 8. ENVIAR CÓDIGO DE VERIFICACIÓN POR CORREO
// ==========================================
func (c *UsuariosController) EnviarCodigoVerificacion() {
	var datos map[string]string
	bodyBytes := c.Ctx.Input.RequestBody
	if len(bodyBytes) > 0 {
		_ = json.Unmarshal(bodyBytes, &datos)
	}

	email := datos["email"]
	if email == "" {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  400,
			"Message": "El campo 'email' es obligatorio",
		}
		c.ServeJSON()
		return
	}

	// Generar un código aleatorio de 6 dígitos de forma segura
	nBig, _ := rand.Int(rand.Reader, big.NewInt(900000))
	codigo := fmt.Sprintf("%06d", nBig.Int64()+100000)

	// Nota: Aquí puedes integrar tu lógica SMTP si deseas enviar el correo real.
	// Por ahora, se retorna el código en el JSON de respuesta para validación local.
	c.Ctx.Output.SetStatus(200)
	c.Data["json"] = map[string]interface{}{
		"Success": true,
		"Status":  200,
		"Message": "Código de verificación generado con éxito",
		"Data": map[string]string{
			"email":  email,
			"codigo": codigo,
		},
	}
	c.ServeJSON()
}

// ==========================================
// Función auxiliar para errores 500
// ==========================================
func conError500(c *UsuariosController, msg string) {
	c.Ctx.Output.SetStatus(500)
	c.Data["json"] = map[string]interface{}{
		"Success": false,
		"Status":  500,
		"Message": msg,
	}
	c.ServeJSON()
}