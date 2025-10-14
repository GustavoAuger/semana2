package ofertaclient

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Client struct {
	baseURL string
}

// NewClient crea un nuevo cliente que usa Traefik como punto de entrada
func NewClient() *Client {
	// Usamos el nombre del servicio de Traefik en la red de Docker
	// No es necesario el puerto 80 ya que es el puerto HTTP por defecto
	return &Client{
		baseURL: "http://traefik/api/v1/ofertas",
	}
}

// GetOfertaByID obtiene una oferta por su ID a través de Traefik
func (c *Client) GetOfertaByID(id int) (map[string]interface{}, error) {
	// Construimos la URL usando la ruta base de Traefik
	url := fmt.Sprintf("%s/%d", c.baseURL, id)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error al hacer la petición al servicio de ofertas: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("error al obtener oferta: %s", string(body))
	}

	var oferta map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&oferta); err != nil {
		return nil, fmt.Errorf("error al decodificar la respuesta de oferta: %v", err)
	}

	return oferta, nil
}
