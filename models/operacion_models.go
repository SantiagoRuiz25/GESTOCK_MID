package models

import (
	"github.com/beego/beego/v2/client/orm"
)

type Operacion struct {
	IdOperacion int `json:"id_operacion"`
	TipoOperacion string `json:"tipo_operacion"`
	Cantidad float64 `json:"cantidad"`
	Fecha string `json:"fecha"`
	IdResponsable int `json:"id_responsable"`
}

// Consultar todas las operaciones

func GetAllOperaciones() ([]Operacion, error) {
	o := orm.NewOrm()
	var operaciones []Operacion
	_, err := o.QueryTable(new(Operacion)).All(&operaciones)
	return operaciones, err
}

// Crear una nueva operación

func CreateOperacion(op *Operacion) error {
	o := orm.NewOrm()
	_, err := o.Insert(op)
	return err
}

// Eliminar una operación por ID

func DeleteOperacion(id int) error {
	o := orm.NewOrm()
	_, err := o.Delete(&Operacion{IdOperacion: id})
	return err
}
