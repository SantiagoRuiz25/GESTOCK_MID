package models

import (
	"errors"
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

func init() {
	orm.RegisterModel(new(Mantenimiento))
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

func UpdateMantenimiento(m *Mantenimiento) error {
	o := orm.NewOrm()
	// Verificamos si el registro existe primero
	mantenimientoExistente := Mantenimiento{ID: m.ID}
	if err := o.Read(&mantenimientoExistente); err != nil {
		return errors.New("mantenimiento no encontrado")
	}

	// Actualizamos los campos
	mantenimientoExistente.EquipoID = m.EquipoID
	mantenimientoExistente.Tipo = m.Tipo
	mantenimientoExistente.Detalles = m.Detalles
	mantenimientoExistente.FechaProgramada = m.FechaProgramada

	_, err := o.Update(&mantenimientoExistente, "EquipoID", "Tipo", "Detalles", "FechaProgramada")
	return err
}