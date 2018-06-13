package services

import (
	"errors"
	"testing"

	"github.com/ambikuk/go-api-marketplace/app"
	"github.com/ambikuk/go-api-marketplace/models"
	"github.com/stretchr/testify/assert"
)

func TestNewcountryService(t *testing.T) {
	dao := newMockCountryDAO()
	s := NewCountryService(dao)
	assert.Equal(t, dao, s.dao)
}

func TestCountryService_Get(t *testing.T) {
	s := NewCountryService(newMockCountryDAO())
	country, err := s.Get(nil, 1)
	if assert.Nil(t, err) && assert.NotNil(t, country) {
		assert.Equal(t, "aaa", country.Name)
	}

	country, err = s.Get(nil, 100)
	assert.NotNil(t, err)
}

func TestCountryService_Query(t *testing.T) {
	s := NewCountryService(newMockCountryDAO())
	result, err := s.Query(nil, 1, 2)
	if assert.Nil(t, err) {
		assert.Equal(t, 2, len(result))
	}
}

func newMockCountryDAO() countryDAO {
	return &mockCountryDAO{
		records: []models.Country{
			{Id: 1, Name: "aaa"},
			{Id: 2, Name: "bbb"},
			{Id: 3, Name: "ccc"},
		},
	}
}

type mockCountryDAO struct {
	records []models.Country
}

func (m *mockCountryDAO) Get(rs app.RequestScope, id int) (*models.Country, error) {
	for _, record := range m.records {
		if record.Id == id {
			return &record, nil
		}
	}
	return nil, errors.New("not found")
}

func (m *mockCountryDAO) Query(rs app.RequestScope, offset, limit int) ([]models.Country, error) {
	return m.records[offset : offset+limit], nil
}

func (m *mockCountryDAO) Count(rs app.RequestScope) (int, error) {
	return len(m.records), nil
}

func (m *mockCountryDAO) Create(rs app.RequestScope, country *models.Country) error {
	if country.Id != 0 {
		return errors.New("Id cannot be set")
	}
	country.Id = len(m.records) + 1
	m.records = append(m.records, *country)
	return nil
}

func (m *mockCountryDAO) Update(rs app.RequestScope, id int, country *models.Country) error {
	country.Id = id
	for i, record := range m.records {
		if record.Id == id {
			m.records[i] = *country
			return nil
		}
	}
	return errors.New("not found")
}

func (m *mockCountryDAO) Delete(rs app.RequestScope, id int) error {
	for i, record := range m.records {
		if record.Id == id {
			m.records = append(m.records[:i], m.records[i+1:]...)
			return nil
		}
	}
	return errors.New("not found")
}
