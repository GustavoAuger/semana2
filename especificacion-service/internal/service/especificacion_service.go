package service

import (
	"errors"
	"especificacion-service/internal/model"
	"especificacion-service/internal/repository"
)

type EspecificacionService interface {
	GetAllEspecificaciones() ([]model.Especificacion, error)
	GetEspecificacionByID(id uint) (*model.Especificacion, error)
	GetEspecificacionByOfertaID(ofertaID int) (*model.Especificacion, error)
	CreateEspecificacion(especificacion *model.Especificacion) error
	UpdateEspecificacion(especificacion *model.Especificacion) error
	DeleteEspecificacion(id uint) error
}

type especificacionService struct {
	repo repository.EspecificacionRepository
}

func NewEspecificacionService(repo repository.EspecificacionRepository) EspecificacionService {
	return &especificacionService{repo: repo}
}

func (s *especificacionService) GetAllEspecificaciones() ([]model.Especificacion, error) {
	return s.repo.FindAll()
}

func (s *especificacionService) GetEspecificacionByID(id uint) (*model.Especificacion, error) {
	return s.repo.FindByID(id)
}

func (s *especificacionService) GetEspecificacionByOfertaID(ofertaID int) (*model.Especificacion, error) {
	return s.repo.FindByOfertaID(ofertaID)
}

func (s *especificacionService) CreateEspecificacion(especificacion *model.Especificacion) error {
	// Verificar si ya existe una especificación para esta oferta
	existing, _ := s.repo.FindByOfertaID(int(especificacion.OfertaID))
	if existing != nil {
		return errors.New("ya existe una especificación para esta oferta")
	}
	return s.repo.Create(especificacion)
}

func (s *especificacionService) UpdateEspecificacion(especificacion *model.Especificacion) error {
	return s.repo.Update(especificacion)
}

func (s *especificacionService) DeleteEspecificacion(id uint) error {
	return s.repo.Delete(id)
}
