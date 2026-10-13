package models

import (
	"github.com/beego/beego/v2/client/orm"
)

type Mantenimiento struct {
	ID              int    `orm:"column(id);pk;auto" json:"id"`
	EquipoID        int    `orm:"column(equipo_id)" json:"equipo_id"`
	Tipo            string `orm:"column(tipo)" json:"tipo"`
	Detalles        string `orm:"column(detalles)" json:"detalles"`
	FechaProgramada string `orm:"column(fecha_programada)" json:"fecha_programada"`
}

func (t *Mantenimiento) TableName() string {
	return "mantenimiento"
}

func GetAllMantenimientos() ([]Mantenimiento, error) {
	o := orm.NewOrm()
	var mantenimientos []Mantenimiento
	_, err := o.QueryTable(new(Mantenimiento)).All(&mantenimientos)
	if err != nil {
		return nil, err
	}
	return mantenimientos, nil
}

func (m *Mantenimiento) Create() error {
	o := orm.NewOrm()
	id, err := o.Insert(m)
	if err != nil {
		return err
	}
	m.ID = int(id)
	return nil
}