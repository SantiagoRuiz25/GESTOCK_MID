package models

import (
	"github.com/beego/beego/v2/client/orm"
)

type Incidencia struct {
	IdIncidencia int `json:"id_incidencia"`
	Descripcion string `json:"descripcion"`
	Estado string `json:"estado"`
	Fecha_creacion string `json:"fecha_creacion"`
	IdUsuario int `json:"id_usuario"`
}

// Consultar

func GetAllIncidencias() ([]Incidencia, error) {
	o := orm.NewOrm()
	var incidencias []Incidencia
	_, err := o.QueryTable(new(Incidencia)).All(&incidencias)
	return incidencias, err
}

// Crear una incidencia
func (i *Incidencia) Create() error {
	o := orm.NewOrm()
	_, err := o.Insert(i)
	return err
}



