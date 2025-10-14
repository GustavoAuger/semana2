package service

import (
	"errors"
	"especificacion-service/internal/clients/ofertaclient"
	"especificacion-service/internal/dto"
	"especificacion-service/internal/model"
	"especificacion-service/internal/repository"
)

type EspecificacionService interface {
	GetAllEspecificaciones() ([]dto.EspecificacionResponse, error)
	GetEspecificacionByID(id uint) (*dto.EspecificacionResponse, error)
	GetEspecificacionByOfertaID(ofertaID int) (*dto.EspecificacionResponse, error)
	CreateEspecificacion(especificacion *model.Especificacion) error
	UpdateEspecificacion(especificacion *model.Especificacion) error
	DeleteEspecificacion(id uint) error
}

type especificacionService struct {
	repo         repository.EspecificacionRepository
	ofertaClient *ofertaclient.Client
}

func NewEspecificacionService(repo repository.EspecificacionRepository, ofertaClient *ofertaclient.Client) EspecificacionService {
	return &especificacionService{
		repo:         repo,
		ofertaClient: ofertaClient,
	}
}

func (s *especificacionService) GetAllEspecificaciones() ([]dto.EspecificacionResponse, error) {
	especificaciones, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}

	responses := make([]dto.EspecificacionResponse, 0, len(especificaciones))
	for _, esp := range especificaciones {
		oferta, _ := s.ofertaClient.GetOfertaByID(int(esp.OfertaID))
		responses = append(responses, dto.EspecificacionResponse{
			Especificacion: esp,
			Oferta:         oferta,
		})
	}

	return responses, nil
}

func (s *especificacionService) GetEspecificacionByID(id uint) (*dto.EspecificacionResponse, error) {
	esp, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	oferta, _ := s.ofertaClient.GetOfertaByID(int(esp.OfertaID))
	return &dto.EspecificacionResponse{
		Especificacion: *esp,
		Oferta:         oferta,
	}, nil
}

func (s *especificacionService) GetEspecificacionByOfertaID(ofertaID int) (*dto.EspecificacionResponse, error) {
	esp, err := s.repo.FindByOfertaID(ofertaID)
	if err != nil {
		return nil, err
	}

	oferta, _ := s.ofertaClient.GetOfertaByID(ofertaID)
	return &dto.EspecificacionResponse{
		Especificacion: *esp,
		Oferta:         oferta,
	}, nil
}

func (s *especificacionService) CreateEspecificacion(especificacion *model.Especificacion) error {
	// Verificar si ya existe una especificación para esta oferta
	existing, _ := s.repo.FindByOfertaID(int(especificacion.OfertaID))
	if existing != nil {
		return errors.New("ya existe una especificación para esta oferta")
	}

	// Verificar que la oferta exista
	_, err := s.ofertaClient.GetOfertaByID(int(especificacion.OfertaID))
	if err != nil {
		return errors.New("la oferta especificada no existe")
	}

	return s.repo.Create(especificacion)
}

func (s *especificacionService) UpdateEspecificacion(especificacion *model.Especificacion) error {
	return s.repo.Update(especificacion)
}

func (s *especificacionService) DeleteEspecificacion(id uint) error {
	return s.repo.Delete(id)
}
