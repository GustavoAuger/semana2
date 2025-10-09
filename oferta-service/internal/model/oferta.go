package model

import (
	"time"

	"gorm.io/gorm"
)

// Oferta representa la entidad de oferta en el sistema
type Oferta struct {
	gorm.Model
	SolicitudID  int        `json:"solicitud_id" gorm:"column:solicitud_id"`
	Titulo       string     `json:"titulo" gorm:"type:varchar(200);not null"`
	Descripcion  string     `json:"descripcion" gorm:"type:text"`
	FechaInicio  *time.Time `json:"fecha_inicio" gorm:"type:date"`
	FechaFin     *time.Time `json:"fecha_fin" gorm:"type:date"`
	SalarioMin   float64    `json:"salario_min" gorm:"type:decimal(12,2)"`
	SalarioMax   float64    `json:"salario_max" gorm:"type:decimal(12,2)"`
	Moneda       string     `json:"moneda" gorm:"type:varchar(3);default:'USD'"`
	Ubicacion    string     `json:"ubicacion" gorm:"type:varchar(100)"`
	TipoContrato string     `json:"tipo_contrato" gorm:"type:varchar(50)"`
	Activa       bool       `json:"activa" gorm:"default:true"`
}

// TableName especifica el nombre de la tabla en la base de datos
func (Oferta) TableName() string {
	return "ofertas"
}
