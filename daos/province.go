package daos

import (
	"github.com/ambikuk/go-api-marketplace/app"
	"github.com/ambikuk/go-api-marketplace/models"
	"github.com/go-ozzo/ozzo-dbx"
)

// ProvinceDAO persists province data in database
type ProvinceDAO struct{}

// NewProvinceDAO creates a new ProvinceDAO
func NewProvinceDAO() *ProvinceDAO {
	return &ProvinceDAO{}
}

// Count returns the number of the province records in the database.
func (dao *ProvinceDAO) Count(rs app.RequestScope, countryID int) (int, error) {
	var count int
	err := rs.Tx().Select("COUNT(*)").Where(dbx.HashExp{"country_id": countryID}).From("province").Row(&count)
	return count, err
}

// Query retrieves the province records with the specified offset and limit from the database.
func (dao *ProvinceDAO) Query(rs app.RequestScope, offset, limit int, countryID int) ([]models.Province, error) {
	provinces := []models.Province{}
	err := rs.Tx().Select().Where(dbx.HashExp{"country_id": countryID}).OrderBy("id").Offset(int64(offset)).Limit(int64(limit)).All(&provinces)
	return provinces, err
}
