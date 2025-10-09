package model

import (
	"time"

	"gorm.io/gorm"
)

// Especificacion representa la entidad de especificación en el sistema
type Especificacion struct {
	gorm.Model
	OfertaID       uint      `json:"oferta_id" gorm:"not null;uniqueIndex"`
	NumeroVacantes int       `json:"numero_vacantes" gorm:"not null"`
	Experiencia    string    `json:"experiencia" gorm:"type:varchar(100)"`
	Educacion      string    `json:"educacion" gorm:"type:varchar(200)"`
	Habilidades    string    `json:"habilidades" gorm:"type:text"`
	Idiomas        string    `json:"idiomas" gorm:"type:text"`
	Disponibilidad string    `json:"disponibilidad" gorm:"type:varchar(100)"`
	FechaInicio    time.Time `json:"fecha_inicio" gorm:"type:date"`
	FechaFin       time.Time `json:"fecha_fin" gorm:"type:date"`
	Activa         bool      `json:"activa" gorm:"default:true"`
}

// TableName especifica el nombre de la tabla en la base de datos
func (Especificacion) TableName() string {
	return "especificaciones"
}
