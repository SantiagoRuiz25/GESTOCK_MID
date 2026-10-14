package services

import (
	"GESTOCK_MID/models"
	"github.com/beego/beego/v2/client/orm"
)

func GetAllIncidenciasService() ([]models.Incidencia, error) {
	o := orm.NewOrm()
	var incidencias []models.Incidencia
	_, err := o.QueryTable(new(models.Incidencia)).All(&incidencias)
	return incidencias, err
}

func CreateIncidenciaService(i *models.Incidencia) error {
	o := orm.NewOrm()
	_, err := o.Insert(i)
	return err
}

func UpdateIncidenciaService(i *models.Incidencia) error {
	o := orm.NewOrm()
	_, err := o.Update(i)
	return err
}

func DeleteIncidenciaService(id int) error {
	o := orm.NewOrm()
	_, err := o.Delete(&models.Incidencia{IdIncidencia: id})
	return err
}