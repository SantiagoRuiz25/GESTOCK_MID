package services

import (
	"context"
	"encoding/json"
)

func SolicitarBodega(ctx context.Context, metodo string, id int64, cuerpo []byte) (int, json.RawMessage, error) {
	return solicitarCRUD(ctx, metodo, "/v1/bodega", id, cuerpo)
}
