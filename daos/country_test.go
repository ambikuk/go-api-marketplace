package daos

import (
	"testing"

	"github.com/ambikuk/go-api-marketplace/app"
	"github.com/ambikuk/go-api-marketplace/testdata"
	"github.com/stretchr/testify/assert"
)

func TestCountryDAO(t *testing.T) {
	db := testdata.ResetDB()
	dao := NewCountryDAO()

	{
		// Get
		testDBCall(db, func(rs app.RequestScope) {
			country, err := dao.Get(rs, 2)
			assert.Nil(t, err)
			if assert.NotNil(t, country) {
				assert.Equal(t, 2, country.Id)
			}
		})
	}

	{
		// Query
		testDBCall(db, func(rs app.RequestScope) {
			countrys, err := dao.Query(rs, 1, 3)
			assert.Nil(t, err)
			assert.Equal(t, 3, len(countrys))
		})
	}

	{
		// Count
		testDBCall(db, func(rs app.RequestScope) {
			count, err := dao.Count(rs)
			assert.Nil(t, err)
			assert.NotZero(t, count)
		})
	}
}
