package service

import (
	"oferta-service/internal/model"
	"oferta-service/internal/repository"
)

type OfertaService interface {
	GetAllOfertas() ([]model.Oferta, error)
	GetOfertaByID(id uint) (*model.Oferta, error)
}

type ofertaService struct {
	repo repository.OfertaRepository
}

func NewOfertaService(repo repository.OfertaRepository) OfertaService {
	return &ofertaService{repo: repo}
}

func (s *ofertaService) GetAllOfertas() ([]model.Oferta, error) {
	return s.repo.FindAll()
}

func (s *ofertaService) GetOfertaByID(id uint) (*model.Oferta, error) {
	return s.repo.FindByID(id)
}
