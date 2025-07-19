package dao

import (
	"fmt"
	"lgc/src/domain"
	"lgc/src/infraestructure/database"
	"time"

	"gorm.io/gorm"
)

type InscripcionDao struct {
	db *gorm.DB
}

func NewInscripcionDao(db *gorm.DB) *InscripcionDao {
	return &InscripcionDao{db: db}
}

type formularioDB struct {
	ID              int64     `gorm:"primaryKey;column:id"`
	Nombre          string    `gorm:"column:nombre"`
	Documento       string    `gorm:"column:documento"`
	Email           string    `gorm:"column:email"`
	Telefono        string    `gorm:"column:telefono"`
	Ciudad          string    `gorm:"column:ciudad"`
	Iglesia         string    `gorm:"column:iglesia"`
	HabeasData      bool      `gorm:"column:habeas_data"`
	MedioPago       string    `gorm:"column:medio_pago"`
	Estado          string    `gorm:"column:estado"`
	Asistencia      string    `gorm:"column:asistencia"`
	ComprobantePago string    `gorm:"column:comprobante_pago"`
	FechaRegistro   time.Time `gorm:"column:fecha_registro;autoCreateTime"`
}

func (formularioDB) TableName() string {
	return "formularios"
}

func toInscripcion(f formularioDB) *domain.Inscripcion {
	ins := domain.NewInscripcion(NewInscripcionDao(database.GetDB()))
	ins.SetID(f.ID)
	ins.SetNombre(f.Nombre)
	ins.SetDocumento(f.Documento)
	ins.SetEmail(f.Email)
	ins.SetTelefono(f.Telefono)
	ins.SetCiudad(f.Ciudad)
	ins.SetIglesia(f.Iglesia)
	ins.SetHabeasData(f.HabeasData)
	ins.SetEstado(f.Estado)
	ins.SetAsistencia(f.Asistencia)
	ins.SetComprobantePago(f.ComprobantePago)
	fechaStr := f.FechaRegistro.Format("2006-01-02 15:04:05")
	ins.SetFechaRegistro(fechaStr)
	return ins
}

func (i *InscripcionDao) Crear(inscripcion *domain.Inscripcion) bool {
	data := formularioDB{
		Nombre:          inscripcion.GetNombre(),
		Documento:       inscripcion.GetDocumento(),
		Email:           inscripcion.GetEmail(),
		Telefono:        inscripcion.GetTelefono(),
		Ciudad:          inscripcion.GetCiudad(),
		Iglesia:         inscripcion.GetIglesia(),
		HabeasData:      inscripcion.GetHabeasData(),
		MedioPago:       inscripcion.GetMedioPago(),
		Estado:          inscripcion.GetEstado(),
		Asistencia:      inscripcion.GetAsistencia(),
		ComprobantePago: inscripcion.GetComprobantePago(),
	}
	result := i.db.Create(&data)
	if result.Error == nil {
		inscripcion.SetID(data.ID)
		return true
	}

	return false
}

func (i *InscripcionDao) BuscarPorID(inscripcionID int64) *domain.Inscripcion {
	var f formularioDB
	if err := i.db.First(&f, inscripcionID).Error; err != nil {
		return toInscripcion(f)
	}
	return toInscripcion(f)
}

func (i *InscripcionDao) BuscarPorDocumento(documento string) *domain.Inscripcion {
	var f formularioDB
	if err := i.db.Where("documento = ? AND estado != 'Anulada'", documento).First(&f).Error; err != nil {
		return toInscripcion(f)
	}
	return toInscripcion(f)
}

func (i *InscripcionDao) ListarInscripcionesPorEstado(estado string) []domain.Inscripcion {
	var resultados []formularioDB
	if err := i.db.Where("estado = ?", estado).Order("fecha_registro DESC").Find(&resultados).Error; err != nil {
		return nil
	}
	var lista []domain.Inscripcion
	for _, f := range resultados {
		ins := toInscripcion(f)
		lista = append(lista, *ins)
	}
	return lista
}

func (i *InscripcionDao) Listar() []domain.Inscripcion {
	var resultados []formularioDB
	if err := i.db.Order("fecha_registro DESC").Find(&resultados).Error; err != nil {
		return nil
	}
	var lista []domain.Inscripcion
	for _, f := range resultados {
		ins := toInscripcion(f)
		lista = append(lista, *ins)
	}
	return lista
}

func (i *InscripcionDao) InscripcionAprobada(inscripcionID int64) bool {
	var estado string
	err := i.db.Table("formularios").Select("estado").Where("id = ?", inscripcionID).Scan(&estado).Error
	return err == nil && estado == "Aprobada"
}

func (i *InscripcionDao) Aprobar(inscripcionID int64) bool {

	result := i.db.Model(&formularioDB{}).
		Where("id = ?", inscripcionID).
		Update("estado", "Aprobada")

	fmt.Println(result.Error)
	return result.Error == nil && result.RowsAffected > 0
}

func (i *InscripcionDao) Anular(inscripcionID int64) bool {

	result := i.db.Model(&formularioDB{}).
		Where("id = ?", inscripcionID).
		Update("estado", "Anulada")

	fmt.Println(result.Error)
	return result.Error == nil && result.RowsAffected > 0
}

func (i *InscripcionDao) TotalInscripcionesPresenciales() int {
	var total int64
	err := i.db.Table("formularios").
		Where("asistencia = ?", "Presencial").
		Where("estado IN ?", []string{"PreAprobada", "Aprobada"}).
		Count(&total).Error

	if err != nil {
		return 0
	}
	return int(total) + 1
}

func (i *InscripcionDao) CrearConValidacionDeCupo(inscripcion *domain.Inscripcion, cupoMax int) (bool, error) {
	err := i.db.Transaction(func(tx *gorm.DB) error {
		var total int64
		if err := tx.Model(&formularioDB{}).
			Where("asistencia = 'Presencial' AND estado != 'Anulada'").
			Count(&total).Error; err != nil {
			return err
		}

		if total >= int64(cupoMax) {
			return fmt.Errorf("cupo lleno")
		}

		data := formularioDB{
			Nombre:          inscripcion.GetNombre(),
			Documento:       inscripcion.GetDocumento(),
			Email:           inscripcion.GetEmail(),
			Telefono:        inscripcion.GetTelefono(),
			Ciudad:          inscripcion.GetCiudad(),
			Iglesia:         inscripcion.GetIglesia(),
			MedioPago:       inscripcion.GetMedioPago(),
			HabeasData:      inscripcion.GetHabeasData(),
			Estado:          inscripcion.GetEstado(),
			Asistencia:      inscripcion.GetAsistencia(),
			ComprobantePago: inscripcion.GetComprobantePago(),
		}

		if err := tx.Create(&data).Error; err != nil {
			return err
		}

		inscripcion.SetID(data.ID)
		return nil
	})

	if err != nil {
		return false, err
	}

	return true, nil
}
