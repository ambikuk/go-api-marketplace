package apis

import (
	"github.com/ambikuk/go-api-marketplace/app"
	"github.com/ambikuk/go-api-marketplace/models"
	"github.com/go-ozzo/ozzo-routing"
)

type (
	// districtService specifies the interface for the district service needed by districtResource.
	districtService interface {
		Query(rs app.RequestScope, offset, limit int, regencyID int) ([]models.District, error)
		Count(rs app.RequestScope, regencyID int) (int, error)
	}

	// districtResource defines the handlers for the CRUD APIs.
	districtResource struct {
		service districtService
	}
)

// ServeDistrictResource sets up the routing of district endpoints and the corresponding handlers.
func ServeDistrictResource(rg *routing.RouteGroup, service districtService) {
	r := &districtResource{service}
	rg.Get("/districts", r.query)
}

func (r *districtResource) query(c *routing.Context) error {
	rs := app.GetRequestScope(c)
	regencyID := parseInt(c.Query("regency_id"), 0)
	count, err := r.service.Count(rs, regencyID)
	if err != nil {
		return err
	}
	paginatedList := getPaginatedListFromRequest(c, count)
	items, err := r.service.Query(app.GetRequestScope(c), paginatedList.Offset(), paginatedList.Limit(), regencyID)
	if err != nil {
		return err
	}
	paginatedList.Items = items
	return c.Write(paginatedList)
}