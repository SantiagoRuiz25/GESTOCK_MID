package models

type Producto struct {
	ID          int     `json:"id"`
	Id_categoria 	int 	`json:"id_categoria"`
	Id_bodega int `json:"id_bodega"`
	Codigo      string  `json:"codigo"`
	Nombre      string  `json:"nombre"`
	Precio      float64 `json:"precio"`
	Stock       int     `json:"stock"`
}