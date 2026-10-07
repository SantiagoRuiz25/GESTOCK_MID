package models

type Auditoria struct {
	Id_auditoria int    `json:"id_auditoria"`
	accion	  string `json:"accion"`
	tabla_afectada string `json:"tabla_afectada"`
	Id_usuario int    `json:"id_usuario"`
	fecha_hora string `json:"fecha_hora"`
	detalles_cambios string `json:"detalles_cambios"`
}