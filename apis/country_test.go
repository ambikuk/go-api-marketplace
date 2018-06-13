package apis

import (
	"net/http"
	"testing"

	"github.com/ambikuk/go-api-marketplace/daos"
	"github.com/ambikuk/go-api-marketplace/services"
	"github.com/ambikuk/go-api-marketplace/testdata"
)

func TestCountry(t *testing.T) {
	testdata.ResetDB()
	router := newRouter()
	ServeCountryResource(&router.RouteGroup, services.NewCountryService(daos.NewCountryDAO()))

	notFoundError := `{"error_code":"NOT_FOUND", "message":"NOT_FOUND"}`

	runAPITests(t, router, []apiTestCase{
		{"t1 - get an artist", "GET", "/countries/2", "", http.StatusOK, `{"id":2,"iso":"AE","iso3":"ARE","iso_number":"784","name":"United Arab Emirates","capital":"Abu Dhabi","continent":"AS","currency":"AED","currency_name":"Dirham","phone":"971"}`},
		{"t2 - get a nonexisting artist", "GET", "/countries/99999", "", http.StatusNotFound, notFoundError},
		{"t10 - get a list of countries", "GET", "/countries?page=3&per_page=2", "", http.StatusOK, `{"page":3,"per_page":2,"page_count":123,"total_count":246,"items":[{"id":5,"iso":"AI","iso3":"AIA","iso_number":"660","name":"Anguilla","capital":"The Valley","continent":"NA","currency":"XCD","currency_name":"Dollar","phone":"+1-264"},{"id":6,"iso":"AL","iso3":"ALB","iso_number":"8","name":"Albania","capital":"Tirana","continent":"EU","currency":"ALL","currency_name":"Lek","phone":"355"}]}`},
	})
}
