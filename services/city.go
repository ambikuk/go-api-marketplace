package services

import (
	"github.com/ambikuk/go-api-marketplace/app"
	"github.com/ambikuk/go-api-marketplace/models"
)

// cityDAO specifies the interface of the artist DAO needed by CityService.
type cityDAO interface {
	// Count returns the number of artists.
	Count(rs app.RequestScope, provinceID int) (int, error)
	// Query returns the list of artists with the given offset and limit.
	Query(rs app.RequestScope, offset, limit int, provinceID int) ([]models.City, error)
}

// CityService provides services related with artists.
type CityService struct {
	dao cityDAO
}

// NewCityService creates a new CityService with the given artist DAO.
func NewCityService(dao cityDAO) *CityService {
	return &CityService{dao}
}

// Count returns the number of artists.
func (s *CityService) Count(rs app.RequestScope,provinceID int) (int, error) {
	return s.dao.Count(rs, provinceID)
}

// Query returns the artists with the specified offset and limit.
func (s *CityService) Query(rs app.RequestScope, offset, limit ,provinceID int) ([]models.City, error) {
	return s.dao.Query(rs, offset, limit, provinceID)
}
