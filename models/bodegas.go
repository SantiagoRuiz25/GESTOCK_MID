package models

type Bodegas struct{
	ID int `json:"id_bodega"`
	Id_Usuario int `json:"id_usuario"`
	Nombre string `json:"nombre"`
	Codigo int `json:"codigo"`
	Direccion string `json:"direccion"`
	Ciudad string `json:"ciudad"`
	Telefono int `json:"telefono"`
	Capacidad_Maxima int `json:"capacidad_maxima"`


}