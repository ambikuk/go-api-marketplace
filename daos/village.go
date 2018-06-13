package daos

import (
	"github.com/ambikuk/go-api-marketplace/app"
	"github.com/ambikuk/go-api-marketplace/models"
	"github.com/go-ozzo/ozzo-dbx"
)

// VillageDAO persists village data in database
type VillageDAO struct{}

// NewVillageDAO creates a new VillageDAO
func NewVillageDAO() *VillageDAO {
	return &VillageDAO{}
}

// Count returns the number of the village records in the database.
func (dao *VillageDAO) Count(rs app.RequestScope, districtID int) (int, error) {
	var count int
	err := rs.Tx().Select("COUNT(*)").Where(dbx.HashExp{"district_id": districtID}).From("village").Row(&count)
	return count, err
}

// Query retrieves the village records with the specified offset and limit from the database.
func (dao *VillageDAO) Query(rs app.RequestScope, offset, limit int, districtID int) ([]models.Village, error) {
	villages := []models.Village{}
	err := rs.Tx().Select().Where(dbx.HashExp{"district_id": districtID}).OrderBy("id").Offset(int64(offset)).Limit(int64(limit)).All(&villages)
	return villages, err
}
