package models

import (
	"context"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type CategoriaPostgres struct {
	ID                int       `json:"id_categoria"`
	Nombre            string    `json:"nombre"`
	Activo            bool      `json:"activo"`
	FechaCreacion     time.Time `json:"fecha_creacion"`
	FechaModificacion time.Time `json:"fecha_modificacion"`
}

func ListarCategorias(ctx context.Context) ([]CategoriaPostgres, error) {
	db, err := orm.GetDB("default")
	if err != nil {
		return nil, err
	}

	filas, err := db.QueryContext(ctx, `
        SELECT id, nombre, activo, fecha_creacion, fecha_modificacion
        FROM inventario.categoria
        ORDER BY nombre, id
    `)
	if err != nil {
		return nil, err
	}
	defer filas.Close()

	categorias := make([]CategoriaPostgres, 0)
	for filas.Next() {
		var categoria CategoriaPostgres
		if err := filas.Scan(
			&categoria.ID, &categoria.Nombre, &categoria.Activo,
			&categoria.FechaCreacion, &categoria.FechaModificacion,
		); err != nil {
			return nil, err
		}
		categorias = append(categorias, categoria)
	}
	return categorias, filas.Err()
}

func CrearCategoria(ctx context.Context, nombre string) (CategoriaPostgres, error) {
	var categoria CategoriaPostgres
	db, err := orm.GetDB("default")
	if err != nil {
		return categoria, err
	}

	err = db.QueryRowContext(ctx, `
INSERT INTO inventario.categoria (nombre)
VALUES ($1)
RETURNING id, nombre, activo, fecha_creacion, fecha_modificacion
`, nombre).Scan(
		&categoria.ID, &categoria.Nombre, &categoria.Activo,
		&categoria.FechaCreacion, &categoria.FechaModificacion,
	)
	return categoria, err
}

func BuscarCategoriaPorID(ctx context.Context, id int64) (CategoriaPostgres, error) {
	var categoria CategoriaPostgres
	db, err := orm.GetDB("default")
	if err != nil {
		return categoria, err
	}

	err = db.QueryRowContext(ctx, `
SELECT id, nombre, activo, fecha_creacion, fecha_modificacion
FROM inventario.categoria
WHERE id = $1
`, id).Scan(
		&categoria.ID, &categoria.Nombre, &categoria.Activo,
		&categoria.FechaCreacion, &categoria.FechaModificacion,
	)
	return categoria, err
}

func ActualizarCategoria(ctx context.Context, id int64, nombre string) (CategoriaPostgres, error) {
	var categoria CategoriaPostgres
	db, err := orm.GetDB("default")
	if err != nil {
		return categoria, err
	}

	err = db.QueryRowContext(ctx, `
UPDATE inventario.categoria
SET nombre = $1
WHERE id = $2
RETURNING id, nombre, activo, fecha_creacion, fecha_modificacion
`, nombre, id).Scan(
		&categoria.ID, &categoria.Nombre, &categoria.Activo,
		&categoria.FechaCreacion, &categoria.FechaModificacion,
	)
	return categoria, err
}

func DesactivarCategoria(ctx context.Context, id int64) error {
	db, err := orm.GetDB("default")
	if err != nil {
		return err
	}

	var idDesactivado int64
	return db.QueryRowContext(ctx, `
UPDATE inventario.categoria
SET activo = false
WHERE id = $1 AND activo
RETURNING id
`, id).Scan(&idDesactivado)
}
