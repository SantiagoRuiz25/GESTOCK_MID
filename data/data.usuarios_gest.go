package data

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

type Usuario struct {
	ID              int       `json:"id"`
	Nombre          string    `json:"nombre"`
	Email           string    `json:"email"`
	PasswordHash    string    `json:"-"` 
	Rol             string    `json:"rol"`
	RolID           int       `json:"rol_id"`
	Activo          bool      `json:"activo"`
	HorasTrabajadas float64   `json:"horasTrabajadas"`
	EmpresaID       *int      `json:"empresa_id,omitempty"`
	CreadoEn        time.Time `json:"creado_en"`
	ActualizadoEn   time.Time `json:"actualizado_en"`
}

type Rol struct {
	ID          int    `json:"id"`
	Nombre      string `json:"nombre"`
	Descripcion string `json:"descripcion"`
}

type Permiso struct {
	ID     int    `json:"id"`
	Nombre string `json:"nombre"`
}

type RolPermiso struct {
	RolID     int `json:"rol_id"`
	PermisoID int `json:"permiso_id"`
}

type Sesion struct {
	ID             int       `json:"id"`
	UsuarioID      int       `json:"usuario_id"`
	Dispositivo    string    `json:"dispositivo"`
	Navegador      string    `json:"navegador"`
	Ubicacion      string    `json:"ubicacion"`
	IP             string    `json:"ip"`
	UltimaActividad time.Time `json:"ultima_actividad"`
	EsActual       bool      `json:"es_actual"`
}

var Usuarios = []Usuario{
	{
		ID:              1,
		Nombre:          "Administrador Sistema",
		Email:           "admin@mid.com",
		PasswordHash:    hashPassword("Admin123!"),
		Rol:             "Administrador",
		RolID:           1,
		Activo:          true,
		HorasTrabajadas: 160,
		EmpresaID:       getIntPtr(1),
		CreadoEn:        time.Now(),
		ActualizadoEn:   time.Now(),
	},
}

var Roles = []Rol{
	{ID: 1, Nombre: "Administrador", Descripcion: "Acceso total al sistema"},
	{ID: 2, Nombre: "Auditor", Descripcion: "Control y auditoría de operaciones"},
	{ID: 3, Nombre: "Operador", Descripcion: "Gestión operativa estándar"},
}

var Permisos = []Permiso{
	{ID: 1, Nombre: "Acceso al Panel"},
	{ID: 2, Nombre: "Gestión de Empresas"},
	{ID: 3, Nombre: "Control de Inventario"},
	{ID: 4, Nombre: "Gestión de Recepción"},
	{ID: 5, Nombre: "Ver Historial Logístico"},
	{ID: 6, Nombre: "Gestión de Auditorías"},
	{ID: 7, Nombre: "Administración de Roles y Usuarios"},
	{ID: 8, Nombre: "Control de Envíos"},
	{ID: 9, Nombre: "Módulo de Programación"},
	{ID: 10, Nombre: "Gestión de Incidencias"},
	{ID: 11, Nombre: "Generación de Reportes"},
	{ID: 12, Nombre: "Configuración del Sistema"},
}

var RolesPermisos = []RolPermiso{}
var Sesiones = []Sesion{}

func hashPassword(pwd string) string {
	hash := sha256.Sum256([]byte(pwd))
	return hex.EncodeToString(hash[:])
}

func getIntPtr(i int) *int {
	return &i
}