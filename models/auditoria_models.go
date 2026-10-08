package models

import (
	"GESTOCK_MID/db"
)

type Auditoria struct {
	IdAuditoria int    `json:"id_auditoria"`
	Accion	  string `json:"accion"`
	TablaAfectada string `json:"tabla_afectada"`
	IdUsuario int    `json:"id_usuario"`
	FechaHora string `json:"fecha_hora"`
	DetallesCambios string `json:"detalles_cambios"`
}

func GetAllAuditorias() ([]Auditoria, error) {
	rows, err := db.DB.Query("SELECT id_auditoria, accion, tabla_afectada, id_usuario, fecha_hora, detalles_cambios FROM auditoria")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var auditorias []Auditoria
	for rows.Next() {
		var a Auditoria
		err := rows.Scan(&a.IdAuditoria, &a.Accion, &a.TablaAfectada, &a.IdUsuario, &a.FechaHora, &a.DetallesCambios)
		if err != nil {
			return nil, err
		}
		auditorias = append(auditorias, a)
	}
	return auditorias, nil
}

// crear

func (a *Auditoria) Create() error {
	_, err := db.DB.Exec("INSERT INTO auditoria (id_auditoria, accion, tabla_afectada, id_usuario, fecha_hora, detalles_cambios) VALUES (?, ?, ?, ?, ?, ?)", a.IdAuditoria, a.Accion, a.TablaAfectada, a.IdUsuario, a.FechaHora, a.DetallesCambios)
	return err
}

// actualizar

func (a *Auditoria) Update(id int) error {
	_, err := db.DB.Exec("UPDATE auditoria SET accion = ?, tabla_afectada = ?, id_usuario = ?, fecha_hora = ?, detalles_cambios = ? WHERE id_auditoria = ?", a.Accion, a.TablaAfectada, a.IdUsuario, a.FechaHora, a.DetallesCambios, id)
	return err
}

// eliminar

func (a *Auditoria) Delete(id int) error {
	_, err := db.DB.Exec("DELETE FROM auditoria WHERE id_auditoria = ?", id)
	return err
}
