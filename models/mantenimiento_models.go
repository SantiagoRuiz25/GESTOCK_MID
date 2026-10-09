package models

import (
	"GESTOCK_MID/db" // Única importación requerida para conectar con la base de datos
)

type Mantenimiento struct {
	ID              int    `json:"id"`
	EquipoID        int    `json:"equipo_id"`
	Tipo            string `json:"tipo"`
	Detalles        string `json:"detalles"`
	FechaProgramada string `json:"fecha_programada"`
}

// Obtener todos los mantenimientos
func GetAllMantenimientos() ([]Mantenimiento, error) {
	// Se usa directamente db.DB que ya está configurado globalmente
	rows, err := db.DB.Query("SELECT id, equipo_id, tipo, detalles, fecha_programada FROM mantenimiento")
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

// Crear un mantenimiento
func (m *Mantenimiento) Create() error {
	query := "INSERT INTO mantenimiento (equipo_id, tipo, detalles, fecha_programada) VALUES ($1, $2, $3, $4) RETURNING id"
	return db.DB.QueryRow(query, m.EquipoID, m.Tipo, m.Detalles, m.FechaProgramada).Scan(&m.ID)
}