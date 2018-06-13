package daos

import (
	"testing"

	"github.com/ambikuk/go-api-marketplace/app"
	"github.com/ambikuk/go-api-marketplace/testdata"
	"github.com/stretchr/testify/assert"
)

func TestDistrictDAO(t *testing.T) {
	db := testdata.ResetDB()
	dao := NewDistrictDAO()

	{
		// Query
		testDBCall(db, func(rs app.RequestScope) {
			districts, err := dao.Query(rs, 1, 3, 3217)
			assert.Nil(t, err)
			assert.Equal(t, 3, len(districts))
		})
	}

	{
		// Count
		testDBCall(db, func(rs app.RequestScope) {
			count, err := dao.Count(rs, 3217)
			assert.Nil(t, err)
			assert.NotZero(t, count)
		})
	}
}
