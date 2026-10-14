package controllers

import (
	"encoding/json"
	"strconv"
	"GESTOCK_MID/models"
	beego "github.com/beego/beego/v2/server/web"
)

type UsuariosController struct {
	beego.Controller
}

var listaUsuarios = []models.Usuario{
	{
		ID:     1,
		Nombre: "Maicol Nore",
		Email:  "maicol@email.com",
		Rol:    "Admin",
		Activo: true,
	},
}

func (c *UsuariosController) GetAll() {
	c.Data["json"] = map[string]interface{}{
		"Success": true,
		"Data":    listaUsuarios,
	}
	_ = c.ServeJSON()
}

func (c *UsuariosController) Post() {
	var usuario models.Usuario
	if err := json.NewDecoder(c.Ctx.Request.Body).Decode(&usuario); err != nil {
		c.Data["json"] = map[string]interface{}{"Success": false, "Message": "Error al leer el JSON"}
		_ = c.ServeJSON()
		return
	}
	if usuario.ID == 0 {
		usuario.ID = len(listaUsuarios) + 1
	}
	listaUsuarios = append(listaUsuarios, usuario)
	c.Data["json"] = map[string]interface{}{"Success": true, "Message": "Usuario creado exitosamente", "Data": usuario}
	_ = c.ServeJSON()
}

func (c *UsuariosController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.Data["json"] = map[string]interface{}{"Success": false, "Message": "ID inválido"}
		_ = c.ServeJSON()
		return
	}

	var usuarioActualizado models.Usuario
	if err := json.NewDecoder(c.Ctx.Request.Body).Decode(&usuarioActualizado); err != nil {
		c.Data["json"] = map[string]interface{}{"Success": false, "Message": "JSON inválido"}
		_ = c.ServeJSON()
		return
	}

	encontrado := false
	for i, u := range listaUsuarios {
		if u.ID == id {
			if usuarioActualizado.Nombre != "" {
				listaUsuarios[i].Nombre = usuarioActualizado.Nombre
			}
			if usuarioActualizado.Email != "" {
				listaUsuarios[i].Email = usuarioActualizado.Email
			}
			if usuarioActualizado.Rol != "" {
				listaUsuarios[i].Rol = usuarioActualizado.Rol
			}
			listaUsuarios[i].Activo = usuarioActualizado.Activo
			encontrado = true
			break
		}
	}

	if !encontrado {
		c.Data["json"] = map[string]interface{}{"Success": false, "Message": "Usuario no encontrado"}
		_ = c.ServeJSON()
		return
	}

	c.Data["json"] = map[string]interface{}{"Success": true, "Message": "Usuario actualizado exitosamente"}
	_ = c.ServeJSON()
}

func (c *UsuariosController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.Data["json"] = map[string]interface{}{"Success": false, "Message": "ID inválido"}
		_ = c.ServeJSON()
		return
	}

	indice := -1
	for i, u := range listaUsuarios {
		if u.ID == id {
			indice = i
			break
		}
	}

	if indice == -1 {
		c.Data["json"] = map[string]interface{}{"Success": false, "Message": "Usuario no encontrado"}
		_ = c.ServeJSON()
		return
	}

	listaUsuarios = append(listaUsuarios[:indice], listaUsuarios[indice+1:]...)
	c.Data["json"] = map[string]interface{}{"Success": true, "Message": "Usuario eliminado exitosamente"}
	_ = c.ServeJSON()
}

func (c *UsuariosController) Login() {
	c.Data["json"] = map[string]interface{}{"Success": true, "Message": "Login exitoso"}
	_ = c.ServeJSON()
}

func (c *UsuariosController) GetActiveSessions() {
	c.Data["json"] = map[string]interface{}{"Success": true, "Data": []string{}}
	_ = c.ServeJSON()
}

func (c *UsuariosController) Logout() {
	c.Data["json"] = map[string]interface{}{"Success": true, "Message": "Sesión cerrada"}
	_ = c.ServeJSON()
}