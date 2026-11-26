package services

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type respuestaProductos struct {
	Success bool            `json:"Success"`
	Status  int             `json:"Status"`
	Message string          `json:"Message"`
	Data    json.RawMessage `json:"Data"`
}

type categoriaProducto struct {
	ID     int    `json:"id_categoria"`
	Nombre string `json:"nombre"`
}

type bodegaProducto struct {
	ID     int    `json:"id_bodega"`
	Nombre string `json:"nombre"`
}

func SolicitarProductoParaCliente(ctx context.Context, metodo string, id int64, cuerpo []byte) (int, json.RawMessage, error) {
	estado, datos, err := SolicitarProducto(ctx, metodo, id, cuerpo)
	if err != nil || metodo != http.MethodGet || estado != http.StatusOK {
		return estado, datos, err
	}

	var respuesta respuestaProductos
	if err := json.Unmarshal(datos, &respuesta); err != nil {
		return 0, nil, err
	}
	if !respuesta.Success || len(respuesta.Data) == 0 {
		return 0, nil, fmt.Errorf("respuesta de productos incompleta")
	}

	var productos []map[string]json.RawMessage
	if id == 0 {
		if err := json.Unmarshal(respuesta.Data, &productos); err != nil {
			return 0, nil, err
		}
		if len(productos) == 0 {
			return estado, datos, nil
		}
	} else {
		var producto map[string]json.RawMessage
		if err := json.Unmarshal(respuesta.Data, &producto); err != nil {
			return 0, nil, err
		}
		productos = []map[string]json.RawMessage{producto}
	}

	categorias, err := leerCatalogo[categoriaProducto](ctx, "/v1/categorias")
	if err != nil {
		return 0, nil, err
	}
	bodegas, err := leerCatalogo[bodegaProducto](ctx, "/v1/bodega")
	if err != nil {
		return 0, nil, err
	}

	nombresCategorias := make(map[int]string, len(categorias))
	for _, categoria := range categorias {
		nombresCategorias[categoria.ID] = categoria.Nombre
	}
	nombresBodegas := make(map[int]string, len(bodegas))
	for _, bodega := range bodegas {
		nombresBodegas[bodega.ID] = bodega.Nombre
	}

	for _, producto := range productos {
		var idCategoria, idBodega int
		if err := json.Unmarshal(producto["id_categoria"], &idCategoria); err != nil {
			return 0, nil, err
		}
		if err := json.Unmarshal(producto["id_bodega"], &idBodega); err != nil {
			return 0, nil, err
		}

		categoria, existeCategoria := nombresCategorias[idCategoria]
		bodega, existeBodega := nombresBodegas[idBodega]
		if !existeCategoria || !existeBodega {
			return 0, nil, fmt.Errorf("el catálogo no contiene la categoría %d o la bodega %d del producto", idCategoria, idBodega)
		}

		producto["categoria_nombre"], _ = json.Marshal(categoria)
		producto["bodega_nombre"], _ = json.Marshal(bodega)
	}

	if id == 0 {
		respuesta.Data, err = json.Marshal(productos)
	} else {
		respuesta.Data, err = json.Marshal(productos[0])
	}
	if err != nil {
		return 0, nil, err
	}

	resultado, err := json.Marshal(respuesta)
	return estado, resultado, err
}

func leerCatalogo[T any](ctx context.Context, ruta string) ([]T, error) {
	estado, datos, err := solicitarCRUD(ctx, http.MethodGet, ruta, 0, nil)
	if err != nil {
		return nil, fmt.Errorf("consultar %s: %w", ruta, err)
	}
	if estado != http.StatusOK {
		return nil, fmt.Errorf("consultar %s: HTTP %d", ruta, estado)
	}

	var respuesta struct {
		Success bool `json:"Success"`
		Data    []T  `json:"Data"`
	}
	if err := json.Unmarshal(datos, &respuesta); err != nil {
		return nil, fmt.Errorf("leer %s: %w", ruta, err)
	}
	if !respuesta.Success || respuesta.Data == nil {
		return nil, fmt.Errorf("respuesta incompleta de %s", ruta)
	}
	return respuesta.Data, nil
}
