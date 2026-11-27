package services

import (
	"context"
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"time"
)

type productoPanel struct {
	ID          int     `json:"id"`
	Nombre      string  `json:"nombre"`
	IDBodega    int     `json:"id_bodega"`
	Precio      float64 `json:"precio"`
	Stock       int     `json:"stock"`
	StockMinimo int     `json:"stock_minimo"`
	Activo      bool    `json:"activo"`
}

type bodegaPanel struct {
	ID        int    `json:"id_bodega"`
	Nombre    string `json:"nombre"`
	Capacidad int64  `json:"capacidad_maxima"`
	Activo    bool   `json:"activo"`
}

type empresaPanel struct {
	ID     int    `json:"id"`
	Nombre string `json:"nombre"`
	Activo bool   `json:"activo"`
}

type alertaPanel struct {
	ID          int    `json:"id"`
	Producto    string `json:"producto"`
	Bodega      string `json:"bodega"`
	Stock       int    `json:"stock"`
	StockMinimo int    `json:"stock_minimo"`
}

type movimientoPanel struct {
	ID          int64     `json:"id"`
	IDProducto  int       `json:"id_producto"`
	Producto    string    `json:"producto"`
	Tipo        string    `json:"tipo"`
	Cantidad    int       `json:"cantidad"`
	Fecha       time.Time `json:"fecha"`
	Responsable *string   `json:"responsable"`
}

type actividadPanel struct {
	UsuariosActivos    int               `json:"usuarios_activos"`
	UltimosMovimientos []movimientoPanel `json:"ultimos_movimientos"`
}

type ResumenPanel struct {
	ProductosActivos    int               `json:"productos_activos"`
	ValorInventario     float64           `json:"valor_inventario"`
	StockBajo           int               `json:"stock_bajo"`
	Agotados            int               `json:"agotados"`
	BodegasActivas      int               `json:"bodegas_activas"`
	CapacidadActiva     int64             `json:"capacidad_activa"`
	UnidadesOcupadas    int64             `json:"unidades_ocupadas"`
	OcupacionPorcentaje *float64          `json:"ocupacion_porcentaje"`
	UsuariosActivos     int               `json:"usuarios_activos"`
	EmpresasActivas     []empresaPanel    `json:"empresas_activas"`
	AlertasStock        []alertaPanel     `json:"alertas_stock"`
	UltimosMovimientos  []movimientoPanel `json:"ultimos_movimientos"`
}

func ObtenerResumenPanel(ctx context.Context) (ResumenPanel, error) {
	resumen := ResumenPanel{
		EmpresasActivas: make([]empresaPanel, 0),
		AlertasStock:    make([]alertaPanel, 0),
	}

	productos, err := leerCatalogo[productoPanel](ctx, "/v1/productos")
	if err != nil {
		return resumen, err
	}
	bodegas, err := leerCatalogo[bodegaPanel](ctx, "/v1/bodega")
	if err != nil {
		return resumen, err
	}
	empresas, err := leerCatalogo[empresaPanel](ctx, "/v1/empresas")
	if err != nil {
		return resumen, err
	}
	actividad, err := obtenerActividadPanel(ctx)
	if err != nil {
		return resumen, err
	}

	bodegasPorID := make(map[int]bodegaPanel, len(bodegas))
	for _, bodega := range bodegas {
		bodegasPorID[bodega.ID] = bodega
		if bodega.Activo {
			resumen.BodegasActivas++
			resumen.CapacidadActiva += bodega.Capacidad
		}
	}

	productosPorID := make(map[int]string, len(productos))
	for _, producto := range productos {
		productosPorID[producto.ID] = producto.Nombre
		if !producto.Activo {
			continue
		}

		resumen.ProductosActivos++
		resumen.ValorInventario += producto.Precio * float64(producto.Stock)

		bodega := bodegasPorID[producto.IDBodega]
		if bodega.Activo {
			resumen.UnidadesOcupadas += int64(producto.Stock)
		}

		switch {
		case producto.Stock == 0:
			resumen.Agotados++
		case producto.Stock <= producto.StockMinimo:
			resumen.StockBajo++
		default:
			continue
		}

		resumen.AlertasStock = append(resumen.AlertasStock, alertaPanel{
			ID: producto.ID, Producto: producto.Nombre, Bodega: bodega.Nombre,
			Stock: producto.Stock, StockMinimo: producto.StockMinimo,
		})
	}

	for _, empresa := range empresas {
		if empresa.Activo {
			resumen.EmpresasActivas = append(resumen.EmpresasActivas, empresa)
		}
	}

	if resumen.CapacidadActiva > 0 {
		porcentaje := 100 * float64(resumen.UnidadesOcupadas) / float64(resumen.CapacidadActiva)
		resumen.OcupacionPorcentaje = &porcentaje
	}
	resumen.ValorInventario = math.Round(resumen.ValorInventario*100) / 100

	resumen.UsuariosActivos = actividad.UsuariosActivos
	resumen.UltimosMovimientos = actividad.UltimosMovimientos
	for i := range resumen.UltimosMovimientos {
		movimiento := &resumen.UltimosMovimientos[i]
		movimiento.Producto = productosPorID[movimiento.IDProducto]
	}

	return resumen, nil
}

func obtenerActividadPanel(ctx context.Context) (actividadPanel, error) {
	var respuesta struct {
		Success bool            `json:"Success"`
		Data    *actividadPanel `json:"Data"`
	}

	estado, datos, err := solicitarCRUD(ctx, http.MethodGet, "/v1/panel/actividad", 0, nil)
	if err != nil {
		return actividadPanel{}, fmt.Errorf("consultar actividad: %w", err)
	}
	if estado != http.StatusOK {
		return actividadPanel{}, fmt.Errorf("consultar actividad: HTTP %d", estado)
	}
	if err := json.Unmarshal(datos, &respuesta); err != nil {
		return actividadPanel{}, fmt.Errorf("leer actividad: %w", err)
	}
	if !respuesta.Success || respuesta.Data == nil || respuesta.Data.UltimosMovimientos == nil {
		return actividadPanel{}, fmt.Errorf("respuesta de actividad incompleta")
	}

	return *respuesta.Data, nil
}
