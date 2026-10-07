package models

type Incidencia struct {
	Id_incidencia int `json:"id_incidencia"`
	descripcion string `json:"descripcion"`
	estado string `json:"estado"`
	fecha_creacion string `json:"fecha_creacion"`
	Id_usuario int `json:"id_usuario"`
}