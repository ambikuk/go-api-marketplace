package apis

import (
	"github.com/ambikuk/go-api-marketplace/app"
	"github.com/ambikuk/go-api-marketplace/models"
	"github.com/go-ozzo/ozzo-routing"
)

type (
	// provinceService specifies the interface for the province service needed by provinceResource.
	provinceService interface {
		Query(rs app.RequestScope, offset, limit int, countryID int) ([]models.Province, error)
		Count(rs app.RequestScope, countryID int) (int, error)
	}

	// provinceResource defines the handlers for the CRUD APIs.
	provinceResource struct {
		service provinceService
	}
)

// ServeProvinceResource sets up the routing of province endpoints and the corresponding handlers.
func ServeProvinceResource(rg *routing.RouteGroup, service provinceService) {
	r := &provinceResource{service}
	rg.Get("/provinces", r.query)
}

func (r *provinceResource) query(c *routing.Context) error {
	rs := app.GetRequestScope(c)
	countryID := parseInt(c.Query("country_id"), 0)
	count, err := r.service.Count(rs, countryID)
	if err != nil {
		return err
	}
	paginatedList := getPaginatedListFromRequest(c, count)
	items, err := r.service.Query(app.GetRequestScope(c), paginatedList.Offset(), paginatedList.Limit(), countryID)
	if err != nil {
		return err
	}
	paginatedList.Items = items
	return c.Write(paginatedList)
}