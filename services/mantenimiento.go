package services

import (
	"GESTOCK_MID/models"
	"github.com/beego/beego/v2/client/orm"
)

func GetAllMantenimientosService() ([]models.Mantenimiento, error) {
	o := orm.NewOrm()
	var mantenimientos []models.Mantenimiento
	_, err := o.QueryTable(new(models.Mantenimiento)).All(&mantenimientos)
	return mantenimientos, err
}

func CreateMantenimientoService(m *models.Mantenimiento) error {
	o := orm.NewOrm()
	_, err := o.Insert(m)
	return err
}

func UpdateMantenimientoService(m *models.Mantenimiento) error {
	o := orm.NewOrm()
	_, err := o.Update(m)
	return err
}

func DeleteMantenimientoService(id int) error {
	o := orm.NewOrm()
	_, err := o.Delete(&models.Mantenimiento{ID: id})
	return err
}