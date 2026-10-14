package services

import (
	"errors"
	"GESTOCK_MID/models"
)

func GetAllUsuariosService()([]models.Usuario, error){
	usuarios, err :=models.GetAllUsuarios()
	if err != nil {
		return nil, err
	}
	return usuarios, nil
}

func GetUsuarioByIdService (id int)(*models.Usuario, error){
	usuario, err := models.GetUsuarioById(id)
	if err != nil {
		return nil, errors.New("Usuario no encontrado")

	}
	return usuario, nil
}

func CreateUsuarioService(usuario *models.Usuario)error{
	err:=usuario.Create()
	if err!=nil{
		return err
	}
	return nil
}

func UpdateUsuarioService(id int, usuario *models.Usuario)



