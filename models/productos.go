package models

type Producto struct {
	ID          int     `json:"id"`
	Codigo      string  `json:"codigo"`
	Nombre      string  `json:"nombre"`
	CategoriaId int     `json:"categoria_id"`
	BodegaId    int     `json:"bodega_id"`
	Precio      float64 `json:"precio"`
	Stock       int     `json:"stock"`
}