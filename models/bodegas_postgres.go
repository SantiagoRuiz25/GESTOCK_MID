package models

import (
	"context"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type BodegaPostgres struct {
	ID                int       `json:"id_bodega"`
	Codigo            string    `json:"codigo"`
	Nombre            string    `json:"nombre"`
	Direccion         string    `json:"direccion"`
	Ciudad            string    `json:"ciudad"`
	Telefono          string    `json:"telefono"`
	Responsable       string    `json:"responsable"`
	Capacidad         int       `json:"capacidad_maxima"`
	Activo            bool      `json:"activo"`
	FechaCreacion     time.Time `json:"fecha_creacion"`
	FechaModificacion time.Time `json:"fecha_modificacion"`
}

func ListarBodegas(ctx context.Context) ([]BodegaPostgres, error) {
	db, err := orm.GetDB("default")
	if err != nil {
		return nil, err
	}

	filas, err := db.QueryContext(ctx, `
        SELECT id, codigo, nombre, direccion, ciudad, telefono,
               responsable, capacidad, activo, fecha_creacion, fecha_modificacion
        FROM inventario.bodega
        ORDER BY id
    `)
	if err != nil {
		return nil, err
	}
	defer filas.Close()

	bodegas := make([]BodegaPostgres, 0)
	for filas.Next() {
		var bodega BodegaPostgres
		if err := filas.Scan(
			&bodega.ID, &bodega.Codigo, &bodega.Nombre, &bodega.Direccion,
			&bodega.Ciudad, &bodega.Telefono, &bodega.Responsable,
			&bodega.Capacidad, &bodega.Activo, &bodega.FechaCreacion,
			&bodega.FechaModificacion,
		); err != nil {
			return nil, err
		}
		bodegas = append(bodegas, bodega)
	}

	return bodegas, filas.Err()
}

func CrearBodega(ctx context.Context, bodega *BodegaPostgres) error {
	db, err := orm.GetDB("default")
	if err != nil {
		return err
	}

	return db.QueryRowContext(ctx, `
        INSERT INTO inventario.bodega
            (codigo, nombre, direccion, ciudad, telefono, responsable, capacidad)
        VALUES ($1, $2, $3, $4, $5, $6, $7)
        RETURNING id, activo, fecha_creacion, fecha_modificacion
    `, bodega.Codigo, bodega.Nombre, bodega.Direccion, bodega.Ciudad,
		bodega.Telefono, bodega.Responsable, bodega.Capacidad,
	).Scan(&bodega.ID, &bodega.Activo, &bodega.FechaCreacion, &bodega.FechaModificacion)
}

func BuscarBodegaPorID(ctx context.Context, id int) (BodegaPostgres, error) {
	var bodega BodegaPostgres
	db, err := orm.GetDB("default")
	if err != nil {
		return bodega, err
	}

	err = db.QueryRowContext(ctx, `
        SELECT id, codigo, nombre, direccion, ciudad, telefono,
               responsable, capacidad, activo, fecha_creacion, fecha_modificacion
        FROM inventario.bodega
        WHERE id = $1
    `, id).Scan(
		&bodega.ID, &bodega.Codigo, &bodega.Nombre, &bodega.Direccion,
		&bodega.Ciudad, &bodega.Telefono, &bodega.Responsable, &bodega.Capacidad,
		&bodega.Activo, &bodega.FechaCreacion, &bodega.FechaModificacion,
	)
	return bodega, err
}

func ActualizarBodega(ctx context.Context, bodega *BodegaPostgres) error {
	db, err := orm.GetDB("default")
	if err != nil {
		return err
	}

	return db.QueryRowContext(ctx, `
        UPDATE inventario.bodega
        SET codigo = $1, nombre = $2, direccion = $3, ciudad = $4,
            telefono = $5, responsable = $6, capacidad = $7
        WHERE id = $8
        RETURNING activo, fecha_creacion, fecha_modificacion
    `, bodega.Codigo, bodega.Nombre, bodega.Direccion, bodega.Ciudad,
		bodega.Telefono, bodega.Responsable, bodega.Capacidad, bodega.ID,
	).Scan(&bodega.Activo, &bodega.FechaCreacion, &bodega.FechaModificacion)
}

func DesactivarBodega(ctx context.Context, id int) error {
	db, err := orm.GetDB("default")
	if err != nil {
		return err
	}

	var idDesactivado int
	return db.QueryRowContext(ctx, `
        UPDATE inventario.bodega
        SET activo = false
        WHERE id = $1 AND activo
        RETURNING id
    `, id).Scan(&idDesactivado)
}
