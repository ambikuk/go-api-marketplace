package daos

import (
	"testing"

	"github.com/ambikuk/go-api-marketplace/app"
	"github.com/ambikuk/go-api-marketplace/testdata"
	"github.com/stretchr/testify/assert"
)

func TestProvinceDAO(t *testing.T) {
	db := testdata.ResetDB()
	dao := NewProvinceDAO()

	{
		// Query
		testDBCall(db, func(rs app.RequestScope) {
			provinces, err := dao.Query(rs, 1, 3, 97)
			assert.Nil(t, err)
			assert.Equal(t, 3, len(provinces))
		})
	}

	{
		// Count
		testDBCall(db, func(rs app.RequestScope) {
			count, err := dao.Count(rs, 97)
			assert.Nil(t, err)
			assert.NotZero(t, count)
		})
	}
}
