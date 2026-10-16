package models

import (
	"GESTOCK_MID/db"
)

type Mantenimiento struct {
	ID              int    `json:"id"`
	EquipoID        int    `json:"equipo_id"`
	Tipo            string `json:"tipo"`
	Detalles        string `json:"detalles"`
	FechaProgramada string `json:"fecha_programada"`
}

// 1. Obtener todos los mantenimientos
func GetAllMantenimientos() ([]Mantenimiento, error) {
	// Corregido: id_mantenimiento en lugar de id
	rows, err := db.DB.Query("SELECT id_mantenimiento, equipo_id, tipo, detalles, fecha_programada FROM mantenimiento")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var mantenimientos []Mantenimiento
	for rows.Next() {
		var m Mantenimiento
		err := rows.Scan(&m.ID, &m.EquipoID, &m.Tipo, &m.Detalles, &m.FechaProgramada)
		if err != nil {
			return nil, err
		}
		mantenimientos = append(mantenimientos, m)
	}
	return mantenimientos, nil
}

// 2. Crear un mantenimiento
func (m *Mantenimiento) Create() error {
	query := "INSERT INTO mantenimiento (equipo_id, tipo, detalles, fecha_programada) VALUES ($1, $2, $3, $4) RETURNING id_mantenimiento"
	return db.DB.QueryRow(query, m.EquipoID, m.Tipo, m.Detalles, m.FechaProgramada).Scan(&m.ID)
}

// 3. Actualizar un mantenimiento
func UpdateMantenimiento(m *Mantenimiento) error {
	query := "UPDATE mantenimiento SET equipo_id = $1, tipo = $2, detalles = $3, fecha_programada = $4 WHERE id_mantenimiento = $5"
	_, err := db.DB.Exec(query, m.EquipoID, m.Tipo, m.Detalles, m.FechaProgramada, m.ID)
	return err
}

// 4. Eliminar un mantenimiento por ID
func DeleteMantenimiento(id int) error {
	query := "DELETE FROM mantenimiento WHERE id_mantenimiento = $1"
	_, err := db.DB.Exec(query, id)
	return err
}