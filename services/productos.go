package services

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

var clienteProductos = &http.Client{
	Timeout: 5 * time.Second,
	CheckRedirect: func(_ *http.Request, _ []*http.Request) error {
		return http.ErrUseLastResponse
	},
}

func SolicitarProducto(ctx context.Context, metodo string, id int64, cuerpo []byte) (int, json.RawMessage, error) {
	return solicitarCRUD(ctx, metodo, "/v1/productos", id, cuerpo)
}

func solicitarCRUD(ctx context.Context, metodo string, ruta string, id int64, cuerpo []byte) (int, json.RawMessage, error) {
	base := strings.TrimRight(os.Getenv("GESTOCK_CRUD_URL"), "/")
	if base == "" {
		base = "http://127.0.0.1:8082"
	}

	url := base + ruta
	if id > 0 {
		url += "/" + strconv.FormatInt(id, 10)
	}

	peticion, err := http.NewRequestWithContext(ctx, metodo, url, bytes.NewReader(cuerpo))
	if err != nil {
		return 0, nil, err
	}
	peticion.Header.Set("Accept", "application/json")
	if metodo == http.MethodPost || metodo == http.MethodPut {
		peticion.Header.Set("Content-Type", "application/json")
	}

	respuesta, err := clienteProductos.Do(peticion)
	if err != nil {
		return 0, nil, err
	}
	defer respuesta.Body.Close()

	datos, err := io.ReadAll(respuesta.Body)
	if err != nil {
		return 0, nil, err
	}
	if respuesta.StatusCode < 200 ||
		(respuesta.StatusCode >= 300 && respuesta.StatusCode < 400) ||
		!json.Valid(datos) {
		return 0, nil, fmt.Errorf("respuesta inválida del CRUD, HTTP %d", respuesta.StatusCode)
	}

	return respuesta.StatusCode, json.RawMessage(datos), nil
}
