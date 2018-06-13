package services

import (
	"github.com/ambikuk/go-api-marketplace/app"
	"github.com/ambikuk/go-api-marketplace/models"
)

// villageDAO specifies the interface of the artist DAO needed by VillageService.
type villageDAO interface {
	// Count returns the number of artists.
	Count(rs app.RequestScope, districtID int) (int, error)
	// Query returns the list of artists with the given offset and limit.
	Query(rs app.RequestScope, offset, limit int, districtID int) ([]models.Village, error)
}

// VillageService provides services related with artists.
type VillageService struct {
	dao villageDAO
}

// NewVillageService creates a new VillageService with the given artist DAO.
func NewVillageService(dao villageDAO) *VillageService {
	return &VillageService{dao}
}

// Count returns the number of artists.
func (s *VillageService) Count(rs app.RequestScope,districtID int) (int, error) {
	return s.dao.Count(rs, districtID)
}

// Query returns the artists with the specified offset and limit.
func (s *VillageService) Query(rs app.RequestScope, offset, limit ,districtID int) ([]models.Village, error) {
	return s.dao.Query(rs, offset, limit, districtID)
}
