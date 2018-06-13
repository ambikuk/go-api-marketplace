package daos

import (
	"github.com/ambikuk/go-api-marketplace/app"
	"github.com/ambikuk/go-api-marketplace/models"
	"github.com/go-ozzo/ozzo-dbx"
)

// CityDAO persists city data in database
type CityDAO struct{}

// NewCityDAO creates a new CityDAO
func NewCityDAO() *CityDAO {
	return &CityDAO{}
}

// Count returns the number of the city records in the database.
func (dao *CityDAO) Count(rs app.RequestScope, provinceID int) (int, error) {
	var count int
	err := rs.Tx().Select("COUNT(*)").Where(dbx.HashExp{"province_id": provinceID}).From("city").Row(&count)
	return count, err
}

// Query retrieves the city records with the specified offset and limit from the database.
func (dao *CityDAO) Query(rs app.RequestScope, offset, limit int, provinceID int) ([]models.City, error) {
	citys := []models.City{}
	err := rs.Tx().Select().Where(dbx.HashExp{"province_id": provinceID}).OrderBy("id").Offset(int64(offset)).Limit(int64(limit)).All(&citys)
	return citys, err
}
