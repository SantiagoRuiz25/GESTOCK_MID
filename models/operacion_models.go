package models

type Operacion struct {
	Id_operacion int `json:"id_operacion"`
	tipo_operacion string `json:"tipo_operacion"`
	cantidad float64 `json:"cantidad"`
	fecha string `json:"fecha"`
	Id_responsable int `json:"id_responsable"`
}