package models

type Mantenimiento struct {
	IdMantenimiento int `json:"id_mantenimiento"`
	IdEquipo int `json:"id_equipo"`
	Tipo string `json:"tipo"`
	Detalles string `json:"detalles"`
	Fecha_programada string `json:"fecha_programada"`
}