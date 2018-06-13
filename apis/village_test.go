package apis

import (
	"net/http"
	"testing"

	"github.com/ambikuk/go-api-marketplace/daos"
	"github.com/ambikuk/go-api-marketplace/services"
	"github.com/ambikuk/go-api-marketplace/testdata"
)

func TestVillage(t *testing.T) {
	testdata.ResetDB()
	router := newRouter()
	ServeVillageResource(&router.RouteGroup, services.NewVillageService(daos.NewVillageDAO()))

	runAPITests(t, router, []apiTestCase{
		{"t10 - get a list of villages", "GET", "/villages?district_id=3217010", "", http.StatusOK, `{"page":1,"per_page":100,"page_count":1,"total_count":8,"items":[{"id":3217010001,"name":"CICADAS","district_id":3217010},{"id":3217010002,"name":"CIBEDUG","district_id":3217010},{"id":3217010003,"name":"SUKAMANAH","district_id":3217010},{"id":3217010004,"name":"BOJONG","district_id":3217010},{"id":3217010005,"name":"BOJONGSALAM","district_id":3217010},{"id":3217010006,"name":"CINENGAH","district_id":3217010},{"id":3217010007,"name":"SUKARESMI","district_id":3217010},{"id":3217010008,"name":"CIBITUNG","district_id":3217010}]}`},
	})
}
