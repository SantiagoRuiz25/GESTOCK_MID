package models

import (
	"GESTOCK_MID/db"
)

type Incidencia struct {
	IdIncidencia   int    `json:"id_incidencia"`
	Descripcion    string `json:"descripcion"`
	Estado         string `json:"estado"`
	Fecha_creacion string `json:"fecha_creacion"`
	IdUsuario      int    `json:"id_usuario"`
}

// 1. GET: Buscar / Obtener todas las incidencias
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

// 2. POST: Crear una incidencia
func (i *Incidencia) Create() error {
	query := "INSERT INTO incidencia (descripcion, estado, fecha_creacion, id_usuario) VALUES ($1, $2, $3, $4)"
	_, err := db.DB.Exec(query, i.Descripcion, i.Estado, i.Fecha_creacion, i.IdUsuario)
	return err
}