
package controllers
import (
<<<<<<< Updated upstream
	"encoding/json"
	"strconv"

	"GESTOCK_MID/models"

	beego "github.com/beego/beego/v2/server/web"
)

// MantenimientoController controla las operaciones de mantenimiento.
type MantenimientoController struct {
	beego.Controller
}

// GET: Listar todos los mantenimientos.
func (c *MantenimientoController) GetAll() {
	mantenimientos, err := models.GetAllMantenimientos()
=======
	"errors"

	"github.com/beego/beego/v2/client/orm"
)

type Mantenimiento struct {
	Id                int    `orm:"column(id_mantenimiento);pk;auto" json:"id_mantenimiento"`
	Equipo            string `orm:"column(equipo)" json:"equipo"`
	TipoMantenimiento string `orm:"column(tipo_mantenimiento)" json:"tipo_mantenimiento"`
	FechaProgramada   string `orm:"column(fecha_programada)" json:"fecha_programada"`
	Responsable       string `orm:"column(responsable)" json:"responsable"`
	Estado            string `orm:"column(estado)" json:"estado"`
}

func (m *Mantenimiento) TableName() string {
	return "mantenimiento"
}

func init() {
	orm.RegisterModel(new(Mantenimiento))
}

// Obtener todos los mantenimientos
func GetAllMantenimientos() ([]Mantenimiento, error) {
	o := orm.NewOrm()
	var lista []Mantenimiento
	_, err := o.QueryTable(new(Mantenimiento)).All(&lista)
>>>>>>> Stashed changes
	if err != nil {
		return nil, err
	}
	return lista, nil
}

<<<<<<< Updated upstream
// POST: Crear un nuevo mantenimiento.
func (c *MantenimientoController) Post() {
	var nuevoMantenimiento models.Mantenimiento

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &nuevoMantenimiento); err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  400,
			"Message": "Datos inválidos o formato JSON incorrecto",
		}
		c.ServeJSON()
		return
	}

	if err := nuevoMantenimiento.Create(); err != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  500,
			"Message": "Error al guardar el mantenimiento",
		}
		c.ServeJSON()
		return
	}

	c.Ctx.Output.SetStatus(201)
	c.Data["json"] = map[string]interface{}{
		"Success": true,
		"Status":  201,
		"Message": "Mantenimiento creado exitosamente",
		"Data":    nuevoMantenimiento,
	}
	c.ServeJSON()
}

// PUT: Actualizar un mantenimiento existente.
// El ID se recibe mediante el parámetro de ruta :id.
func (c *MantenimientoController) Put() {
	idParam := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idParam)

	if err != nil || id <= 0 {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  400,
			"Message": "ID de mantenimiento inválido",
		}
		c.ServeJSON()
		return
	}

	var mantenimiento models.Mantenimiento

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &mantenimiento); err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  400,
			"Message": "Datos inválidos o formato JSON incorrecto",
		}
		c.ServeJSON()
		return
	}

	// El ID de la URL tiene prioridad sobre el ID enviado en el JSON.
	mantenimiento.ID = id

	if err := models.UpdateMantenimiento(&mantenimiento); err != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  500,
			"Message": "Error al actualizar el mantenimiento",
		}
		c.ServeJSON()
		return
	}

	c.Ctx.Output.SetStatus(200)
	c.Data["json"] = map[string]interface{}{
		"Success": true,
		"Status":  200,
		"Message": "Mantenimiento actualizado exitosamente",
		"Data":    mantenimiento,
	}
	c.ServeJSON()
}
=======
// Crear mantenimiento
func (m *Mantenimiento) Create() error {
	o := orm.NewOrm()
	_, err := o.Insert(m)
	return err
}

// Actualizar mantenimiento directamente con ORM
func (m *Mantenimiento) Update() error {
	o := orm.NewOrm()
	if o.Read(m) == nil {
		_, err := o.Update(m)
		return err
	}
	return errors.New("mantenimiento no encontrado")
}
>>>>>>> Stashed changes
