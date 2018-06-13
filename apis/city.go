package apis

import (
	"github.com/ambikuk/go-api-marketplace/app"
	"github.com/ambikuk/go-api-marketplace/models"
	"github.com/go-ozzo/ozzo-routing"
)

type (
	// cityService specifies the interface for the city service needed by cityResource.
	cityService interface {
		Query(rs app.RequestScope, offset, limit int, provinceID int) ([]models.City, error)
		Count(rs app.RequestScope, provinceID int) (int, error)
	}

	// cityResource defines the handlers for the CRUD APIs.
	cityResource struct {
		service cityService
	}
)

// ServeCityResource sets up the routing of city endpoints and the corresponding handlers.
func ServeCityResource(rg *routing.RouteGroup, service cityService) {
	r := &cityResource{service}
	rg.Get("/cities", r.query)
}

func (r *cityResource) query(c *routing.Context) error {
	rs := app.GetRequestScope(c)
	provinceID := parseInt(c.Query("province_id"), 0)
	count, err := r.service.Count(rs, provinceID)
	if err != nil {
		return err
	}
	paginatedList := getPaginatedListFromRequest(c, count)
	items, err := r.service.Query(app.GetRequestScope(c), paginatedList.Offset(), paginatedList.Limit(), provinceID)
	if err != nil {
		return err
	}
	paginatedList.Items = items
	return c.Write(paginatedList)
}