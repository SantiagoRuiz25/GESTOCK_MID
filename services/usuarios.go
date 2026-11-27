package services

import (
	"errors"
	"time"

	"GESTOCK_MID/models"

	"github.com/beego/beego/v2/client/orm"
)

// GetAllUsuariosService obtiene todos los usuarios desde la base de datos.
func GetAllUsuariosService() ([]models.Usuario, error) {
	o := orm.NewOrm()
	var usuarios []models.Usuario
	_, err := o.QueryTable(new(models.Usuario)).All(&usuarios)
	if err != nil {
		return nil, err
	}
	return usuarios, nil
}

// GetUsuarioByIdService busca un usuario por su ID.
func GetUsuarioByIdService(id int) (*models.Usuario, error) {
	o := orm.NewOrm()
	usuario := &models.Usuario{ID: id}
	if err := o.Read(usuario); err != nil {
		if err == orm.ErrNoRows {
			return nil, errors.New("usuario no encontrado")
		}
		return nil, err
	}
	return usuario, nil
}

// CreateUsuarioService crea un nuevo usuario en la base de datos.
func CreateUsuarioService(usuario *models.Usuario) error {
	if usuario == nil {
		return errors.New("datos de usuario inválidos")
	}

	now := time.Now()
	if usuario.CreadoEn.IsZero() {
		usuario.CreadoEn = now
	}
	usuario.ActualizadoEn = now

	o := orm.NewOrm()
	_, err := o.Insert(usuario)
	return err
}

// UpdateUsuarioService actualiza los datos de un usuario existente.
func UpdateUsuarioService(id int, usuario *models.Usuario) error {
	if usuario == nil {
		return errors.New("datos de usuario inválidos")
	}

	o := orm.NewOrm()
	existente := &models.Usuario{ID: id}
	if err := o.Read(existente); err != nil {
		if err == orm.ErrNoRows {
			return errors.New("usuario no encontrado")
		}
		return err
	}

	usuario.ID = id
	usuario.CreadoEn = existente.CreadoEn
	usuario.ActualizadoEn = time.Now()
	if usuario.PasswordHash == "" {
		usuario.PasswordHash = existente.PasswordHash
	}

	_, err := o.Update(usuario)
	return err
}

// DeleteUsuarioService elimina un usuario por su ID.
func DeleteUsuarioService(id int) error {
	o := orm.NewOrm()
	_, err := o.Delete(&models.Usuario{ID: id})
	return err
}