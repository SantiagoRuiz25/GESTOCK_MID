package models

type Mantenimiento struct {
	Id_mantenimiento int `json:"id_mantenimiento"`
	Id_equipo int `json:"id_equipo"`
	tipo string `json:"tipo"`
	detalles string `json:"detalles"`
	fecha_programada string `json:"fecha_programada"`
}