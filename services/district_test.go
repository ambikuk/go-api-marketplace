package services

import (
	"testing"

	"github.com/ambikuk/go-api-marketplace/app"
	"github.com/ambikuk/go-api-marketplace/models"
	"github.com/stretchr/testify/assert"
)

func TestNewdistrictService(t *testing.T) {
	dao := newMockDistrictDAO()
	s := NewDistrictService(dao)
	assert.Equal(t, dao, s.dao)
}

func TestDistrictService_Query(t *testing.T) {
	s := NewDistrictService(newMockDistrictDAO())
	result, err := s.Query(nil, 1, 2, 97)
	if assert.Nil(t, err) {
		assert.Equal(t, 2, len(result))
	}
}

func newMockDistrictDAO() districtDAO {
	return &mockDistrictDAO{
		records: []models.District{
			{Id: 1, Name: "aaa", RegencyId: 97},
			{Id: 2, Name: "bbb", RegencyId: 97},
			{Id: 3, Name: "ccc", RegencyId: 97},
		},
	}
}
func TestDistrictService_Count(t *testing.T) {
	s := NewDistrictService(newMockDistrictDAO())
	result, err := s.Count(nil, 97)
	if assert.Nil(t, err) {
		assert.Equal(t, 3, result)
	}
}

type mockDistrictDAO struct {
	records []models.District
}

func (m *mockDistrictDAO) Query(rs app.RequestScope, offset, limit int, regencyID int) ([]models.District, error) {
	return m.records[offset : offset+limit], nil
}

func (m *mockDistrictDAO) Count(rs app.RequestScope, regencyID int) (int, error) {
	return len(m.records), nil
}
