package models

import (
	"github.com/beego/beego/v2/client/orm"
)

type Mantenimiento struct {
	IdMantenimiento int    `orm:"column(id_mantenimiento);pk;auto" json:"id_mantenimiento"`
	Tipo            string `orm:"column(tipo)" json:"tipo"`
	Descripcion     string `orm:"column(descripcion)" json:"descripcion"`
	Fecha           string `orm:"column(fecha)" json:"fecha"`
	IdBodega        int    `orm:"column(id_bodega)" json:"id_bodega"`
}

// TableName define el nombre exacto de la tabla en PostgreSQL
func (t *Mantenimiento) TableName() string {
	return "mantenimiento"
}

// init registra el modelo automáticamente en el ORM de Beego
func init() {
	orm.RegisterModel(new(Mantenimiento))
}

// GetAllMantenimientos obtiene todos los registros de la tabla mantenimiento
func GetAllMantenimientos() ([]Mantenimiento, error) {
	o := orm.NewOrm()
	var mantenimientos []Mantenimiento
	_, err := o.QueryTable(new(Mantenimiento)).All(&mantenimientos)
	return mantenimientos, err
}

// Create inserta un nuevo registro de mantenimiento
func (m *Mantenimiento) Create() error {
	o := orm.NewOrm()
	_, err := o.Insert(m)
	return err
}

// Update actualiza un mantenimiento existente por su ID
func (m *Mantenimiento) Update(id int) error {
	o := orm.NewOrm()
	mantenimiento := Mantenimiento{IdMantenimiento: id}
	if err := o.Read(&mantenimiento); err == nil {
		mantenimiento.Tipo = m.Tipo
		mantenimiento.Descripcion = m.Descripcion
		mantenimiento.Fecha = m.Fecha
		mantenimiento.IdBodega = m.IdBodega
		_, err := o.Update(&mantenimiento)
		return err
	} else {
		return err
	}
}