package services

import (
	"github.com/ambikuk/go-api-marketplace/app"
	"github.com/ambikuk/go-api-marketplace/models"
)

// districtDAO specifies the interface of the artist DAO needed by DistrictService.
type districtDAO interface {
	// Count returns the number of artists.
	Count(rs app.RequestScope, regencyID int) (int, error)
	// Query returns the list of artists with the given offset and limit.
	Query(rs app.RequestScope, offset, limit int, regencyID int) ([]models.District, error)
}

// DistrictService provides services related with artists.
type DistrictService struct {
	dao districtDAO
}

// NewDistrictService creates a new DistrictService with the given artist DAO.
func NewDistrictService(dao districtDAO) *DistrictService {
	return &DistrictService{dao}
}

// Count returns the number of artists.
func (s *DistrictService) Count(rs app.RequestScope,regencyID int) (int, error) {
	return s.dao.Count(rs, regencyID)
}

// Query returns the artists with the specified offset and limit.
func (s *DistrictService) Query(rs app.RequestScope, offset, limit ,regencyID int) ([]models.District, error) {
	return s.dao.Query(rs, offset, limit, regencyID)
}
