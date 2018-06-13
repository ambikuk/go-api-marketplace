package daos

import (
	"testing"

	"github.com/ambikuk/go-api-marketplace/app"
	"github.com/ambikuk/go-api-marketplace/testdata"
	"github.com/stretchr/testify/assert"
)

func TestVillageDAO(t *testing.T) {
	db := testdata.ResetDB()
	dao := NewVillageDAO()

	{
		// Query
		testDBCall(db, func(rs app.RequestScope) {
			villages, err := dao.Query(rs, 1, 3, 3217010)
			assert.Nil(t, err)
			assert.Equal(t, 3, len(villages))
		})
	}

	{
		// Count
		testDBCall(db, func(rs app.RequestScope) {
			count, err := dao.Count(rs, 3217010)
			assert.Nil(t, err)
			assert.NotZero(t, count)
		})
	}
}
