package models

import (
	"github.com/beego/beego/v2/client/orm"
)

type Mantenimiento struct {
	ID              int    `json:"id"`
	EquipoID        int    `json:"equipo_id"`
	Tipo            string `json:"tipo"`
	Detalles        string `json:"detalles"`
	FechaProgramada string `json:"fecha_programada"`
}

// Consultar todos los mantenimientos

func GetAllMantenimientos() ([]Mantenimiento, error) {
	o := orm.NewOrm()
	var mantenimientos []Mantenimiento
	_, err := o.QueryTable(new(Mantenimiento)).All(&mantenimientos)
	return mantenimientos, err
}

// Crear un nuevo mantenimiento
func (m *Mantenimiento) Create() error {
	o := orm.NewOrm()
	_, err := o.Insert(m)
	return err
}

// Actualizar un mantenimiento existente
func UpdateMantenimiento(m *Mantenimiento) error {
	o := orm.NewOrm()
	_, err := o.Update(m)
	return err
}

// Eliminar un mantenimiento por ID
func DeleteMantenimiento(id int) error {
	o := orm.NewOrm()
	_, err := o.Delete(&Mantenimiento{ID: id})
	return err
}

