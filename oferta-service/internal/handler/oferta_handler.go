package handler

import (
	"net/http"
	"strconv"

	"oferta-service/internal/model"
	"oferta-service/internal/service"

	"github.com/gin-gonic/gin"
)

type OfertaHandler struct {
	ofertaService service.OfertaService
}

func NewOfertaHandler(service service.OfertaService) *OfertaHandler {
	return &OfertaHandler{ofertaService: service}
}

// GetOfertas obtiene todas las ofertas
func (h *OfertaHandler) GetOfertas(c *gin.Context) {
	ofertas, err := h.ofertaService.GetAllOfertas()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener las ofertas"})
		return
	}

	if ofertas == nil {
		ofertas = []model.Oferta{} // Retornar array vacío en lugar de null
	}

	c.JSON(http.StatusOK, ofertas)
}

// GetOferta obtiene una oferta por su ID
func (h *OfertaHandler) GetOferta(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	oferta, err := h.ofertaService.GetOfertaByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Oferta no encontrada"})
		return
	}

	c.JSON(http.StatusOK, oferta)
}
