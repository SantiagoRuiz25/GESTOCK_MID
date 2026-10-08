package models

type Operacion struct {
	IdOperacion int `json:"id_operacion"`
	TipoOperacion string `json:"tipo_operacion"`
	Cantidad float64 `json:"cantidad"`
	Fecha string `json:"fecha"`
	IdResponsable int `json:"id_responsable"`
}