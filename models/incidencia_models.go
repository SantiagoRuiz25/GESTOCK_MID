package models

import (
	"github.com/beego/beego/v2/client/orm"
)

type Incidencia struct {
	IdIncidencia   int    `orm:"column(id_incidencia);pk;auto" json:"id_incidencia"`
	Descripcion    string `orm:"column(descripcion)" json:"descripcion"`
	Estado         string `orm:"column(estado)" json:"estado"`
	Fecha_creacion string `orm:"column(fecha_creacion)" json:"fecha_creacion"`
	IdUsuario      int    `orm:"column(id_usuario)" json:"id_usuario"`
}

// Define el nombre exacto de la tabla en tu base de datos de pgAdmin
func (t *Incidencia) TableName() string {
	return "incidencia"
}

// Registra el modelo automáticamente al iniciar
func init() {
	orm.RegisterModel(new(Incidencia))
}

// Obtener todas las incidencias
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

// Actualizar una incidencia
func (i *Incidencia) Update() error {
	o := orm.NewOrm()
	_, err := o.Update(i)
	return err
}

// Eliminar una incidencia por ID
func DeleteIncidencia(id int) error {
	o := orm.NewOrm()
	_, err := o.Delete(&Incidencia{IdIncidencia: id})
	return err
}