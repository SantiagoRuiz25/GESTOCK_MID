package models

import (
	"github.com/beego/beego/v2/client/orm"
)

type Incidencia struct {
	IdIncidencia   int    `orm:"column(id_incidencia);pk;auto" json:"id_incidencia"`
	Descripcion    string `orm:"column(descripcion)" json:"descripcion"`
	Estado         string `orm:"column(estado)" json:"estado"`
	FechaCreacion  string `orm:"column(fecha_creacion)" json:"fecha_creacion"`
	IdUsuario      int    `orm:"column(id_usuario)" json:"id_usuario"`
}

func (t *Incidencia) TableName() string {
	return "incidencia"
}

func GetAllIncidencias() ([]Incidencia, error) {
	o := orm.NewOrm()
	var incidencias []Incidencia
	_, err := o.QueryTable(new(Incidencia)).All(&incidencias)
	if err != nil {
		return nil, err
	}
	return incidencias, nil
}

func (i *Incidencia) Create() error {
	o := orm.NewOrm()
	id, err := o.Insert(i)
	if err != nil {
		return err
	}
	i.IdIncidencia = int(id)
	return nil
}