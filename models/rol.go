package models

// Rol estructura para mapear la información de los roles
type Rol struct {
	ID          int    `json:"id_rol"`
	Nombre      string `json:"nombre"`
	Descripcion string `json:"descripcion"`
}

