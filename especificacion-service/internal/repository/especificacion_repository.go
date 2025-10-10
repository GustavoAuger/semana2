package repository

import (
	"especificacion-service/internal/model"
	"gorm.io/gorm"
)

type EspecificacionRepository interface {
	FindAll() ([]model.Especificacion, error)
	FindByID(id uint) (*model.Especificacion, error)
	FindByOfertaID(ofertaID int) (*model.Especificacion, error)
	Create(especificacion *model.Especificacion) error
	Update(especificacion *model.Especificacion) error
	Delete(id uint) error
}

type especificacionRepository struct {
	DB *gorm.DB
}

func NewEspecificacionRepository(db *gorm.DB) EspecificacionRepository {
	return &especificacionRepository{DB: db}
}

func (r *especificacionRepository) FindAll() ([]model.Especificacion, error) {
	var especificaciones []model.Especificacion
	if err := r.DB.Find(&especificaciones).Error; err != nil {
		return nil, err
	}
	return especificaciones, nil
}

func (r *especificacionRepository) FindByID(id uint) (*model.Especificacion, error) {
	var especificacion model.Especificacion
	if err := r.DB.First(&especificacion, id).Error; err != nil {
		return nil, err
	}
	return &especificacion, nil
}

func (r *especificacionRepository) FindByOfertaID(ofertaID int) (*model.Especificacion, error) {
	var especificacion model.Especificacion
	if err := r.DB.Where("oferta_id = ?", ofertaID).First(&especificacion).Error; err != nil {
		return nil, err
	}
	return &especificacion, nil
}

func (r *especificacionRepository) Create(especificacion *model.Especificacion) error {
	return r.DB.Create(especificacion).Error
}

func (r *especificacionRepository) Update(especificacion *model.Especificacion) error {
	return r.DB.Save(especificacion).Error
}

func (r *especificacionRepository) Delete(id uint) error {
	return r.DB.Delete(&model.Especificacion{}, id).Error
}
