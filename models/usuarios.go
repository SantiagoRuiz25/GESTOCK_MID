package models


// Usuario estructura principal con la relación rol_id
type Usuario struct {
	ID     int    `json:"id_usuario"`
	Nombre string `json:"nombre"`
	Email  string `json:"email"`
	RolId  int    `json:"rol_id"` // ID numérico del rol para relacionarlo con la base de datos
	Activo bool   `json:"activo"`
}