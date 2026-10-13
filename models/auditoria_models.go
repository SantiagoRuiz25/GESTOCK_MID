 package models

import (
	"github.com/beego/beego/v2/client/orm"
)

type Auditoria struct {
	IdAuditoria     int    `orm:"column(id_auditoria);pk;auto" json:"id_auditoria"`
	Accion          string `orm:"column(accion)" json:"accion"`
	TablaAfectada   string `orm:"column(tabla_afectada)" json:"tabla_afectada"`
	IdUsuario       int    `orm:"column(id_usuario)" json:"id_usuario"`
	FechaHora       string `orm:"column(fecha_hora)" json:"fecha_hora"`
	DetallesCambios string `orm:"column(detalles_cambios)" json:"detalles_cambios"`
}

// Define el nombre exacto de la tabla en tu base de datos de pgAdmin
func (t *Auditoria) TableName() string {
	return "auditoria"
}

// Registra el modelo automáticamente al iniciar
func init() {
	orm.RegisterModel(new(Auditoria))
}

// Obtener todas las auditorías
func GetAllAuditorias() ([]Auditoria, error) {
	o := orm.NewOrm()
	var auditorias []Auditoria
	_, err := o.QueryTable(new(Auditoria)).All(&auditorias)
	return auditorias, err
}

// Crear una auditoría
func (a *Auditoria) Create() error {
	o := orm.NewOrm()
	_, err := o.Insert(a)
	return err
}

// Actualizar una auditoría
func (a *Auditoria) Update() error {
	o := orm.NewOrm()
	_, err := o.Update(a)
	return err
}

// Eliminar una auditoría por ID
func DeleteAuditoria(id int) error {
	o := orm.NewOrm()
	_, err := o.Delete(&Auditoria{IdAuditoria: id})
	return err
}