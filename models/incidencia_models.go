package models

import (
	"GESTOCK_MID/db"
)

type Incidencia struct {
	IdIncidencia int `json:"id_incidencia"`
	Descripcion string `json:"descripcion"`
	Estado string `json:"estado"`
	Fecha_creacion string `json:"fecha_creacion"`
	IdUsuario int `json:"id_usuario"`
}

// Consultar

func GetAllIncidencias() ([]Incidencia, error) {
	rows, err := db.DB.Query("SELECT id_incidencia, descripcion, estado, fecha_creacion, id_usuario FROM incidencia")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var incidencias []Incidencia
	for rows.Next() {
		var i Incidencia
		err := rows.Scan(&i.IdIncidencia, &i.Descripcion, &i.Estado, &i.Fecha_creacion, &i.IdUsuario)
		if err != nil {
			return nil, err
		}
		incidencias = append(incidencias, i)
	}
	return incidencias, nil
}

// Crear

func (i *Incidencia) Create() error {
	_, err := db.DB.Exec("INSERT INTO incidencia (descripcion, estado, fecha_creacion, id_usuario) VALUES (?, ?, ?, ?)", i.Descripcion, i.Estado, i.Fecha_creacion, i.IdUsuario)
	return err
}
