package services

import (
	"github.com/ambikuk/go-api-marketplace/app"
	"github.com/ambikuk/go-api-marketplace/models"
)

// countryDAO specifies the interface of the artist DAO needed by CountryService.
type countryDAO interface {
	// Get returns the artist with the specified artist ID.
	Get(rs app.RequestScope, id int) (*models.Country, error)
	// Count returns the number of artists.
	Count(rs app.RequestScope) (int, error)
	// Query returns the list of artists with the given offset and limit.
	Query(rs app.RequestScope, offset, limit int) ([]models.Country, error)
}

// CountryService provides services related with artists.
type CountryService struct {
	dao countryDAO
}

// NewCountryService creates a new CountryService with the given artist DAO.
func NewCountryService(dao countryDAO) *CountryService {
	return &CountryService{dao}
}

// Get returns the artist with the specified the artist ID.
func (s *CountryService) Get(rs app.RequestScope, id int) (*models.Country, error) {
	return s.dao.Get(rs, id)
}

// Count returns the number of artists.
func (s *CountryService) Count(rs app.RequestScope) (int, error) {
	return s.dao.Count(rs)
}

// Query returns the artists with the specified offset and limit.
func (s *CountryService) Query(rs app.RequestScope, offset, limit int) ([]models.Country, error) {
	return s.dao.Query(rs, offset, limit)
}
