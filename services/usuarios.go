package services

import (
	"errors"
	"time"

	"GESTOCK_MID/data"
	"GESTOCK_MID/models"
)

func toModelUsuario(u data.Usuario) models.Usuario {
	return models.Usuario{
		ID:              u.ID,
		Nombre:          u.Nombre,
		Email:           u.Email,
		PasswordHash:    u.PasswordHash,
		Rol:             u.Rol,
		RolID:           u.RolID,
		Activo:          u.Activo,
		HorasTrabajadas: u.HorasTrabajadas,
		EmpresaID:       u.EmpresaID,
		CreadoEn:        u.CreadoEn,
		ActualizadoEn:   u.ActualizadoEn,
	}
}

func GetAllUsuariosService() ([]models.Usuario, error) {
	result := make([]models.Usuario, 0, len(data.Usuarios))
	for _, u := range data.Usuarios {
		result = append(result, toModelUsuario(u))
	}
	return result, nil
}

func GetUsuarioByIdService(id int) (models.Usuario, error) {
	for _, u := range data.Usuarios {
		if u.ID == id {
			return toModelUsuario(u), nil
		}
	}
	return models.Usuario{}, errors.New("Usuario no encontrado")
}

func CreateUsuarioService(usuario *models.Usuario) error {
	if usuario == nil {
		return errors.New("Datos de usuario inválidos")
	}

	nextID := 1
	for _, u := range data.Usuarios {
		if u.ID >= nextID {
			nextID = u.ID + 1
		}
	}

	now := time.Now()
	usuario.ID = nextID
	if usuario.CreadoEn.IsZero() {
		usuario.CreadoEn = now
	}
	usuario.ActualizadoEn = now
	if !usuario.Activo {
		usuario.Activo = true
	}

	nuevo := data.Usuario{
		ID:              usuario.ID,
		Nombre:          usuario.Nombre,
		Email:           usuario.Email,
		PasswordHash:    usuario.PasswordHash,
		Rol:             usuario.Rol,
		RolID:           usuario.RolID,
		Activo:          usuario.Activo,
		HorasTrabajadas: usuario.HorasTrabajadas,
		EmpresaID:       usuario.EmpresaID,
		CreadoEn:        usuario.CreadoEn,
		ActualizadoEn:   usuario.ActualizadoEn,
	}

	data.Usuarios = append(data.Usuarios, nuevo)
	return nil
}

func UpdateUsuarioService(id int, usuario *models.Usuario) error {
	if usuario == nil {
		return errors.New("Datos de usuario inválidos")
	}

	for i, u := range data.Usuarios {
		if u.ID == id {
			now := time.Now()
			usuario.ID = id
			usuario.ActualizadoEn = now

			data.Usuarios[i].Nombre = usuario.Nombre
			data.Usuarios[i].Email = usuario.Email
			if usuario.PasswordHash != "" {
				data.Usuarios[i].PasswordHash = usuario.PasswordHash
			}
			data.Usuarios[i].Rol = usuario.Rol
			data.Usuarios[i].RolID = usuario.RolID
			data.Usuarios[i].Activo = usuario.Activo
			data.Usuarios[i].HorasTrabajadas = usuario.HorasTrabajadas
			data.Usuarios[i].EmpresaID = usuario.EmpresaID
			data.Usuarios[i].ActualizadoEn = now

			*usuario = toModelUsuario(data.Usuarios[i])
			return nil
		}
	}

	return errors.New("Usuario no encontrado")
}