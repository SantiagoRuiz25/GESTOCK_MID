package services

import (
	"GESTOCK_MID/models"
	"github.com/beego/beego/v2/client/orm"
)

func GetAllAuditoriasService() ([]models.Auditoria, error) {
	o := orm.NewOrm()
	var auditorias []models.Auditoria
	_, err := o.QueryTable(new(models.Auditoria)).All(&auditorias)
	return auditorias, err
}

func CreateAuditoriaService(a *models.Auditoria) error {
	o := orm.NewOrm()
	_, err := o.Insert(a)
	return err
}

func UpdateAuditoriaService(a *models.Auditoria) error {
	o := orm.NewOrm()
	_, err := o.Update(a)
	return err
}

func DeleteAuditoriaService(id int) error {
	o := orm.NewOrm()
	_, err := o.Delete(&models.Auditoria{IdAuditoria: id})
	return err
}