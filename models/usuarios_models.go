package models

import (
	"time"
)

type Usuario struct {
	ID                int       `json:"id" db:"id"`
	Nombre            string    `json:"nombre" db:"nombre"`
	Email             string    `json:"email" db:"email"`
	PasswordHash      string    `json:"-" db:"password_hash"` // Oculto en JSON
	Rol               string    `json:"rol" db:"rol"`
	RolID             int       `json:"rol_id" db:"rol_id"`
	Activo            bool      `json:"activo" db:"activo"`
	HorasTrabajadas   float64   `json:"horasTrabajadas" db:"horas_trabajadas"`
	EmpresaID         *int      `json:"empresa_id,omitempty" db:"empresa_id"`
	CreadoEn          time.Time `json:"creado_en" db:"creado_en"`
	ActualizadoEn     time.Time `json:"actualizado_en" db:"actualizado_en"`
}

type Rol struct {
	ID          int    `json:"id" db:"id"`
	Nombre      string `json:"nombre" db:"nombre"`
	Descripcion string `json:"descripcion" db:"descripcion"`
}

type Permiso struct {
	ID     int    `json:"id" db:"id"`
	Nombre string `json:"nombre" db:"nombre"`
}

type RolPermiso struct {
	RolID     int `json:"rol_id" db:"rol_id"`
	PermisoID int `json:"permiso_id" db:"permiso_id"`
}

type SesionActiva struct {
	ID                int       `json:"id" db:"id"`
	UsuarioID         int       `json:"usuario_id" db:"usuario_id"`
	Dispositivo       string    `json:"dispositivo" db:"dispositivo"`
	Navegador         string    `json:"navegador" db:"navegador"`
	Ubicacion         string    `json:"ubicacion" db:"ubicacion"`
	IP                string    `json:"ip" db:"ip"`
	UltimaActividad   time.Time `json:"ultima_actividad" db:"ultima_actividad"`
	EsActual          bool      `json:"es_actual" db:"es_actual"`
}

// DTO para la respuesta enriquecida de la lista de usuarios (Caso Compuesto)
type UsuarioListaDTO struct {
	ID                int     `json:"id"`
	Nombre            string  `json:"nombre"`
	Email             string  `json:"email"`
	Rol               string  `json:"rol"`
	Activo            bool    `json:"activo"`
	EstadoTexto       string  `json:"estadoTexto"`
	HorasTrabajadas   float64 `json:"horasTrabajadas"`
}