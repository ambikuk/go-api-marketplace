package daos

import (
	"github.com/ambikuk/go-api-marketplace/app"
	"github.com/ambikuk/go-api-marketplace/models"
)

// CountryDAO persists country data in database
type CountryDAO struct{}

// NewCountryDAO creates a new CountryDAO
func NewCountryDAO() *CountryDAO {
	return &CountryDAO{}
}

// Get reads the country with the specified ID from the database.
func (dao *CountryDAO) Get(rs app.RequestScope, id int) (*models.Country, error) {
	var country models.Country
	err := rs.Tx().Select().Model(id, &country)
	return &country, err
}

// Count returns the number of the country records in the database.
func (dao *CountryDAO) Count(rs app.RequestScope) (int, error) {
	var count int
	err := rs.Tx().Select("COUNT(*)").From("country").Row(&count)
	return count, err
}

// Query retrieves the country records with the specified offset and limit from the database.
func (dao *CountryDAO) Query(rs app.RequestScope, offset, limit int) ([]models.Country, error) {
	countrys := []models.Country{}
	err := rs.Tx().Select().OrderBy("id").Offset(int64(offset)).Limit(int64(limit)).All(&countrys)
	return countrys, err
}
