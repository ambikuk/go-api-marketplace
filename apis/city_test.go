package apis

import (
	"net/http"
	"testing"

	"github.com/ambikuk/go-api-marketplace/daos"
	"github.com/ambikuk/go-api-marketplace/services"
	"github.com/ambikuk/go-api-marketplace/testdata"
)

func TestCity(t *testing.T) {
	testdata.ResetDB()
	router := newRouter()
	ServeCityResource(&router.RouteGroup, services.NewCityService(daos.NewCityDAO()))

	runAPITests(t, router, []apiTestCase{
		{"t10 - get a list of citys", "GET", "/cities?province_id=32", "", http.StatusOK, `{"page":1,"per_page":100,"page_count":1,"total_count":27,"items":[{"id":3201,"name":"KABUPATEN BOGOR","province_id":32},{"id":3202,"name":"KABUPATEN SUKABUMI","province_id":32},{"id":3203,"name":"KABUPATEN CIANJUR","province_id":32},{"id":3204,"name":"KABUPATEN BANDUNG","province_id":32},{"id":3205,"name":"KABUPATEN GARUT","province_id":32},{"id":3206,"name":"KABUPATEN TASIKMALAYA","province_id":32},{"id":3207,"name":"KABUPATEN CIAMIS","province_id":32},{"id":3208,"name":"KABUPATEN KUNINGAN","province_id":32},{"id":3209,"name":"KABUPATEN CIREBON","province_id":32},{"id":3210,"name":"KABUPATEN MAJALENGKA","province_id":32},{"id":3211,"name":"KABUPATEN SUMEDANG","province_id":32},{"id":3212,"name":"KABUPATEN INDRAMAYU","province_id":32},{"id":3213,"name":"KABUPATEN SUBANG","province_id":32},{"id":3214,"name":"KABUPATEN PURWAKARTA","province_id":32},{"id":3215,"name":"KABUPATEN KARAWANG","province_id":32},{"id":3216,"name":"KABUPATEN BEKASI","province_id":32},{"id":3217,"name":"KABUPATEN BANDUNG BARAT","province_id":32},{"id":3218,"name":"KABUPATEN PANGANDARAN","province_id":32},{"id":3271,"name":"KOTA BOGOR","province_id":32},{"id":3272,"name":"KOTA SUKABUMI","province_id":32},{"id":3273,"name":"KOTA BANDUNG","province_id":32},{"id":3274,"name":"KOTA CIREBON","province_id":32},{"id":3275,"name":"KOTA BEKASI","province_id":32},{"id":3276,"name":"KOTA DEPOK","province_id":32},{"id":3277,"name":"KOTA CIMAHI","province_id":32},{"id":3278,"name":"KOTA TASIKMALAYA","province_id":32},{"id":3279,"name":"KOTA BANJAR","province_id":32}]}`},
	})
}
