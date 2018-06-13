package daos

import (
	"github.com/ambikuk/go-api-marketplace/app"
	"github.com/ambikuk/go-api-marketplace/models"
	"github.com/go-ozzo/ozzo-dbx"
)

// DistrictDAO persists district data in database
type DistrictDAO struct{}

// NewDistrictDAO creates a new DistrictDAO
func NewDistrictDAO() *DistrictDAO {
	return &DistrictDAO{}
}

// Count returns the number of the district records in the database.
func (dao *DistrictDAO) Count(rs app.RequestScope, regencyID int) (int, error) {
	var count int
	err := rs.Tx().Select("COUNT(*)").Where(dbx.HashExp{"regency_id": regencyID}).From("district").Row(&count)
	return count, err
}

// Query retrieves the district records with the specified offset and limit from the database.
func (dao *DistrictDAO) Query(rs app.RequestScope, offset, limit int, regencyID int) ([]models.District, error) {
	districts := []models.District{}
	err := rs.Tx().Select().Where(dbx.HashExp{"regency_id": regencyID}).OrderBy("id").Offset(int64(offset)).Limit(int64(limit)).All(&districts)
	return districts, err
}
