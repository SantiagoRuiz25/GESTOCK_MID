package models

type Empresas struct {
	Id        int    `json:"id"`
	Nombre    string `json:"nombre"`
	Nit       string `json:"nit"`
	Direccion string `json:"direccion"`
	Telefono  string `json:"telefono"`
	Activo    bool   `json:"activo"`
}