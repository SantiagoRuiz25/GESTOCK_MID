package models

import (
	"errors"
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
	if err != nil {
		return nil, err
	}
	return auditorias, nil
}

// Crear una auditoría
func (a *Auditoria) Create() error {
	o := orm.NewOrm()
	_, err := o.Insert(a)
	return err
}

func (a *Auditoria) Update(id int) error {
	o := orm.NewOrm()
	auditoria := Auditoria{IdAuditoria: id}
	
	if err := o.Read(&auditoria); err == nil {
		auditoria.Accion = a.Accion
		auditoria.TablaAfectada = a.TablaAfectada
		auditoria.IdUsuario = a.IdUsuario
		auditoria.FechaHora = a.FechaHora
		auditoria.DetallesCambios = a.DetallesCambios
		
		_, err := o.Update(&auditoria, "Accion", "TablaAfectada", "IdUsuario", "FechaHora", "DetallesCambios")
		return err
	}
	return errors.New("auditoria no encontrada")
}

func (a *Auditoria) Delete(id int) error {
	o := orm.NewOrm()
	auditoria := Auditoria{IdAuditoria: id}
	
	if err := o.Read(&auditoria); err == nil {
		_, err := o.Delete(&auditoria)
		return err
	}
	return errors.New("auditoria no encontrada")
}

// Eliminar una auditoría por ID
func DeleteAuditoria(id int) error {
	o := orm.NewOrm()
	_, err := o.Delete(&Auditoria{IdAuditoria: id})
	return err
}