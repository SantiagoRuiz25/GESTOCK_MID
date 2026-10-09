package models

import (
	"GESTOCK_MID/db"
	"errors"
)

type Usuario struct {
	IdUsuario int `json:"id_usuario"`
	Nombre    string `json:"nombre"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Rol       string `json:"rol"`
}

type SesionActiva struct {
	IdSesion  int `json:"id_sesion"`
	IdUsuario int  `json:"id_usuario"`
	Token     string `json:"token"`
	FechaHora string `json:"fecha_hora"`
}

func GetAllUsuarios() ([]Usuario, error) {
	rows, err := db.DB.Query("SELECT id_usuario, nombre, email, password, rol FROM usuario")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var usuarios []Usuario
	for rows.Next() {
		var u Usuario
		err := rows.Scan(&u.IdUsuario, &u.Nombre, &u.Email, &u.Password, &u.Rol)
		if err != nil {
			return nil, err
		}
		usuarios = append(usuarios, u)
	}
	return usuarios, nil
}

func (u *Usuario) Create() error {
	_, err := db.DB.Exec("INSERT INTO usuario (id_usuario, nombre, email, password, rol) VALUES (?, ?, ?, ?, ?)", u.IdUsuario, u.Nombre, u.Email, u.Password, u.Rol)
	return err
}

func (u *Usuario) Update(id int) error {
	_, err := db.DB.Exec("UPDATE usuario SET nombre = ?, email = ?, password = ?, rol = ? WHERE id_usuario = ?", u.Nombre, u.Email, u.Password, u.Rol, id)
	return err
}

func (u *Usuario) Delete(id int) error {
	_, err := db.DB.Exec("DELETE FROM usuario WHERE id_usuario = ?", id)
	return err
}

// Login 
func Login(email string, password string) (*Usuario, error) {
	var u Usuario
	err := db.DB.QueryRow("SELECT id_usuario, nombre, email, password, rol FROM usuario WHERE email = ? AND password = ?", email, password).
		Scan(&u.IdUsuario, &u.Nombre, &u.Email, &u.Password, &u.Rol)
	
	if err != nil {
		return nil, errors.New("credenciales inválidas o usuario no encontrado")
	}
	return &u, nil
}

// Get para consultar sesiones
func GetActiveSessions() ([]SesionActiva, error) {
	rows, err := db.DB.Query("SELECT id_sesion, id_usuario, token, fecha_hora FROM sesiones_activas")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var sesiones []SesionActiva
	for rows.Next() {
		var s SesionActiva
		err := rows.Scan(&s.IdSesion, &s.IdUsuario, &s.Token, &s.FechaHora)
		if err != nil {
			return nil, err
		}
		sesiones = append(sesiones, s)
	}
	return sesiones, nil
}

// Logout elimina o invalida la sesión activa de un usuario por su ID
func Logout(idUsuario int) error {
	_, err := db.DB.Exec("DELETE FROM sesiones_activas WHERE id_usuario = ?", idUsuario)
	return err
}