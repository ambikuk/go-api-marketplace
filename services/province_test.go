package services

import (
	"testing"

	"github.com/ambikuk/go-api-marketplace/app"
	"github.com/ambikuk/go-api-marketplace/models"
	"github.com/stretchr/testify/assert"
)

func TestNewprovinceService(t *testing.T) {
	dao := newMockProvinceDAO()
	s := NewProvinceService(dao)
	assert.Equal(t, dao, s.dao)
}

func TestProvinceService_Query(t *testing.T) {
	s := NewProvinceService(newMockProvinceDAO())
	result, err := s.Query(nil, 1, 2, 97)
	if assert.Nil(t, err) {
		assert.Equal(t, 2, len(result))
	}
}

func newMockProvinceDAO() provinceDAO {
	return &mockProvinceDAO{
		records: []models.Province{
			{Id: 1, Name: "aaa", ProvinceId: 97},
			{Id: 2, Name: "bbb", ProvinceId: 97},
			{Id: 3, Name: "ccc", ProvinceId: 97},
		},
	}
}
func TestProvinceService_Count(t *testing.T) {
	s := NewProvinceService(newMockProvinceDAO())
	result, err := s.Count(nil)
	if assert.Nil(t, err) {
		assert.Equal(t, 3, result)
	}
}

type mockProvinceDAO struct {
	records []models.Province
}

func (m *mockProvinceDAO) Query(rs app.RequestScope, offset, limit int, countryID int) ([]models.Province, error) {
	return m.records[offset : offset+limit], nil
}

func (m *mockProvinceDAO) Count(rs app.RequestScope, countryID int) (int, error) {
	return len(m.records), nil
}
