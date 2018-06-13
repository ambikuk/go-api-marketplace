package apis

import (
	"net/http"
	"testing"

	"github.com/ambikuk/go-api-marketplace/daos"
	"github.com/ambikuk/go-api-marketplace/services"
	"github.com/ambikuk/go-api-marketplace/testdata"
)

func TestProvince(t *testing.T) {
	testdata.ResetDB()
	router := newRouter()
	ServeProvinceResource(&router.RouteGroup, services.NewProvinceService(daos.NewProvinceDAO()))

	runAPITests(t, router, []apiTestCase{
		{"t10 - get a list of provinces", "GET", "/provinces?country_id=97", "", http.StatusOK, `{"page":1,"per_page":100,"page_count":1,"total_count":34,"items":[{"id":11,"name":"ACEH","country_id":97},{"id":12,"name":"SUMATERA UTARA","country_id":97},{"id":13,"name":"SUMATERA BARAT","country_id":97},{"id":14,"name":"RIAU","country_id":97},{"id":15,"name":"JAMBI","country_id":97},{"id":16,"name":"SUMATERA SELATAN","country_id":97},{"id":17,"name":"BENGKULU","country_id":97},{"id":18,"name":"LAMPUNG","country_id":97},{"id":19,"name":"KEPULAUAN BANGKA BELITUNG","country_id":97},{"id":21,"name":"KEPULAUAN RIAU","country_id":97},{"id":31,"name":"DKI JAKARTA","country_id":97},{"id":32,"name":"JAWA BARAT","country_id":97},{"id":33,"name":"JAWA TENGAH","country_id":97},{"id":34,"name":"DI YOGYAKARTA","country_id":97},{"id":35,"name":"JAWA TIMUR","country_id":97},{"id":36,"name":"BANTEN","country_id":97},{"id":51,"name":"BALI","country_id":97},{"id":52,"name":"NUSA TENGGARA BARAT","country_id":97},{"id":53,"name":"NUSA TENGGARA TIMUR","country_id":97},{"id":61,"name":"KALIMANTAN BARAT","country_id":97},{"id":62,"name":"KALIMANTAN TENGAH","country_id":97},{"id":63,"name":"KALIMANTAN SELATAN","country_id":97},{"id":64,"name":"KALIMANTAN TIMUR","country_id":97},{"id":65,"name":"KALIMANTAN UTARA","country_id":97},{"id":71,"name":"SULAWESI UTARA","country_id":97},{"id":72,"name":"SULAWESI TENGAH","country_id":97},{"id":73,"name":"SULAWESI SELATAN","country_id":97},{"id":74,"name":"SULAWESI TENGGARA","country_id":97},{"id":75,"name":"GORONTALO","country_id":97},{"id":76,"name":"SULAWESI BARAT","country_id":97},{"id":81,"name":"MALUKU","country_id":97},{"id":82,"name":"MALUKU UTARA","country_id":97},{"id":91,"name":"PAPUA BARAT","country_id":97},{"id":94,"name":"PAPUA","country_id":97}]}`},
	})
}
