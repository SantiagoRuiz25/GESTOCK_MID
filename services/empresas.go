package services

import (
	"context"
	"encoding/json"
)

func SolicitarEmpresa(ctx context.Context, metodo string, id int64, cuerpo []byte) (int, json.RawMessage, error) {
	return solicitarCRUD(ctx, metodo, "/v1/empresas", id, cuerpo)
}
