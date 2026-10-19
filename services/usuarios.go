package services

import (
	"errors"
	"GESTOCK_MID/models"

	"github.com/beego/beego/v2/client/orm"
)

// Obtener todos los usuarios
func GetAllUsuariosService() ([]models.Usuario, error) {
	o := orm.NewOrm()
	var usuarios []models.Usuario
	_, err := o.QueryTable(new(models.Usuario)).All(&usuarios)
	if err != nil {
		return nil, err
	}
	return usuarios, nil
}

// Obtener un usuario por ID
func GetUsuarioByIdService(id int) (*models.Usuario, error) {
	o := orm.NewOrm()
	usuario := &models.Usuario{ID: id}
	err := o.Read(usuario)
	if err != nil {
		return nil, errors.New("usuario no encontrado")
	}
	return usuario, nil
}

// Crear un nuevo usuario
func CreateUsuarioService(usuario *models.Usuario) error {
	o := orm.NewOrm()
	_, err := o.Insert(usuario)
	return err
}

// Actualizar un usuario existente
func UpdateUsuarioService(id int, usuario *models.Usuario) error {
	o := orm.NewOrm()
	usuario.ID = id
	_, err := o.Update(usuario)
	return err
}

// Eliminar un usuario por ID
func DeleteUsuarioService(id int) error {
	o := orm.NewOrm()
	_, err := o.Delete(&models.Usuario{ID: id})
	if err != nil {
		return err
	}
	return nil
}