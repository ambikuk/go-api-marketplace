package services

import (
	"testing"

	"github.com/ambikuk/go-api-marketplace/app"
	"github.com/ambikuk/go-api-marketplace/models"
	"github.com/stretchr/testify/assert"
)

func TestNewvillageService(t *testing.T) {
	dao := newMockVillageDAO()
	s := NewVillageService(dao)
	assert.Equal(t, dao, s.dao)
}

func TestVillageService_Query(t *testing.T) {
	s := NewVillageService(newMockVillageDAO())
	result, err := s.Query(nil, 1, 2, 97)
	if assert.Nil(t, err) {
		assert.Equal(t, 2, len(result))
	}
}

func newMockVillageDAO() villageDAO {
	return &mockVillageDAO{
		records: []models.Village{
			{Id: 1, Name: "aaa", DistrictId: 97},
			{Id: 2, Name: "bbb", DistrictId: 97},
			{Id: 3, Name: "ccc", DistrictId: 97},
		},
	}
}
func TestVillageService_Count(t *testing.T) {
	s := NewVillageService(newMockVillageDAO())
	result, err := s.Count(nil, 97)
	if assert.Nil(t, err) {
		assert.Equal(t, 3, result)
	}
}

type mockVillageDAO struct {
	records []models.Village
}

func (m *mockVillageDAO) Query(rs app.RequestScope, offset, limit int, districtID int) ([]models.Village, error) {
	return m.records[offset : offset+limit], nil
}

func (m *mockVillageDAO) Count(rs app.RequestScope, districtID int) (int, error) {
	return len(m.records), nil
}
