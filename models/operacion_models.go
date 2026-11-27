package models

import (
	"github.com/beego/beego/v2/client/orm"
)

type Operacion struct {
	IdOperacion int    `orm:"column(id_operacion);pk;auto" json:"id_operacion"`
	Tipo        string `orm:"column(tipo)" json:"tipo"`
	Cantidad    int    `orm:"column(cantidad)" json:"cantidad"`
	Fecha       string `orm:"column(fecha)" json:"fecha"`
	IdProducto  int    `orm:"column(id_producto)" json:"id_producto"`
	IdBodega    int    `orm:"column(id_bodega)" json:"id_bodega"`
}

// TableName define el nombre exacto de la tabla en PostgreSQL
func (t *Operacion) TableName() string {
	return "operacion"
}

// init registra el modelo automáticamente en el ORM de Beego
func init() {
	orm.RegisterModel(new(Operacion))
}

// GetAllOperaciones obtiene todas las operaciones registradas
func GetAllOperaciones() ([]Operacion, error) {
	o := orm.NewOrm()
	var operaciones []Operacion
	_, err := o.QueryTable(new(Operacion)).All(&operaciones)
	return operaciones, err
}

// Create inserta una nueva operación en la base de datos
func (op *Operacion) Create() error {
	o := orm.NewOrm()
	_, err := o.Insert(op)
	return err
}

// DeleteOperacion elimina una operación según su ID
func DeleteOperacion(id int) error {
	o := orm.NewOrm()
	operacion := Operacion{IdOperacion: id}
	_, err := o.Delete(&operacion)
	return err
}