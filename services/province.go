package services

import (
	"github.com/ambikuk/go-api-marketplace/app"
	"github.com/ambikuk/go-api-marketplace/models"
)

// provinceDAO specifies the interface of the artist DAO needed by ProvinceService.
type provinceDAO interface {
	// Count returns the number of artists.
	Count(rs app.RequestScope, countryID int) (int, error)
	// Query returns the list of artists with the given offset and limit.
	Query(rs app.RequestScope, offset, limit int, countryID int) ([]models.Province, error)
}

// ProvinceService provides services related with artists.
type ProvinceService struct {
	dao provinceDAO
}

// NewProvinceService creates a new ProvinceService with the given artist DAO.
func NewProvinceService(dao provinceDAO) *ProvinceService {
	return &ProvinceService{dao}
}

// Count returns the number of artists.
func (s *ProvinceService) Count(rs app.RequestScope,countryID int) (int, error) {
	return s.dao.Count(rs, countryID)
}

// Query returns the artists with the specified offset and limit.
func (s *ProvinceService) Query(rs app.RequestScope, offset, limit ,countryID int) ([]models.Province, error) {
	return s.dao.Query(rs, offset, limit, countryID)
}
