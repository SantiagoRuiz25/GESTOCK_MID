package controllers

import (
	"GESTOCK_MID/models"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)

// buscar todos los usuarios registrados
func GetUsuarios(c *gin.Context) {
	usuarios, err := models.GetAllUsuarios()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, usuarios)
}

// crear
func CreateUsuario(c *gin.Context) {
	var usuario models.Usuario
	if err := c.ShouldBindJSON(&usuario); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := usuario.Create(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, usuario)
}

// actualizar
func UpdateUsuario(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inbalido"})
		return
	}

	var usuarioActualizado models.Usuario
	if err := c.ShouldBindJSON(&usuarioActualizado); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := usuarioActualizado.Update(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Usuario actualizado correctamente"})
}

// eliminar
func DeleteUsuario(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID invalido"})
		return
	}

	var usuario models.Usuario
	if err := usuario.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Usuario eliminado correctamente"})
}

// Iniciar sesión
func LoginUsuario(c *gin.Context) {
	var credenciales struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&credenciales); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	usuario, err := models.Login(credenciales.Email, credenciales.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Credenciales invalidas"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Sesión iniciada corectamente",
		"usuario": usuario,
	})
}

// Obtener sesiones activas
func GetSesionesActivas(c *gin.Context) {
	sesiones, err := models.GetActiveSessions()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, sesiones)
}

// Cerrar sesión
func LogoutUsuario(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de usuario inválido"})
		return
	}

	if err := models.Logout(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Secion cerrada correctamente"})
}