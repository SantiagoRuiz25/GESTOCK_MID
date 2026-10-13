package models

type Categoria struct{
	ID 				int 			`json:"id_categoria"`
	Nombre 			string 			`json:"nombre"`
	Codigo 			int 			`json:"codigo"`
	Descripcion 	string 			`json:"descripcion"`
	Activo 			bool 			`json:"activo"`
	


}


