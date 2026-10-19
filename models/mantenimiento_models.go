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

func init() {
	orm.RegisterModel(new(Mantenimiento))
}

// 1. Obtener todos los mantenimientos (GET)
func GetAllMantenimientos() ([]Mantenimiento, error) {
	o := orm.NewOrm()
	var mantenimientos []Mantenimiento
	_, err := o.QueryTable(new(Mantenimiento)).All(&mantenimientos)
	return mantenimientos, err
}

// 2. Crear un mantenimiento (POST)
func (m *Mantenimiento) Create() error {
	o := orm.NewOrm()
	_, err := o.Insert(m)
	return err
}

// 3. Actualizar un mantenimiento (PUT) - Función Global
func UpdateMantenimiento(m *Mantenimiento) error {
	o := orm.NewOrm()
	_, err := o.Update(m)
	return err
}

// 4. Eliminar un mantenimiento por ID (DELETE)
func DeleteMantenimiento(id int) error {
	o := orm.NewOrm()
	_, err := o.Delete(&Mantenimiento{ID: id})
	return err
}