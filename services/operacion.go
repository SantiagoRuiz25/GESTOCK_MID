package services

import (
	"GESTOCK_MID/models"
	"github.com/beego/beego/v2/client/orm"
)

func GetAllOperacionesService() ([]models.Operacion, error) {
	o := orm.NewOrm()
	var operaciones []models.Operacion
	_, err := o.QueryTable(new(models.Operacion)).All(&operaciones)
	return operaciones, err
}

func CreateOperacionService(o *models.Operacion) error {
	ormInstance := orm.NewOrm()
	_, err := ormInstance.Insert(o)
	return err
}

func DeleteOperacionService(id int) error {
	ormInstance := orm.NewOrm()
	_, err := ormInstance.Delete(&models.Operacion{IdOperacion: id})
	return err
}