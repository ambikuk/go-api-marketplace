package services

import (
	"testing"

	"github.com/ambikuk/go-api-marketplace/app"
	"github.com/ambikuk/go-api-marketplace/models"
	"github.com/stretchr/testify/assert"
)

func TestNewcityService(t *testing.T) {
	dao := newMockCityDAO()
	s := NewCityService(dao)
	assert.Equal(t, dao, s.dao)
}

func TestCityService_Query(t *testing.T) {
	s := NewCityService(newMockCityDAO())
	result, err := s.Query(nil, 1, 2, 97)
	if assert.Nil(t, err) {
		assert.Equal(t, 2, len(result))
	}
}

func newMockCityDAO() cityDAO {
	return &mockCityDAO{
		records: []models.City{
			{Id: 1, Name: "aaa"},
			{Id: 2, Name: "bbb"},
			{Id: 3, Name: "ccc"},
		},
	}
}
func TestCityService_Count(t *testing.T) {
	s := NewCityService(newMockCityDAO())
	result, err := s.Count(nil, 32)
	if assert.Nil(t, err) {
		assert.Equal(t, 3, result)
	}
}

type mockCityDAO struct {
	records []models.City
}

func (m *mockCityDAO) Query(rs app.RequestScope, offset, limit int, provinceID int) ([]models.City, error) {
	return m.records[offset : offset+limit], nil
}

func (m *mockCityDAO) Count(rs app.RequestScope, provinceID int) (int, error) {
	return len(m.records), nil
}
