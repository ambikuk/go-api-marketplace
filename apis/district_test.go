package apis

import (
	"net/http"
	"testing"

	"github.com/ambikuk/go-api-marketplace/daos"
	"github.com/ambikuk/go-api-marketplace/services"
	"github.com/ambikuk/go-api-marketplace/testdata"
)

func TestDistrict(t *testing.T) {
	testdata.ResetDB()
	router := newRouter()
	ServeDistrictResource(&router.RouteGroup, services.NewDistrictService(daos.NewDistrictDAO()))

	runAPITests(t, router, []apiTestCase{
		{"t10 - get a list of districts", "GET", "/districts?regency_id=3217", "", http.StatusOK, `{"page":1,"per_page":100,"page_count":1,"total_count":16,"items":[{"id":3217010,"name":"RONGGA","regency_id":3217},{"id":3217020,"name":"GUNUNGHALU","regency_id":3217},{"id":3217030,"name":"SINDANGKERTA","regency_id":3217},{"id":3217040,"name":"CILILIN","regency_id":3217},{"id":3217050,"name":"CIHAMPELAS","regency_id":3217},{"id":3217060,"name":"CIPONGKOR","regency_id":3217},{"id":3217070,"name":"BATUJAJAR","regency_id":3217},{"id":3217071,"name":"SAGULING","regency_id":3217},{"id":3217080,"name":"CIPATAT","regency_id":3217},{"id":3217090,"name":"PADALARANG","regency_id":3217},{"id":3217100,"name":"NGAMPRAH","regency_id":3217},{"id":3217110,"name":"PARONGPONG","regency_id":3217},{"id":3217120,"name":"LEMBANG","regency_id":3217},{"id":3217130,"name":"CISARUA","regency_id":3217},{"id":3217140,"name":"CIKALONG WETAN","regency_id":3217},{"id":3217150,"name":"CIPEUNDEUY","regency_id":3217}]}`},
	})
}
