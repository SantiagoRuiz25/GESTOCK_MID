package models

import (
	"context"

	"github.com/beego/beego/v2/client/orm"
)

func ListarProductos(ctx context.Context) ([]Producto, error) {
	db, err := orm.GetDB("default")
	if err != nil {
		return nil, err
	}

	filas, err := db.QueryContext(ctx, `
        SELECT id, id_categoria, id_bodega, codigo, nombre, precio, stock
        FROM public.producto
        ORDER BY id
    `)
	if err != nil {
		return nil, err
	}
	defer filas.Close()

	productos := make([]Producto, 0)
	for filas.Next() {
		var producto Producto
		if err := filas.Scan(
			&producto.ID, &producto.Id_categoria, &producto.Id_bodega,
			&producto.Codigo, &producto.Nombre, &producto.Precio, &producto.Stock,
		); err != nil {
			return nil, err
		}
		productos = append(productos, producto)
	}

	return productos, filas.Err()
}

func CrearProducto(ctx context.Context, producto *Producto) error {
	db, err := orm.GetDB("default")
	if err != nil {
		return err
	}

	return db.QueryRowContext(ctx, `
        INSERT INTO public.producto
            (id_categoria, id_bodega, codigo, nombre, precio, stock)
        VALUES ($1, $2, $3, $4, $5, $6)
        RETURNING id, precio
    `, producto.Id_categoria, producto.Id_bodega, producto.Codigo,
		producto.Nombre, producto.Precio, producto.Stock,
	).Scan(&producto.ID, &producto.Precio)
}

func BuscarProductoPorID(ctx context.Context, id int64) (Producto, error) {
	var producto Producto
	db, err := orm.GetDB("default")
	if err != nil {
		return producto, err
	}

	err = db.QueryRowContext(ctx, `
        SELECT id, id_categoria, id_bodega, codigo, nombre, precio, stock
        FROM public.producto
        WHERE id = $1
    `, id).Scan(
		&producto.ID, &producto.Id_categoria, &producto.Id_bodega,
		&producto.Codigo, &producto.Nombre, &producto.Precio, &producto.Stock,
	)
	return producto, err
}

func ActualizarProducto(ctx context.Context, producto *Producto) error {
	db, err := orm.GetDB("default")
	if err != nil {
		return err
	}

	return db.QueryRowContext(ctx, `
        UPDATE public.producto
        SET id_categoria = $1, id_bodega = $2, codigo = $3,
            nombre = $4, precio = $5, stock = $6
        WHERE id = $7
        RETURNING precio
    `, producto.Id_categoria, producto.Id_bodega, producto.Codigo,
		producto.Nombre, producto.Precio, producto.Stock, producto.ID,
	).Scan(&producto.Precio)
}

func EliminarProducto(ctx context.Context, id int64) error {
	db, err := orm.GetDB("default")
	if err != nil {
		return err
	}

	var idEliminado int64
	return db.QueryRowContext(ctx, `
        DELETE FROM public.producto
        WHERE id = $1
        RETURNING id
    `, id).Scan(&idEliminado)
}
