package services

import (
	"encoding/json"
	"errors"
	"GESTOCK_MID/models"

	"github.com/beego/beego/v2/client/orm"
)

// 1. Obtener todos los productos
func ObtenerTodosProductos() ([]models.Producto, error) {
	o := orm.NewOrm()
	var productos []models.Producto
	_, err := o.QueryTable(new(models.Producto)).All(&productos)
	if err != nil {
		return nil, err
	}
	return productos, nil
}

// 2. Obtener producto por ID
func ObtenerProductoPorId(id int) (*models.Producto, error) {
	o := orm.NewOrm()
	producto := &models.Producto{ID: id}
	err := o.Read(producto)
	if err != nil {
		return nil, errors.New("recurso no encontrado en el sistema")
	}
	return producto, nil
}

// 3. Crear un nuevo producto (recibe el JSON crudo del RequestBody)
func CrearProducto(body []byte) (int64, error) {
	var producto models.Producto
	if err := json.Unmarshal(body, &producto); err != nil {
		return 0, errors.New("el formato del JSON recibido es inválido")
	}

	o := orm.NewOrm()
	id, err := o.Insert(&producto)
	if err != nil {
		return 0, err
	}
	return id, nil
}

// 4. Actualizar un producto (recibe ID y el JSON crudo del RequestBody)
func ActualizarProducto(id int, body []byte) error {
	var producto models.Producto
	if err := json.Unmarshal(body, &producto); err != nil {
		return errors.New("el formato del JSON recibido es inválido")
	}

	producto.ID = id
	o := orm.NewOrm()
	_, err := o.Update(&producto)
	if err != nil {
		return err
	}
	return nil
}

// 5. Eliminar un producto por ID
func EliminarProducto(id int) error {
	o := orm.NewOrm()
	_, err := o.Delete(&models.Producto{ID: id})
	if err != nil {
		return err
	}
	return nil
}