package services

import (
	"GESTOCK_MID/models"
	"github.com/beego/beego/v2/client/orm"
)

// Obtener todos los mantenimientos
func GetAllMantenimientosService() ([]models.Mantenimiento, error) {
	o := orm.NewOrm()
	var mantenimientos []models.Mantenimiento
	_, err := o.QueryTable(new(models.Mantenimiento)).All(&mantenimientos)
	return mantenimientos, err
}

// Crear un mantenimiento
func CreateMantenimientoService(m *models.Mantenimiento) error {
	o := orm.NewOrm()
	_, err := o.Insert(m)
	return err
}

// Actualizar un mantenimiento
func UpdateMantenimientoService(m *models.Mantenimiento) error {
	o := orm.NewOrm()
	_, err := o.Update(m)
	return err
}

// Eliminar un mantenimiento por ID
func DeleteMantenimientoService(id int) error {
	o := orm.NewOrm()
	_, err := o.Delete(&models.Mantenimiento{IdMantenimiento: id})
	return err
}