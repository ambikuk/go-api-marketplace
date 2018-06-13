package daos

import (
	"testing"

	"github.com/ambikuk/go-api-marketplace/app"
	"github.com/ambikuk/go-api-marketplace/testdata"
	"github.com/stretchr/testify/assert"
)

func TestCityDAO(t *testing.T) {
	db := testdata.ResetDB()
	dao := NewCityDAO()

	{
		// Query
		testDBCall(db, func(rs app.RequestScope) {
			city, err := dao.Query(rs, 1, 3, 32)
			assert.Nil(t, err)
			assert.Equal(t, 3, len(city))
		})
	}

	{
		// Count
		testDBCall(db, func(rs app.RequestScope) {
			count, err := dao.Count(rs, 32)
			assert.Nil(t, err)
			assert.NotZero(t, count)
		})
	}
}
